package syncer

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/filecoin-project/venus/pkg/consensus"
	"github.com/filecoin-project/venus/pkg/crypto"
	"github.com/filecoin-project/venus/pkg/repo"
	"github.com/filecoin-project/venus/pkg/statemanger"
	"github.com/hashicorp/go-multierror"
	"github.com/ipfs-force-community/metrics"

	"golang.org/x/sync/errgroup"

	actorsTypes "github.com/filecoin-project/go-state-types/actors"
	syncTypes "github.com/filecoin-project/venus/pkg/chainsync/types"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/venus/pkg/chain"
	"github.com/filecoin-project/venus/pkg/clock"
	"github.com/filecoin-project/venus/pkg/constants"
	"github.com/filecoin-project/venus/pkg/fork"
	"github.com/filecoin-project/venus/pkg/metrics/tracing"
	"github.com/filecoin-project/venus/pkg/net/exchange"
	"github.com/filecoin-project/venus/venus-shared/actors/policy"
	blockstoreutil "github.com/filecoin-project/venus/venus-shared/blockstore"
	"github.com/filecoin-project/venus/venus-shared/types"
	blockstore "github.com/ipfs/boxo/blockstore"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

// Syncer updates its chain.Store according to the methods of its
// consensus.Protocol.  It uses a bad tipset cache and a limit on new
// blocks to traverse during chain collection.  The Syncer can query the
// network for blocks.  The Syncer maintains the following invariant on
// its bsstore: all tipsets that pass the syncer's validity checks are added to the
// chain bsstore along with their state root CID.
//
// Ideally the code that syncs the chain according to consensus rules should
// be independent of any particular implementation of consensus.  Currently the
// Syncer is coupled to details of Expected Consensus. This dependence
// exists in the widen function, the fact that widen is called on only one
// tipset in the incoming chain, and assumptions regarding the existence of
// grandparent state in the bsstore.

var (
	// ErrForkTooLong is return when the syncing chain has fork with local
	ErrForkTooLong = fmt.Errorf("fork longer than threshold")
	// ErrChainHasBadTipSet is returned when the syncer traverses a chain with a cached bad tipset.
	ErrChainHasBadTipSet = errors.New("input chain contains a cached bad tipset")
	// ErrNewChainTooLong is returned when processing a fork that split off from the main chain too many blocks ago.
	ErrNewChainTooLong = errors.New("input chain forked from best chain past finality limit")
	// ErrUnexpectedStoreState indicates that the syncer's chain bsstore is violating expected invariants.
	ErrUnexpectedStoreState = errors.New("the chain bsstore is in an unexpected state")
	ErrForkCheckpoint       = fmt.Errorf("fork would require us to diverge from checkpointed block")

	logSyncer = logging.Logger("chainsync.syncer")
)

// metrics handlers
var (
	// epoch should not use as a label, because it has a lot of values
	syncOneTimer      = metrics.NewTimerMs("sync/sync_one", "Duration of single tipset validation in milliseconds")
	syncStatus        = metrics.NewInt64("sync/status", "The current status of the syncer. 0 = sync delay, 1 = sync done", "")
	reorgCnt          = metrics.NewCounter("chain/reorg_count", "The number of reorgs that have occurred.") // nolint
	epochGauge        = metrics.NewInt64("chain/epoch", "The current epoch.", "")
	netVersionGuage   = metrics.NewInt64("chain/net_version", "The current network version.", "")
	actorVersionGuage = metrics.NewInt64("chain/actor_version", "The current actor version.", "")
)

// StateProcessor does semantic validation on fullblocks.
type StateProcessor interface {
	// RunStateTransition returns the state root CID resulting from applying the input ts to the
	// prior `stateRoot`.  It returns an error if the transition is invalid.
	RunStateTransition(ctx context.Context, ts *types.TipSet) (root cid.Cid, receipt cid.Cid, err error)
}

// BlockValidator used to validate full block
type BlockValidator interface {
	ValidateFullBlock(ctx context.Context, blk *types.BlockHeader) error
}

