// FETCHED FROM LOTUS: builtin/multisig/message.go.template

package multisig

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	actorstypes "github.com/filecoin-project/go-state-types/actors"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	init16 "github.com/filecoin-project/go-state-types/builtin/v16/init"
	"github.com/filecoin-project/go-state-types/manifest"

	builtintypes "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/venus/venus-shared/actors"
	init_ "github.com/filecoin-project/venus/venus-shared/actors/builtin/init"
	"github.com/filecoin-project/venus/venus-shared/actors/types"
)

type message15 struct{ message0 }

func (m message15) Create(
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {

	lenAddrs := uint64(len(signers))

	if lenAddrs < threshold {
		return nil, fmt.Errorf("cannot require signing of more addresses than provided for multisig")
	}

	if threshold == 0 {
		threshold = lenAddrs
	}

	if m.from == address.Undef {
		return nil, fmt.Errorf("must provide source address")
	}

	// Set up constructor parameters for multisig
	msigParams := &multisig15.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,
		StartEpoch:            unlockStart,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {
		return nil, actErr
	}

	code, ok := actors.GetActorCodeID(actorstypes.Version15, manifest.MultisigKey)
	if !ok {
		return nil, fmt.Errorf("failed to get multisig code ID")
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init16.ExecParams{
		CodeCID:           code,
		ConstructorParams: enc,
	}

	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Method: builtintypes.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}
