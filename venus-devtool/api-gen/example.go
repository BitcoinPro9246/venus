package main

import (
	"encoding/json"
	"fmt"
	"go/token"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/filecoin-project/go-state-types/actors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	datatransfer "github.com/filecoin-project/go-data-transfer/v2"
	"github.com/filecoin-project/go-fil-markets/filestore"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/exitcode"
	lminer "github.com/filecoin-project/venus/venus-shared/actors/builtin/miner"
	"github.com/filecoin-project/venus/venus-shared/types/gateway"
	"github.com/filecoin-project/venus/venus-shared/types/market"
	auuid "github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-graphsync"
	textselector "github.com/ipld/go-ipld-selector-text-lite"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/metrics"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-f3/certs"
	"github.com/filecoin-project/go-f3/gpbft"
	"github.com/filecoin-project/go-f3/manifest"
	"github.com/filecoin-project/venus/pkg/constants"
	"github.com/filecoin-project/venus/venus-shared/api/chain"
	"github.com/filecoin-project/venus/venus-shared/types"
	"github.com/filecoin-project/venus/venus-shared/types/market/client"
	"github.com/filecoin-project/venus/venus-shared/types/messager"
	"github.com/filecoin-project/venus/venus-shared/types/wallet"
)

var ExampleValues = map[reflect.Type]interface{}{
	reflect.TypeOf(auth.Permission("")):    auth.Permission("write"),
	reflect.TypeOf(""):                     "string value",
	reflect.TypeOf(market.PieceStatus("")): market.Undefine,
	reflect.TypeOf(uint64(42)):             uint64(42),
	reflect.TypeOf(uint(42)):               uint(42),
	reflect.TypeOf(byte(7)):                byte(7),
	reflect.TypeOf([]byte{}):               []byte("byte array"),
}

func addExample(v interface{}) {
	ExampleValues[reflect.TypeOf(v)] = v
}