// ChainReaderWriter reads and writes the chain bsstore.
type ChainReaderWriter interface {
	GetHead() *types.TipSet
	GetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
	GetTipSetStateRoot(context.Context, *types.TipSet) (cid.Cid, error)
	GetTipSetReceiptsRoot(context.Context, *types.TipSet) (cid.Cid, error)
	HasTipSetAndState(context.Context, *types.TipSet) bool
	GetTipsetMetadata(context.Context, *types.TipSet) (*chain.TipSetMetadata, error)
	PutTipSetMetadata(context.Context, *chain.TipSetMetadata) error
	SetHead(context.Context, *types.TipSet) error
	GetLatestBeaconEntry(context.Context, *types.TipSet) (*types.BeaconEntry, error)
	GetGenesisBlock(context.Context) (*types.BlockHeader, error)
}

// messageStore used to save and load message from db
type messageStore interface {
	LoadTipSetMessage(ctx context.Context, ts *types.TipSet) ([]types.BlockMessagesInfo, error)
	LoadMetaMessages(context.Context, cid.Cid) ([]*types.SignedMessage, []*types.Message, error)
	LoadReceipts(context.Context, cid.Cid) ([]types.MessageReceipt, error)
	StoreReceipts(context.Context, []types.MessageReceipt) (cid.Cid, error)
}

// ChainSelector chooses the heaviest between chains.
type ChainSelector interface {
	// IsHeavier returns true if tipset a is heavier than tipset b and false if
	// tipset b is heavier than tipset a.
	IsHeavier(ctx context.Context, a, b *types.TipSet) (bool, error)
	// Weight returns the weight of a tipset after the upgrade to version 1
	Weight(ctx context.Context, ts *types.TipSet) (big.Int, error)
}

// Syncer used to synchronize the block from the specified target, including acquiring the relevant block data and message data,
// verifying the block machine messages one by one and calculating them, checking the weight of the target after the calculation,
// and check whether it can become the latest tipset
type Syncer struct {
	exchangeClient exchange.Client
	// BadTipSetCache is used to filter out collections of invalid blocks.
	badTipSets *syncTypes.BadTipSetCache

	// Evaluates tipset messages and stores the resulting states.
	stmgr *statemanger.Stmgr
	// Validates headers and message structure
	blockValidator BlockValidator
	// Provides and stores validated tipsets and their state roots.
	chainStore *chain.Store
	// Provides message collections given cids
	messageProvider messageStore

	clock clock.Clock

	bsstore blockstoreutil.Blockstore

	fork fork.IFork

	delayRunTx *delayRunTsTransition
}

// NewSyncer constructs a Syncer ready for use.  The chain reader must have a
// head tipset to initialize the staging field.
func NewSyncer(stmgr *statemanger.Stmgr,
	hv BlockValidator,
	s *chain.Store,
	m messageStore,
	bsstore blockstoreutil.Blockstore,
	exchangeClient exchange.Client,
	c clock.Clock,
	fork fork.IFork,
) (*Syncer, error) {
	if constants.InsecurePoStValidation {
		logSyncer.Warn("*********************************************************************************************")
		logSyncer.Warn(" [INSECURE-POST-VALIDATION] Insecure test validation is enabled. If you see this outside of a test, it is a severe bug! ")
		logSyncer.Warn("*********************************************************************************************")
	}

	syncer := &Syncer{
		exchangeClient:  exchangeClient,
		badTipSets:      syncTypes.NewBadTipSetCache(),
		blockValidator:  hv,
		bsstore:         bsstore,
		chainStore:      s,
		messageProvider: m,
		clock:           c,
		fork:            fork,
		stmgr:           stmgr,
	}

	defer func() {
		syncer.delayRunTx = newDelayRunTsTransition(syncer)
		syncer.delayRunTx.run()
	}()

	return syncer, nil
}

