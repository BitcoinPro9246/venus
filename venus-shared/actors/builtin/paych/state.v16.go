// FETCHED FROM LOTUS: builtin/paych/state.go.template

package paych

import (
	"fmt"

	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/venus/venus-shared/actors"
	"github.com/filecoin-project/venus/venus-shared/actors/adt"

	paych16 "github.com/filecoin-project/go-state-types/builtin/v16/paych"
	adt16 "github.com/filecoin-project/go-state-types/builtin/v16/util/adt"
)

var _ State = (*state16)(nil)

func load16(store adt.Store, root cid.Cid) (State, error) {
	out := state16{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make16(store adt.Store) (State, error) {
	out := state16{store: store}
	out.State = paych16.State{}
	return &out, nil
}

type state16 struct {
	paych16.State
	store adt.Store
	lsAmt *adt16.Array
}

// Channel owner, who has funded the actor
func (s *state16) From() (address.Address, error) {
	return s.State.From, nil
}

// Recipient of payouts from channel
func (s *state16) To() (address.Address, error) {
	return s.State.To, nil
}

// Height at which the channel can be `Collected`
func (s *state16) SettlingAt() (abi.ChainEpoch, error) {
	return s.State.SettlingAt, nil
}

// Amount successfully redeemed through the payment channel, paid out on `Collect()`
func (s *state16) ToSend() (abi.TokenAmount, error) {
	return s.State.ToSend, nil
}

func (s *state16) getOrLoadLsAmt() (*adt16.Array, error) {
	if s.lsAmt != nil {
		return s.lsAmt, nil
	}

	// Get the lane state from the chain
	lsamt, err := adt16.AsArray(s.store, s.State.LaneStates, paych16.LaneStatesAmtBitwidth)
	if err != nil {
		return nil, err
	}

	s.lsAmt = lsamt
	return lsamt, nil
}

// Get total number of lanes
func (s *state16) LaneCount() (uint64, error) {
	lsamt, err := s.getOrLoadLsAmt()
	if err != nil {
		return 0, err
	}
	return lsamt.Length(), nil
}

func (s *state16) GetState() interface{} {
	return &s.State
}

// Iterate lane states
func (s *state16) ForEachLaneState(cb func(idx uint64, dl LaneState) error) error {
	// Get the lane state from the chain
	lsamt, err := s.getOrLoadLsAmt()
	if err != nil {
		return err
	}

	// Note: we use a map instead of an array to store laneStates because the
	// client sets the lane ID (the index) and potentially they could use a
	// very large index.
	var ls paych16.LaneState
	return lsamt.ForEach(&ls, func(i int64) error {
		return cb(uint64(i), &laneState16{ls})
	})
}

type laneState16 struct {
	paych16.LaneState
}

func (ls *laneState16) Redeemed() (big.Int, error) {
	return ls.LaneState.Redeemed, nil
}

func (ls *laneState16) Nonce() (uint64, error) {
	return ls.LaneState.Nonce, nil
}

func (s *state16) ActorKey() string {
	return manifest.PaychKey
}

func (s *state16) ActorVersion() actorstypes.Version {
	return actorstypes.Version16
}

func (s *state16) Code() cid.Cid {
	code, ok := actors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}