func init() {
	c, err := cid.Decode("bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4")
	if err != nil {
		panic(err)
	}

	ExampleValues[reflect.TypeOf(c)] = c

	c2, err := cid.Decode("bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve")
	if err != nil {
		panic(err)
	}

	tsk := types.NewTipSetKey(c, c2)

	ExampleValues[reflect.TypeOf(tsk)] = tsk

	addr, err := address.NewIDAddress(1234)
	if err != nil {
		panic(err)
	}

	ExampleValues[reflect.TypeOf(addr)] = addr

	pid, err := peer.Decode("12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf")
	if err != nil {
		panic(err)
	}
	addExample(pid)
	addExample(&pid)
	uuid, err := types.ParseUUID("e26f1e5c-47f7-4561-a11d-18fab6e748af")
	if err != nil {
		panic(err)
	}
	addExample(constants.TestNetworkVersion)
	addExample(actors.Version6)
	allocationID := types.AllocationId(0)
	addExample(allocationID)
	addExample(&allocationID)
	addExample(map[types.AllocationId]types.Allocation{})
	claimID := types.ClaimId(0)
	addExample(claimID)
	addExample(&claimID)
	addExample(map[types.ClaimId]types.Claim{})
	textSelExample := textselector.Expression("Links/21/Hash/Links/42/Hash")
	clientEvent := retrievalmarket.ClientEventDealAccepted
	addExample(bitfield.NewFromSet([]uint64{5}))
	addExample(abi.RegisteredSealProof_StackedDrg32GiBV1_1)
	addExample(abi.RegisteredPoStProof_StackedDrgWindow32GiBV1)
	addExample(abi.ChainEpoch(10101))
	addExample(crypto.SigTypeBLS)
	addExample(types.KTBLS)
	addExample(types.MTChainMsg)
	addExample(int64(9))
	addExample(12.3)
	addExample(123)
	addExample(uintptr(0))
	addExample(abi.MethodNum(1))
	addExample(exitcode.ExitCode(0))
	addExample(crypto.DomainSeparationTag_ElectionProofProduction)
	addExample(true)
	addExample(abi.UnpaddedPieceSize(1024))
	addExample(abi.UnpaddedPieceSize(1024).Padded())
	addExample(abi.DealID(5432))
	addExample(abi.SectorNumber(9))
	addExample(abi.SectorSize(32 * 1024 * 1024 * 1024))
	addExample(types.MpoolChange(0))
	addExample(network.Connected)
	addExample(types.NetworkName("mainnet"))
	addExample(types.SyncStateStage(1))
	addExample(chain.FullAPIVersion1)
	addExample(types.PCHInbound)
	addExample(time.Minute)
	reqIDBytes, err := uuid.MarshalBinary()
	if err != nil {
		panic(err)
	}
	reqID, err := graphsync.ParseRequestID(reqIDBytes)
	if err != nil {
		panic(err)
	}
	block := blocks.Block(&blocks.BasicBlock{})
	ExampleValues[reflect.TypeOf(&block).Elem()] = block
	addExample(reqID)
	addExample(datatransfer.TransferID(3))
	addExample(datatransfer.Ongoing)
	addExample(clientEvent)
	addExample(&clientEvent)
	addExample(retrievalmarket.ClientEventDealAccepted)
	addExample(retrievalmarket.DealStatusNew)
	addExample(&textSelExample)
	addExample(network.ReachabilityPublic)
	addExample(map[string]int{"name": 42})
	addExample(map[string]time.Time{"name": time.Unix(1615243938, 0).UTC()})
	addExample(map[string]cid.Cid{})
	addExample(&types.ExecutionTrace{
		Msg:    ExampleValue("init", reflect.TypeOf(types.MessageTrace{}), nil).(types.MessageTrace),
		MsgRct: ExampleValue("init", reflect.TypeOf(types.ReturnTrace{}), nil).(types.ReturnTrace),
	})
	addExample(abi.ActorID(1000))
	addExample(map[string]types.Actor{
		"t01236": ExampleValue("init", reflect.TypeOf(types.Actor{}), nil).(types.Actor),
	})
	addExample(json.RawMessage(`"json raw message"`))
	addExample(map[address.Address]*types.Actor{
		addr: {
			Code:    c,
			Head:    c2,
			Nonce:   10,
			Balance: abi.NewTokenAmount(100),
		},
	})
	addExample(map[string]*types.MarketDeal{
		"t026363": ExampleValue("init", reflect.TypeOf(&types.MarketDeal{}), nil).(*types.MarketDeal),
	})
	addExample(map[string]types.MarketBalance{
		"t026363": ExampleValue("init", reflect.TypeOf(types.MarketBalance{}), nil).(types.MarketBalance),
	})
	addExample([]*types.EstimateMessage{
		{
			Msg:  ExampleValue("init", reflect.TypeOf(&types.Message{}), nil).(*types.Message),
			Spec: ExampleValue("init", reflect.TypeOf(&types.MessageSendSpec{}), nil).(*types.MessageSendSpec),
		},
	})
	addExample(map[string]*pubsub.TopicScoreSnapshot{
		"/blocks": {
			TimeInMesh:               time.Minute,
			FirstMessageDeliveries:   122,
			MeshMessageDeliveries:    1234,
			InvalidMessageDeliveries: 3,
		},
	})
	addExample(map[string]metrics.Stats{
		"12D3KooWSXmXLJmBR1M7i9RW9GQPNUhZSzXKzxDHWtAgNuJAbyEJ": {
			RateIn:   100,
			RateOut:  50,
			TotalIn:  174000,
			TotalOut: 12500,
		},
	})
	addExample(map[protocol.ID]metrics.Stats{
		"/fil/hello/1.0.0": {
			RateIn:   100,
			RateOut:  50,
			TotalIn:  174000,
			TotalOut: 12500,
		},
	})
	maddr, err := multiaddr.NewMultiaddr("/ip4/52.36.61.156/tcp/1347/p2p/12D3KooWFETiESTf1v4PGUvtnxMAcEFMzLZbJGg4tjWfGEimYior")
	if err != nil {
		panic(err)
	}
	// because reflect.TypeOf(maddr) returns the concrete type...
	ExampleValues[reflect.TypeOf(struct{ A multiaddr.Multiaddr }{}).Field(0).Type] = maddr
	si := uint64(12)
	addExample(&si)
	addExample(retrievalmarket.DealID(5))
	addExample(abi.ActorID(1000))
	addExample(map[abi.SectorNumber]string{
		123: "can't acquire read lock",
	})
	addExample([]abi.SectorNumber{123, 124})
	addExample(types.CheckStatusCode(0))
	addExample(map[string]interface{}{"abc": 123})
	addExample(types.HCApply)
	addExample(lminer.SectorOnChainInfoFlags(0))

	// messager
	i64 := int64(10000)
	addExample(uuid)
	addExample(messager.OnChainMsg)
	addExample(messager.AddressStateAlive)
	addExample(&i64)
	addExample(ExampleValue("init", reflect.TypeOf(&messager.Address{}), nil).(*messager.Address))
	addExample(&messager.Node{
		ID:    uuid,
		Name:  "venus",
		URL:   "/ip4/127.0.0.1/tcp/3453",
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0._eHBJJAiBzQmfcbD_vVmtTrkgyJQ-LOgGOiHfb8rU1I",
		Type:  messager.LightNode,
	})
	addExample(&messager.Message{
		ID:          uuid.String(),
		UnsignedCid: &c,
		SignedCid:   &c,
		Message:     ExampleValue("init", reflect.TypeOf(types.Message{}), nil).(types.Message),
		Signature:   ExampleValue("init", reflect.TypeOf(&crypto.Signature{}), nil).(*crypto.Signature),
		Height:      100,
		Confidence:  10,
		Receipt:     ExampleValue("init", reflect.TypeOf(&types.MessageReceipt{}), nil).(*types.MessageReceipt),
		TipSetKey:   tsk,
		Meta:        ExampleValue("init", reflect.TypeOf(&messager.SendSpec{}), nil).(*messager.SendSpec),
		WalletName:  "test",
		State:       messager.UnFillMsg,
	})
	addExample(ExampleValue("init", reflect.TypeOf(&messager.SendSpec{}), nil).(*messager.SendSpec))
	addExample(messager.QuickSendParamsCodecJSON)

	// wallet
	addExample(wallet.MEChainMsg)

	// used in gateway
	addExample(types.PaddedByteIndex(10))

	// used in market
	addExample(filestore.Path("/some/path"))

	clientDataSelector := client.DataSelector("/ipld/a/b/c")
	addExample(clientDataSelector)
	addExample(&clientDataSelector)

	addExample(client.ImportID(1234))

	uuidTmp := auuid.MustParse("102334ec-35a3-4b36-be9f-02883844503a")
	addExample(&uuidTmp)

	addExample(market.DirectDealState(1))

	// eth types
	ethint := types.EthUint64(5)
	addExample(ethint)
	addExample(&ethint)
	ethaddr, _ := types.ParseEthAddress("0x5CbEeCF99d3fDB3f25E309Cc264f240bb0664031")
	addExample(&ethaddr)
	ethhash, _ := types.EthHashFromCid(c)
	addExample(&ethhash)
	ethFeeHistoryReward := [][]types.EthBigInt{}
	addExample(&ethFeeHistoryReward)

	filterid := types.EthFilterID(ethhash)
	addExample(filterid)
	addExample(&filterid)

	subID := types.EthSubscriptionID(ethhash)
	addExample(subID)
	addExample(&subID)

	pstring := func(s string) *string { return &s }
	addExample(&types.EthFilterSpec{
		FromBlock: pstring("2301220"),
		Address:   []types.EthAddress{ethaddr},
	})

	addExample(&types.ActorEventBlock{
		Codec: 0x51,
		Value: []byte("ddata"),
	})

	addExample(&types.ActorEventFilter{
		Addresses: []address.Address{addr},
		Fields: map[string][]types.ActorEventBlock{
			"abc": {
				{
					Codec: 0x51,
					Value: []byte("ddata"),
				},
			},
		},
		FromHeight: epochPtr(1010),
		ToHeight:   epochPtr(1020),
	})

	percent := types.Percent(123)
	addExample(percent)
	addExample(&percent)
	addExample(gateway.UnsealStateFinished)
	addExample(types.UnpaddedByteIndex(0))

	addExample(retrievalmarket.CborGenCompatibleNode{})
	addExample(gateway.HostNode)
	addExample(&certs.FinalityCertificate{})
	addExample(gpbft.ActorID(1000))

	addExample(&manifest.Manifest{})
	addExample(gpbft.NetworkName("filecoin"))
	addExample(gpbft.INITIAL_PHASE)

	after := types.EthUint64(0)
	count := types.EthUint64(100)

	ethTraceFilterCriteria := types.EthTraceFilterCriteria{
		FromBlock:   pstring("latest"),
		ToBlock:     pstring("latest"),
		FromAddress: types.EthAddressList{ethaddr},
		ToAddress:   types.EthAddressList{ethaddr},
		After:       &after,
		Count:       &count,
	}
	addExample(&ethTraceFilterCriteria)
	addExample(ethTraceFilterCriteria)

	f3Lease := types.F3ParticipationLease{
		Network:      "filecoin",
		Issuer:       pid.String(),
		MinerID:      1234,
		FromInstance: 10,
		ValidityTerm: 15,
	}
	addExample(f3Lease)
	addExample(&f3Lease)
	ecchain := &gpbft.ECChain{
		TipSets: []*gpbft.TipSet{
			{
				Epoch:      0,
				Key:        tsk.Bytes(),
				PowerTable: c,
			},
		},
	}
	f3Cert := certs.FinalityCertificate{
		GPBFTInstance: 0,
		ECChain:       ecchain,
		SupplementalData: gpbft.SupplementalData{
			PowerTable: c,
		},
		Signers:   bitfield.NewFromSet([]uint64{2, 3, 5, 7}),
		Signature: []byte("UnDadaSeA"),
		PowerTableDelta: []certs.PowerTableDelta{
			{
				ParticipantID: 0,
				PowerDelta:    gpbft.StoragePower{},
				SigningKey:    []byte("BaRrelEYe"),
			},
		},
	}
	addExample(f3Cert)
	addExample(&f3Cert)
	addExample(gpbft.InstanceProgress{
		Instant: gpbft.Instant{
			ID:    1413,
			Round: 1,
			Phase: gpbft.COMMIT_PHASE,
		},
		Input: ecchain,
	})
}