// syncOne syncs a single tipset with the chain bsstore. syncOne calculates the
// parent state of the tipset and calls into consensus to run a state transition
// in order to validate the tipset.  In the case the input tipset is valid,
// syncOne calls into consensus to check its weight, and then updates the head
// of the bsstore if this tipset is the heaviest.
// todo mark bad-block
func (syncer *Syncer) syncOne(ctx context.Context, parent, next *types.TipSet) error {
	logSyncer.Infof("syncOne tipset, height:%d, blocks:%s", next.Height(), next.Key().String())
	priorHeadKey := syncer.chainStore.GetHead()
	// if tipset is already priorHeadKey, we've been here before. do nothing.
	if priorHeadKey.Equals(next) {
		return nil
	}

	stopwatch := syncOneTimer.Start()
	defer stopwatch(ctx)

	var wg errgroup.Group
	for i := 0; i < next.Len(); i++ {
		blk := next.At(i)
		wg.Go(func() error {
			// Fetch the URL.
			err := syncer.blockValidator.ValidateFullBlock(ctx, blk)
			if err == nil {
				if err := syncer.chainStore.AddToTipSetTracker(ctx, blk); err != nil {
					return fmt.Errorf("failed to add validated header to tipset tracker: %w", err)
				}
			}
			return err
		})
	}
	err := wg.Wait()
	if err != nil {
		var rootNotMatch bool // nolint

		if merr, isok := err.(*multierror.Error); isok {
			for _, e := range merr.Errors {
				if isRootNotMatch(e) {
					rootNotMatch = true
					break
				}
			}
		} else {
			rootNotMatch = isRootNotMatch(err) // nolint
		}

		if rootNotMatch { // nolint
			// todo: should here rollback, and re-compute?
			_ = syncer.stmgr.Rollback(ctx, parent, next)
		}

		return fmt.Errorf("validate mining failed %w", err)
	}

	syncer.chainStore.PersistTipSetKey(ctx, next.Key())

	// chain related metrics
	height := next.Height()
	epochGauge.Set(ctx, int64(height))

	timeStamp := next.MinTimestamp()
	if timeStamp+repo.Config.NetworkParams.BlockDelay*uint64(time.Second) >= uint64(syncer.clock.Now().Unix()) {
		// update to latest
		syncStatus.Set(ctx, 1)
	} else {
		syncStatus.Set(ctx, 0)
	}
	netVersion := syncer.fork.GetNetworkVersion(ctx, height)
	netVersionGuage.Set(ctx, int64(netVersion))
	actorVersion, _ := actorsTypes.VersionForNetwork(netVersion)
	actorVersionGuage.Set(ctx, int64(actorVersion))

	return nil
}

func isRootNotMatch(err error) bool {
	return errors.Is(err, consensus.ErrStateRootMismatch) || errors.Is(err, consensus.ErrReceiptRootMismatch)
}

// HandleNewTipSet validates and syncs the chain rooted at the provided tipset
// to a chain bsstore.  Iff catchup is false then the syncer will set the head.
func (syncer *Syncer) HandleNewTipSet(ctx context.Context, target *syncTypes.Target) (err error) {
	ctx, span := trace.StartSpan(ctx, "Syncer.HandleNewTipSet")
	span.AddAttributes(trace.StringAttribute("tipset", target.Head.String()))
	span.AddAttributes(trace.Int64Attribute("height", int64(target.Head.Height())))

	now := time.Now()

	defer func() {
		if err != nil {
			target.Err = err
			target.State = syncTypes.StageSyncErrored
		} else {
			target.State = syncTypes.StageSyncComplete
		}
		tracing.AddErrorEndSpan(ctx, span, &err)
		span.End()
		logSyncer.Infof("handle tipset height %d, count %d, took %.4f(s)", target.Head.Height(), target.Head.Len(), time.Since(now).Seconds())
	}()

	logSyncer.Debugf("begin fetch and sync of chain with head %v from %s at height %v", target.Head.Key(), target.Sender.String(), target.Head.Height())
	head := syncer.chainStore.GetHead()
	// If the store already has this tipset then the syncer is finished.
	if target.Head.At(0).ParentWeight.LessThan(head.At(0).ParentWeight) {
		return errors.New("do not sync to a target with less weight")
	}

	if syncer.chainStore.HasTipSetAndState(ctx, target.Head) || target.Head.Key().Equals(head.Key()) {
		return errors.New("do not sync to a target has synced before")
	}

	if target.Head.Height() == head.Height() {
		// check if maybeHead is fully contained in headTipSet
		// meaning we already synced all the blocks that are a part of maybeHead
		// if that is the case, there is nothing for us to do
		// we need to exit out early, otherwise checkpoint-fork logic might wrongly reject it
		fullyContained := true
		for _, c := range target.Head.Cids() {
			if !head.Contains(c) {
				fullyContained = false
				break
			}
		}
		if fullyContained {
			return nil
		}
	}

	syncer.exchangeClient.AddPeer(target.Sender)
	tipsets, err := syncer.fetchChainBlocks(ctx, head, target.Head, false)
	if err != nil {
		return errors.Wrapf(err, "failure fetching or validating headers")
	}
	logSyncer.Debugf("fetch header success at %v %s ...", tipsets[0].Height(), tipsets[0].Key())

	if err = syncer.syncSegement(ctx, target, tipsets); err == nil {
		syncer.delayRunTx.update(tipsets[len(tipsets)-1])
	}

	return err
}

