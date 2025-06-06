package v1

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multihash"

	"github.com/filecoin-project/go-address"
	datatransfer "github.com/filecoin-project/go-data-transfer/v2"
	"github.com/filecoin-project/go-fil-markets/piecestore"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/paych"

	"github.com/filecoin-project/venus/venus-shared/api"
	"github.com/filecoin-project/venus/venus-shared/types"
	"github.com/filecoin-project/venus/venus-shared/types/gateway"
	"github.com/filecoin-project/venus/venus-shared/types/market"
)

type IMarket interface {
	ActorUpsert(context.Context, market.User) (bool, error)                   //perm:admin
	ActorDelete(context.Context, address.Address) error                       //perm:admin
	ActorList(context.Context) ([]market.User, error)                         //perm:read
	ActorExist(ctx context.Context, addr address.Address) (bool, error)       //perm:read
	ActorSectorSize(context.Context, address.Address) (abi.SectorSize, error) //perm:read

	MarketImportDealData(ctx context.Context, propcid cid.Cid, path string) error                                                                                                                               //perm:admin
	MarketImportPublishedDeal(ctx context.Context, deal market.MinerDeal) error                                                                                                                                 //perm:write
	MarketListDeals(ctx context.Context, addrs []address.Address) ([]*types.MarketDeal, error)                                                                                                                  //perm:read
	MarketGetDeal(ctx context.Context, dealPropCid cid.Cid) (*market.MinerDeal, error)                                                                                                                          //perm:read
	MarketListRetrievalDeals(ctx context.Context, params *market.RetrievalDealQueryParams) ([]market.ProviderDealState, error)                                                                                  //perm:read
	MarketGetRetrievalDeal(ctx context.Context, receiver peer.ID, dealID uint64) (*market.ProviderDealState, error)                                                                                             //perm:read
	MarketGetDealUpdates(ctx context.Context) (<-chan market.MinerDeal, error)                                                                                                                                  //perm:admin
	MarketListIncompleteDeals(ctx context.Context, params *market.StorageDealQueryParams) ([]market.MinerDeal, error)                                                                                           //perm:read
	MarketSetAsk(ctx context.Context, mAddr address.Address, price types.BigInt, verifiedPrice types.BigInt, duration abi.ChainEpoch, minPieceSize abi.PaddedPieceSize, maxPieceSize abi.PaddedPieceSize) error //perm:admin
	MarketGetAsk(ctx context.Context, mAddr address.Address) (*market.SignedStorageAsk, error)                                                                                                                  //perm:read
	MarketListStorageAsk(ctx context.Context) ([]*market.SignedStorageAsk, error)                                                                                                                               //perm:read
	MarketSetRetrievalAsk(ctx context.Context, mAddr address.Address, rask *retrievalmarket.Ask) error                                                                                                          //perm:admin
	MarketGetRetrievalAsk(ctx context.Context, mAddr address.Address) (*retrievalmarket.Ask, error)                                                                                                             //perm:read
	MarketListRetrievalAsk(ctx context.Context) ([]*market.RetrievalAsk, error)                                                                                                                                 //perm:read
	MarketListDataTransfers(ctx context.Context) ([]market.DataTransferChannel, error)                                                                                                                          //perm:admin
	MarketDataTransferUpdates(ctx context.Context) (<-chan market.DataTransferChannel, error)                                                                                                                   //perm:admin
	// MarketRestartDataTransfer attempts to restart a data transfer with the given transfer ID and other peer
	MarketRestartDataTransfer(ctx context.Context, transferID datatransfer.TransferID, otherPeer peer.ID, isInitiator bool) error //perm:admin
	// MarketCancelDataTransfer cancels a data transfer with the given transfer ID and other peer
	MarketCancelDataTransfer(ctx context.Context, transferID datatransfer.TransferID, otherPeer peer.ID, isInitiator bool) error //perm:admin
	MarketPendingDeals(ctx context.Context) ([]market.PendingDealInfo, error)                                                    //perm:write
	MarketPublishPendingDeals(ctx context.Context) error                                                                         //perm:admin

	PiecesListPieces(ctx context.Context) ([]cid.Cid, error)                                 //perm:read
	PiecesListCidInfos(ctx context.Context) ([]cid.Cid, error)                               //perm:read
	PiecesGetPieceInfo(ctx context.Context, pieceCid cid.Cid) (*piecestore.PieceInfo, error) //perm:read
	PiecesGetCIDInfo(ctx context.Context, payloadCid cid.Cid) (*piecestore.CIDInfo, error)   //perm:read

	DealsImportData(ctx context.Context, ref market.ImportDataRef, skipCommP bool) error                      //perm:admin
	DealsBatchImportData(ctx context.Context, refs market.ImportDataRefs) ([]*market.ImportDataResult, error) //perm:admin
	DealsImport(ctx context.Context, deals []*market.MinerDeal) error                                         //perm:admin

	DealsConsiderOnlineStorageDeals(context.Context, address.Address) (bool, error)      //perm:read
	DealsSetConsiderOnlineStorageDeals(context.Context, address.Address, bool) error     //perm:write
	DealsConsiderOnlineRetrievalDeals(context.Context, address.Address) (bool, error)    //perm:read
	DealsSetConsiderOnlineRetrievalDeals(context.Context, address.Address, bool) error   //perm:write
	DealsPieceCidBlocklist(context.Context, address.Address) ([]cid.Cid, error)          //perm:read
	DealsSetPieceCidBlocklist(context.Context, address.Address, []cid.Cid) error         //perm:write
	DealsConsiderOfflineStorageDeals(context.Context, address.Address) (bool, error)     //perm:read
	DealsSetConsiderOfflineStorageDeals(context.Context, address.Address, bool) error    //perm:write
	DealsConsiderOfflineRetrievalDeals(context.Context, address.Address) (bool, error)   //perm:read
	DealsSetConsiderOfflineRetrievalDeals(context.Context, address.Address, bool) error  //perm:write
	DealsConsiderVerifiedStorageDeals(context.Context, address.Address) (bool, error)    //perm:read
	DealsSetConsiderVerifiedStorageDeals(context.Context, address.Address, bool) error   //perm:write
	DealsConsiderUnverifiedStorageDeals(context.Context, address.Address) (bool, error)  //perm:read
	DealsSetConsiderUnverifiedStorageDeals(context.Context, address.Address, bool) error //perm:write
	// SectorGetExpectedSealDuration gets the time that a newly-created sector
	// waits for more deals before it starts sealing
	SectorGetExpectedSealDuration(context.Context, address.Address) (time.Duration, error) //perm:read
	// SectorSetExpectedSealDuration sets the expected time for a sector to seal
	SectorSetExpectedSealDuration(context.Context, address.Address, time.Duration) error    //perm:write
	DealsMaxStartDelay(context.Context, address.Address) (time.Duration, error)             //perm:read
	DealsSetMaxStartDelay(context.Context, address.Address, time.Duration) error            //perm:write
	MarketDataTransferPath(context.Context, address.Address) (string, error)                //perm:admin
	MarketSetDataTransferPath(context.Context, address.Address, string) error               //perm:admin
	DealsPublishMsgPeriod(context.Context, address.Address) (time.Duration, error)          //perm:read
	DealsSetPublishMsgPeriod(context.Context, address.Address, time.Duration) error         //perm:write
	MarketMaxDealsPerPublishMsg(context.Context, address.Address) (uint64, error)           //perm:read
	MarketSetMaxDealsPerPublishMsg(context.Context, address.Address, uint64) error          //perm:write
	DealsMaxProviderCollateralMultiplier(context.Context, address.Address) (uint64, error)  //perm:read
	DealsSetMaxProviderCollateralMultiplier(context.Context, address.Address, uint64) error //perm:write
	DealsMaxPublishFee(context.Context, address.Address) (types.FIL, error)                 //perm:read
	DealsSetMaxPublishFee(context.Context, address.Address, types.FIL) error                //perm:write
	MarketMaxBalanceAddFee(context.Context, address.Address) (types.FIL, error)             //perm:read
	MarketSetMaxBalanceAddFee(context.Context, address.Address, types.FIL) error            //perm:write

	// messager
	MessagerWaitMessage(ctx context.Context, mid cid.Cid) (*types.MsgLookup, error)                            //perm:read
	MessagerPushMessage(ctx context.Context, msg *types.Message, meta *types.MessageSendSpec) (cid.Cid, error) //perm:write
	MessagerGetMessage(ctx context.Context, mid cid.Cid) (*types.Message, error)                               //perm:read

	MarketAddBalance(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error)                   //perm:sign
	MarketGetReserved(ctx context.Context, addr address.Address) (types.BigInt, error)                                       //perm:sign
	MarketReserveFunds(ctx context.Context, wallet address.Address, addr address.Address, amt types.BigInt) (cid.Cid, error) //perm:sign
	MarketReleaseFunds(ctx context.Context, addr address.Address, amt types.BigInt) error                                    //perm:sign
	MarketWithdraw(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error)                     //perm:sign

	NetAddrsListen(context.Context) (peer.AddrInfo, error) //perm:read
	ID(context.Context) (peer.ID, error)                   //perm:read

	// DagstoreListShards returns information about all shards known to the
	// DAG store. Only available on nodes running the markets subsystem.
	DagstoreListShards(ctx context.Context) ([]market.DagstoreShardInfo, error) //perm:admin

	// DagstoreInitializeShard initializes an uninitialized shard.
	//
	// Initialization consists of fetching the shard's data (deal payload) from
	// the storage subsystem, generating an index, and persisting the index
	// to facilitate later retrievals, and/or to publish to external sources.
	//
	// This operation is intended to complement the initial migration. The
	// migration registers a shard for every unique piece CID, with lazy
	// initialization. Thus, shards are not initialized immediately to avoid
	// IO activity competing with proving. Instead, shard are initialized
	// when first accessed. This method forces the initialization of a shard by
	// accessing it and immediately releasing it. This is useful to warm up the
	// cache to facilitate subsequent retrievals, and to generate the indexes
	// to publish them externally.
	//
	// This operation fails if the shard is not in ShardStateNew state.
	// It blocks until initialization finishes.
	DagstoreInitializeShard(ctx context.Context, key string) error //perm:admin

	// DagstoreRecoverShard attempts to recover a failed shard.
	//
	// This operation fails if the shard is not in ShardStateErrored state.
	// It blocks until recovery finishes. If recovery failed, it returns the
	// error.
	DagstoreRecoverShard(ctx context.Context, key string) error //perm:admin

	// DagstoreInitializeAll initializes all uninitialized shards in bulk,
	// according to the policy passed in the parameters.
	//
	// It is recommended to set a maximum concurrency to avoid extreme
	// IO pressure if the storage subsystem has a large amount of deals.
	//
	// It returns a stream of events to report progress.
	DagstoreInitializeAll(ctx context.Context, params market.DagstoreInitializeAllParams) (<-chan market.DagstoreInitializeAllEvent, error) //perm:admin

	// DagstoreInitializeStorage initializes all pieces in specify storage
	DagstoreInitializeStorage(context.Context, string, market.DagstoreInitializeAllParams) (<-chan market.DagstoreInitializeAllEvent, error) //perm:admin

	// DagstoreGC runs garbage collection on the DAG store.
	DagstoreGC(ctx context.Context) ([]market.DagstoreShardResult, error) //perm:admin

	// DagstoreDestroyShard destroy shard by key
	DagstoreDestroyShard(ctx context.Context, key string) error //perm:admin

	MarkDealsAsPacking(ctx context.Context, miner address.Address, deals []abi.DealID) error                                                                          //perm:write
	UpdateDealOnPacking(ctx context.Context, miner address.Address, dealID abi.DealID, sectorid abi.SectorNumber, offset abi.PaddedPieceSize) error                   //perm:write
	UpdateDealStatus(ctx context.Context, miner address.Address, dealID abi.DealID, pieceStatus market.PieceStatus, dealStatus storagemarket.StorageDealStatus) error //perm:write
	GetDeals(ctx context.Context, miner address.Address, pageIndex, pageSize int) ([]*market.DealInfo, error)                                                         //perm:read
	AssignUnPackedDeals(ctx context.Context, sid abi.SectorID, ssize abi.SectorSize, spec *market.GetDealSpec) ([]*market.DealInfoIncludePath, error)                 //perm:write
	// ReleaseDeals is used to release the deals that have been assigned by AssignUnPackedDeals method.
	ReleaseDeals(ctx context.Context, miner address.Address, deals []abi.DealID) error                                                                //perm:write
	GetUnPackedDeals(ctx context.Context, miner address.Address, spec *market.GetDealSpec) ([]*market.DealInfoIncludePath, error)                     //perm:read
	UpdateStorageDealStatus(ctx context.Context, dealProposalCid cid.Cid, state storagemarket.StorageDealStatus, pieceState market.PieceStatus) error //perm:write
	AssignDeals(ctx context.Context, sid abi.SectorID, ssize abi.SectorSize, spec *market.GetDealSpec) ([]*market.DealInfoV2, error)                  //perm:write
	ReleaseDirectDeals(ctx context.Context, miner address.Address, allocationIDs []types.AllocationId) error                                          //perm:write
	// market event
	ResponseMarketEvent(ctx context.Context, resp *gateway.ResponseEvent) error                                        //perm:read
	ListenMarketEvent(ctx context.Context, policy *gateway.MarketRegisterPolicy) (<-chan *gateway.RequestEvent, error) //perm:read

	// Paych
	PaychVoucherList(ctx context.Context, pch address.Address) ([]*paych.SignedVoucher, error) //perm:read

	AddFsPieceStorage(ctx context.Context, name string, path string, readonly bool) error //perm:admin

	AddS3PieceStorage(ctx context.Context, name, endpoit, bucket, subdir, accessKey, secretKey, token string, readonly bool) error //perm:admin

	RemovePieceStorage(ctx context.Context, name string) error //perm:admin

	ListPieceStorageInfos(ctx context.Context) market.PieceStorageInfos //perm:read

	// GetStorageDealStatistic get storage deal statistic information
	// if set miner address to address.Undef, return all storage deal info
	GetStorageDealStatistic(ctx context.Context, miner address.Address) (*market.StorageDealStatistic, error) //perm:read

	// GetRetrievalDealStatistic get retrieval deal statistic information
	// todo address undefined is invalid, it is currently not possible to directly associate an order with a miner
	GetRetrievalDealStatistic(ctx context.Context, miner address.Address) (*market.RetrievalDealStatistic, error) //perm:read

	// ImportDirectDeal import direct deals
	ImportDirectDeal(ctx context.Context, deal *market.DirectDealParams) error                                   //perm:write
	GetDirectDeal(ctx context.Context, id uuid.UUID) (*market.DirectDeal, error)                                 //perm:read
	GetDirectDealByAllocationID(ctx context.Context, id types.AllocationId) (*market.DirectDeal, error)          //perm:read
	ListDirectDeals(ctx context.Context, queryParams market.DirectDealQueryParams) ([]*market.DirectDeal, error) //perm:read
	UpdateDirectDealState(ctx context.Context, id uuid.UUID, state market.DirectDealState) error                 //perm:write
	UpdateDirectDealPayloadCID(ctx context.Context, id uuid.UUID, payloadCID cid.Cid) error                      //perm:write

	UpdateStorageDealPayloadSize(ctx context.Context, dealProposal cid.Cid, payloadSize uint64) error //perm:write

	IndexerAnnounceAllDeals(ctx context.Context, minerAddr address.Address) error                //perm:admin
	IndexerListMultihashes(ctx context.Context, contextID []byte) ([]multihash.Multihash, error) //perm:read
	IndexerAnnounceLatest(ctx context.Context) (cid.Cid, error)                                  //perm:admin
	IndexerAnnounceLatestHttp(ctx context.Context, urls []string) (cid.Cid, error)               //perm:admin
	IndexerAnnounceDealRemoved(ctx context.Context, contextID []byte) (cid.Cid, error)           //perm:admin
	IndexerAnnounceDeal(ctx context.Context, contextID []byte) (cid.Cid, error)                  //perm:admin

	api.Version
}