func ExampleValue(method string, t, parent reflect.Type) interface{} {
	v, ok := ExampleValues[t]
	if ok {
		return v
	}

	switch t.Kind() {
	case reflect.Slice:
		out := reflect.New(t).Elem()
		out = reflect.Append(out, reflect.ValueOf(ExampleValue(method, t.Elem(), t)))
		return out.Interface()
	case reflect.Chan:
		return ExampleValue(method, t.Elem(), nil)
	case reflect.Struct:
		es := exampleStruct(method, t, parent)
		v := reflect.ValueOf(es).Interface()
		ExampleValues[t] = v
		return v
	case reflect.Array:
		out := reflect.New(t).Elem()
		for i := 0; i < t.Len(); i++ {
			out.Index(i).Set(reflect.ValueOf(ExampleValue(method, t.Elem(), t)))
		}
		return out.Interface()
	case reflect.Map:
		out := reflect.MakeMap(t)
		out.SetMapIndex(reflect.ValueOf(ExampleValue(method, t.Key(), parent)), reflect.ValueOf(ExampleValue(method, t.Elem(), parent)))
		return out.Interface()
	case reflect.Ptr:
		out := reflect.New(t.Elem())
		out.Elem().Set(reflect.ValueOf(ExampleValue(method, t.Elem(), t)))
		return out.Interface()

	case reflect.Interface:
		if t.Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			return fmt.Errorf("empty error")
		}
		return struct{}{}
	}

	_, _ = fmt.Fprintf(os.Stderr, "Warnning: No example value for type: %s (method '%s')\n", t, method)
	return nil
}

func exampleStruct(method string, t, parent reflect.Type) interface{} {
	ns := reflect.New(t).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if shouldIgnoreField(f, parent) {
			continue
		}

		if strings.Title(f.Name) == f.Name {
			ns.Field(i).Set(reflect.ValueOf(ExampleValue(method, f.Type, t)))
		}
	}

	return ns.Interface()
}

func shouldIgnoreField(f reflect.StructField, parentType reflect.Type) bool {
	if f.Type == parentType {
		return true
	}

	if len(f.Name) == 0 {
		return true
	}

	if !token.IsExported(f.Name) {
		return true
	}

	jtag := f.Tag.Get("json")
	if len(jtag) == 0 {
		return false
	}

	return strings.Split(jtag, ",")[0] == "-"
}

func epochPtr(ei int64) *abi.ChainEpoch {
	ep := abi.ChainEpoch(ei)
	return &ep
}