func (syncer *Syncer) syncSegement(ctx context.Context, target *syncTypes.Target, tipsets []*types.TipSet) error {
	parent, err := syncer.chainStore.GetTipSet(ctx, tipsets[0].Parents())
	if err != nil {
		return err
	}

	errProcessChan := make(chan error, 1)
	errProcessChan <- nil // init
	var wg sync.WaitGroup
	// todo  write a pipeline segment processor function
	if err = rangeProcess(tipsets, func(segTipset []*types.TipSet) error {
		// fetch messages
		startTip := segTipset[0].Height()
		emdTipset := segTipset[len(segTipset)-1].Height()
		logSyncer.Debugf("start to fetch message segement %d-%d", startTip, emdTipset)
		_, err := syncer.fetchSegMessage(ctx, segTipset)
		if err != nil {
			return err
		}
		logSyncer.Debugf("finish to fetch message segement %d-%d", startTip, emdTipset)
		err = <-errProcessChan
		if err != nil {
			return fmt.Errorf("process message failed %v", err)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			logSyncer.Debugf("start to process message segement %d-%d", startTip, emdTipset)
			defer logSyncer.Debugf("finish to process message segement %d-%d", startTip, emdTipset)
			var processErr error
			parent, processErr = syncer.processTipSetSegment(ctx, target, parent, segTipset)
			if processErr != nil {
				errProcessChan <- processErr
				return
			}
			if !parent.Key().Equals(syncer.chainStore.GetCheckPoint().Key()) {
				logSyncer.Debugf("set chain head, height:%d, blocks:%d", parent.Height(), parent.Len())
				if err := syncer.chainStore.RefreshHeaviestTipSet(ctx, parent.Height()); err != nil {
					errProcessChan <- err
					return
				}
			}
			errProcessChan <- nil
		}()
		return nil
	}); err != nil {
		return err
	}

	wg.Wait()
	select {
	case err = <-errProcessChan:
		return err
	default:
		return nil
	}
}

// fetchChainBlocks get the block data, from targettip to knowntip.
// if local db has the block used that block
// if local db not exist, get block from network(libp2p),
// if there is a fork, get the common root tipset of knowntip and targettip, and return the block data from root tipset to targettip
// local(···->A->B) + incoming(C->D->E)  => ···->A->B->C->D->E
func (syncer *Syncer) fetchChainBlocks(ctx context.Context, knownTip *types.TipSet, targetTip *types.TipSet, ignoreCheckpoint bool) ([]*types.TipSet, error) {
	chainTipsets := []*types.TipSet{targetTip}
	flushDB := func(saveTips []*types.TipSet) error {
		bs := blockstoreutil.NewTemporary()
		cborStore := cbor.NewCborStore(bs)
		for _, tips := range saveTips {
			for _, blk := range tips.Blocks() {
				_, err := cborStore.Put(ctx, blk)
				if err != nil {
					return err
				}
			}
		}
		return blockstoreutil.CopyBlockstore(ctx, bs, syncer.bsstore)
	}

	untilHeight := knownTip.Height()
	count := 0
loop:
	for chainTipsets[len(chainTipsets)-1].Height() > untilHeight {
		tipSet, err := syncer.chainStore.GetTipSet(ctx, targetTip.Parents())
		if err == nil {
			chainTipsets = append(chainTipsets, tipSet)
			targetTip = tipSet
			count++
			if count%500 == 0 {
				logSyncer.Info("load from local db ", "Height: ", tipSet.Height())
			}
			continue
		}

		windows := targetTip.Height() - untilHeight
		if windows > 500 {
			windows = 500
		}

		fetchHeaders, err := syncer.exchangeClient.GetBlocks(ctx, targetTip.Parents(), int(windows))
		if err != nil {
			return nil, err
		}

		if len(fetchHeaders) == 0 {
			break loop
		}

		logSyncer.Infof("fetch blocks %d height from %d-%d", len(fetchHeaders), fetchHeaders[0].Height(), fetchHeaders[len(fetchHeaders)-1].Height())
		if err = flushDB(fetchHeaders); err != nil {
			return nil, err
		}
		for _, b := range fetchHeaders {
			if b.Height() < untilHeight {
				break loop
			}
			chainTipsets = append(chainTipsets, b)
			targetTip = b
		}
	}

	base := chainTipsets[len(chainTipsets)-1]
	if base.Equals(knownTip) {
		chainTipsets = chainTipsets[:len(chainTipsets)-1]
		base = chainTipsets[len(chainTipsets)-1]
	}

	if base.IsChildOf(knownTip) {
		// common case: receiving blocks that are building on top of our best tipset
		chain.Reverse(chainTipsets)
		return chainTipsets, nil
	}

	knownParent, err := syncer.chainStore.GetTipSet(ctx, knownTip.Parents())
	if err != nil {
		return nil, fmt.Errorf("failed to load next local tipset: %w", err)
	}

	if !ignoreCheckpoint {
		if chkpt := syncer.chainStore.GetCheckPoint(); chkpt != nil && base.Height() <= chkpt.Height() {
			return nil, fmt.Errorf("merge point affecting the checkpoint: %w", ErrForkCheckpoint)
		}
	}

	if base.IsChildOf(knownParent) {
		// common case: receiving a block thats potentially part of the same tipset as our best block
		chain.Reverse(chainTipsets)
		return chainTipsets, nil
	}

	logSyncer.Warnf("(fork detected) synced header chain, base: %v(%d), knownTip: %v(%d)", base.Key(), base.Height(),
		knownTip.Key(), knownTip.Height())
	fork, err := syncer.syncFork(ctx, base, knownTip, ignoreCheckpoint)
	if err != nil {
		if errors.Is(err, ErrForkTooLong) {
			// TODO: we're marking this block bad in the same way that we mark invalid blocks bad. Maybe distinguish?
			logSyncer.Warn("adding forked chain to our bad tipset cache")
			/*		for _, b := range incoming.Blocks() {
					syncer.bad.Add(b.Cid(), NewBadBlockReason(incoming.Cids(), "fork past finality"))
				}*/
		}
		return nil, fmt.Errorf("failed to sync fork: %w", err)
	}
	err = flushDB(fork)
	if err != nil {
		return nil, err
	}
	chainTipsets = append(chainTipsets, fork...)
	chain.Reverse(chainTipsets)
	return chainTipsets, nil
}

// syncFork tries to obtain the chain fragment that links a fork into a common
// ancestor in our view of the chain.
//
// If the fork is too long (build.ForkLengthThreshold), or would cause us to diverge from the checkpoint (ErrForkCheckpoint),
// we add the entire subchain to the denylist. Else, we find the common ancestor, and add the missing chain
// fragment until the fork point to the returned []TipSet.
//
//		D->E-F(targetTip）
//	A						 => D->E>F
//		B-C(knownTip)
func (syncer *Syncer) syncFork(ctx context.Context, incoming *types.TipSet, known *types.TipSet, ignoreCheckpoint bool) ([]*types.TipSet, error) {
	var chkpt *types.TipSet
	if !ignoreCheckpoint {
		chkpt = syncer.chainStore.GetCheckPoint()
		if known.Equals(chkpt) {
			return nil, ErrForkCheckpoint
		}
	}

	incomingParentsTsk := incoming.Parents()
	commonParent := false
	for _, incomingParent := range incomingParentsTsk.Cids() {
		if known.Contains(incomingParent) {
			commonParent = true
		}
	}

	if commonParent {
		// known contains at least one of incoming's Parents => the common ancestor is known's Parents (incoming's Grandparents)
		// in this case, we need to return {incoming.Parents()}
		incomingParents, err := syncer.chainStore.GetTipSet(ctx, incomingParentsTsk)
		if err != nil {
			// fallback onto the network
			tips, err := syncer.exchangeClient.GetBlocks(ctx, incoming.Parents(), 1)
			if err != nil {
				return nil, fmt.Errorf("failed to fetch incomingParents from the network: %w", err)
			}

			if len(tips) == 0 {
				return nil, fmt.Errorf("network didn't return any tipsets")
			}

			incomingParents = tips[0]
		}

		return []*types.TipSet{incomingParents}, nil
	}

	now := time.Now()
	defer func() {
		logSyncer.Infof("fork detected, fetch blocks %d(%v) took: %v", incoming.Height(), incoming.Parents(),
			time.Since(now))
	}()

	// TODO: Does this mean we always ask for ForkLengthThreshold blocks from the network, even if we just need, like, 2?
	// Would it not be better to ask in smaller chunks, given that an ~ForkLengthThreshold is very rare?
	tips, err := syncer.exchangeClient.GetBlocks(ctx, incoming.Parents(), int(policy.ChainFinality))
	if err != nil {
		return nil, err
	}

	genesisBlock, err := syncer.chainStore.GetGenesisBlock(ctx)
	if err != nil {
		return nil, err
	}

	nts, err := syncer.chainStore.GetTipSet(ctx, known.Parents())
	if err != nil {
		return nil, fmt.Errorf("failed to load next local tipset: %w", err)
	}

	for cur := 0; cur < len(tips); {
		if nts.Height() == 0 {
			if !genesisBlock.Equals(nts.At(0)) {
				return nil, fmt.Errorf("somehow synced chain that linked back to a different genesis (bad genesis: %s)", nts.Key())
			}
			return nil, fmt.Errorf("synced chain forked at genesis, refusing to sync; incoming: %s", incoming.ToSlice())
		}

		if nts.Equals(tips[cur]) {
			return tips[:cur], nil
		}

		if nts.Height() < tips[cur].Height() {
			cur++
		} else {
			nts, err = syncer.chainStore.GetTipSet(ctx, nts.Parents())
			if err != nil {
				return nil, fmt.Errorf("loading next local tipset: %w", err)
			}
		}
	}

	return nil, ErrForkTooLong
}

// fetchSegMessage get message in tipset
func (syncer *Syncer) fetchSegMessage(ctx context.Context, segTipset []*types.TipSet) ([]*types.FullTipSet, error) {
	// get message from local bsstore
	if len(segTipset) == 0 {
		return []*types.FullTipSet{}, nil
	}

	chain.Reverse(segTipset)
	defer chain.Reverse(segTipset)

	fullTipSets := make([]*types.FullTipSet, len(segTipset))
	defer types.ReverseFullBlock(fullTipSets)

	var leftChain []*types.TipSet
	var leftFullChain []*types.FullTipSet
	for index, tip := range segTipset {
		fullTipset, err := syncer.getFullBlock(ctx, tip)
		if err != nil {
			leftChain = segTipset[index:]
			leftFullChain = fullTipSets[index:]
			break
		}
		fullTipSets[index] = fullTipset
	}

	if len(leftChain) == 0 {
		return fullTipSets, nil
	}
	// fetch message from remote nodes
	bs := blockstoreutil.NewTemporary()
	cborStore := cbor.NewCborStore(bs)

	messages, err := syncer.exchangeClient.GetChainMessages(ctx, leftChain)
	if err != nil {
		return nil, err
	}

	for index := range messages {
		tip := leftChain[index]
		fts, err := zipTipSetAndMessages(bs, tip, messages[index].Bls, messages[index].Secpk, messages[index].BlsIncludes, messages[index].SecpkIncludes)
		if err != nil {
			return nil, fmt.Errorf("message processing failed: %w", err)
		}
		leftFullChain[index] = fts

		// save message
		for _, m := range messages[index].Bls {
			if _, err := cborStore.Put(ctx, m); err != nil {
				return nil, fmt.Errorf("BLS message processing failed: %w", err)
			}
		}

		for _, m := range messages[index].Secpk {
			if m.Signature.Type != crypto.SigTypeSecp256k1 && m.Signature.Type != crypto.SigTypeDelegated {
				return nil, fmt.Errorf("unknown signature type on message %s: %q", m.Cid(), m.Signature.Type)
			}
			if _, err := cborStore.Put(ctx, m); err != nil {
				return nil, fmt.Errorf("SECP message processing failed: %w", err)
			}
		}
	}

	err = blockstoreutil.CopyBlockstore(ctx, bs, syncer.bsstore)
	if err != nil {
		return nil, errors.Wrapf(err, "failure fetching full blocks")
	}
	return fullTipSets, nil
}

// getFullBlock get full block from message store
func (syncer *Syncer) getFullBlock(ctx context.Context, tipset *types.TipSet) (*types.FullTipSet, error) {
	fullBlocks := make([]*types.FullBlock, tipset.Len())
	for index, blk := range tipset.Blocks() {
		secpMsg, blsMsg, err := syncer.messageProvider.LoadMetaMessages(ctx, blk.Messages)
		if err != nil {
			return nil, err
		}
		fullBlocks[index] = &types.FullBlock{
			Header:       blk,
			BLSMessages:  blsMsg,
			SECPMessages: secpMsg,
		}
	}
	return types.NewFullTipSet(fullBlocks), nil
}

// processTipSetSegment process a batch of tipset in turn，
func (syncer *Syncer) processTipSetSegment(ctx context.Context, target *syncTypes.Target, parent *types.TipSet, segTipset []*types.TipSet) (*types.TipSet, error) {
	for i, ts := range segTipset {
		err := syncer.syncOne(ctx, parent, ts)
		if err != nil {
			// While `syncOne` can indeed fail for reasons other than consensus,
			// adding to the badTipSets at this point is the simplest, since we
			// have access to the chain. If syncOne fails for non-consensus reasons,
			// there is no assumption that the running node's data is valid at all,
			// so we don't really lose anything with this simplification.
			syncer.badTipSets.AddChain(segTipset[i:])
			return nil, errors.Wrapf(err, "failed to sync tipset %s, number %d of %d in chain", ts.Key().String(), i, len(segTipset))
		}
		parent = ts
		target.Current = ts
	}
	return parent, nil
}

// Head get latest head from chain store
func (syncer *Syncer) Head() *types.TipSet {
	return syncer.chainStore.GetHead()
}

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk.IsEmpty() {
		return fmt.Errorf("called with empty tsk")
	}

	ts, err := syncer.chainStore.GetTipSet(ctx, tsk)
	if err != nil {
		tss, err := syncer.exchangeClient.GetBlocks(ctx, tsk, 1)
		if err != nil {
			return fmt.Errorf("failed to fetch tipset: %w", err)
		} else if len(tss) != 1 {
			return fmt.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}

	head := syncer.Head()
	target := &syncTypes.Target{
		Head:    ts,
		Base:    syncer.Head(),
		Current: syncer.Head(),
		Start:   time.Now(),
	}
	if !head.Equals(ts) {
		if anc, err := syncer.chainStore.IsAncestorOf(ctx, ts, head); err != nil {
			return fmt.Errorf("failed to walk the chain when checkpointing: %w", err)
		} else if !anc {
			tipsets, err := syncer.fetchChainBlocks(ctx, head, target.Head, true)
			if err != nil {
				return errors.Wrapf(err, "failure fetching or validating headers")
			}
			logSyncer.Debugf("fetch header success at %v %s ...", tipsets[0].Height(), tipsets[0].Key())

			err = syncer.syncSegement(ctx, target, tipsets)
			if err != nil {
				return fmt.Errorf("failed to sync segment: %w", err)
			}

		} // else new checkpoint is on the current chain, we definitely have the tipsets.
	} // else current head, no need to switch.

	if err := syncer.chainStore.SetCheckpoint(ctx, ts); err != nil {
		return fmt.Errorf("failed to set the chain checkpoint: %w", err)
	}

	return nil
}

// TODO: this function effectively accepts unchecked input from the network,
// either validate it here, or ensure that its validated elsewhere (maybe make
// sure the blocksync code checks it?)
// maybe this code should actually live in blocksync??
func zipTipSetAndMessages(bs blockstore.Blockstore, ts *types.TipSet, allbmsgs []*types.Message, allsmsgs []*types.SignedMessage, bmi, smi [][]uint64) (*types.FullTipSet, error) {
	if len(ts.Blocks()) != len(smi) || len(ts.Blocks()) != len(bmi) {
		return nil, fmt.Errorf("msgincl length didn't match tipset size")
	}

	fts := &types.FullTipSet{}
	for bi, b := range ts.Blocks() {
		if msgc := len(bmi[bi]) + len(smi[bi]); msgc > constants.BlockMessageLimit {
			return nil, fmt.Errorf("block %q has too many messages (%d)", b.Cid(), msgc)
		}

		var smsgs []*types.SignedMessage
		var smsgCids []cid.Cid
		for _, m := range smi[bi] {
			smsgs = append(smsgs, allsmsgs[m])
			mCid := allsmsgs[m].Cid()
			smsgCids = append(smsgCids, mCid)
		}

		var bmsgs []*types.Message
		var bmsgCids []cid.Cid
		for _, m := range bmi[bi] {
			bmsgs = append(bmsgs, allbmsgs[m])
			mCid := allbmsgs[m].Cid()
			bmsgCids = append(bmsgCids, mCid)
		}

		mrcid, err := chain.ComputeMsgMeta(bs, bmsgCids, smsgCids)
		if err != nil {
			return nil, err
		}

		if b.Messages != mrcid {
			return nil, fmt.Errorf("messages didn't match message root in header for ts %s", ts.Key())
		}

		fb := &types.FullBlock{
			Header:       b,
			BLSMessages:  bmsgs,
			SECPMessages: smsgs,
		}

		fts.Blocks = append(fts.Blocks, fb)
	}

	return fts, nil
}

const maxProcessLen = 8

func rangeProcess(ts []*types.TipSet, cb func(ts []*types.TipSet) error) (err error) {
	for {
		if len(ts) == 0 {
			break
		} else if len(ts) < maxProcessLen {
			// break out if less than process len
			err = cb(ts)
			break
		} else {
			processTS := ts[0:maxProcessLen]
			err = cb(processTS)
			if err != nil {
				break
			}
			ts = ts[maxProcessLen:]
		}
		logSyncer.Infof("Sync Process End,Remaining: %v, err: %v ...", len(ts), err)
	}
	return err
}

type delayRunTsTransition struct { // nolint
	ch           chan *types.TipSet
	toRunTS      *types.TipSet
	syncer       *Syncer
	runningCount int64
}

func newDelayRunTsTransition(syncer *Syncer) *delayRunTsTransition {
	return &delayRunTsTransition{
		ch:     make(chan *types.TipSet, 10),
		syncer: syncer,
	}
}

func (d *delayRunTsTransition) run() {
	go d.listenUpdate()
}

func (d *delayRunTsTransition) stop() { // nolint
	close(d.ch)
}

func (d *delayRunTsTransition) update(ts *types.TipSet) {
	d.ch <- ts
}

func (d *delayRunTsTransition) listenUpdate() {
	duration := time.Second * 6
	ticker := time.NewTicker(duration)
	for {
		select {
		case t, isok := <-d.ch:
			if !isok {
				return
			}
			if d.toRunTS == nil || (d.toRunTS != nil && !d.toRunTS.Parents().Equals(t.Parents())) {
				ticker.Reset(duration)
			}
			d.toRunTS = t
		case <-ticker.C:
			if d.toRunTS != nil {
				if atomic.LoadInt64(&d.runningCount) < maxProcessLen {
					atomic.AddInt64(&d.runningCount, 1)
					go func(ts *types.TipSet) {
						_, _, err := d.syncer.stmgr.RunStateTransition(context.TODO(), ts, nil, false)
						if err != nil {
							logSyncer.Errorf("stmgr.runStateTransaction failed:%s", err.Error())
						}
						atomic.AddInt64(&d.runningCount, -1)
					}(d.toRunTS)
				}
				d.toRunTS = nil
			}
		}
	}
}
