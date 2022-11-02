package stateutil

import (
	"math/big"

	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v5/config/params"
	"github.com/prysmaticlabs/prysm/v5/consensus-types/primitives"
)

// UnrealizedCheckpointBalances returns the total current active balance, the
// total previous epoch correctly attested for target balance, and the total
// current epoch correctly attested for target balance. It takes the current and
// previous epoch participation bits as parameters so implicitly only works for
// beacon states post-Altair.
func UnrealizedCheckpointBalances(cp, pp []byte, validators ValReader, currentEpoch primitives.Epoch) (*big.Int, *big.Int, *big.Int, error) {
	targetIdx := params.BeaconConfig().TimelyTargetFlagIndex
	activeBalance := big.NewInt(0)
	currentTarget := big.NewInt(0)
	prevTarget := big.NewInt(0)
	if len(cp) < validators.Len() || len(pp) < validators.Len() {
		return activeBalance, currentTarget, prevTarget, errors.New("participation does not match validator set")
	}

	valLength := validators.Len()
	for i := 0; i < valLength; i++ {
		v, err := validators.At(i)
		if err != nil {
			return activeBalance, currentTarget, prevTarget, err
		}
		activeCurrent := v.ActivationEpoch <= currentEpoch && currentEpoch < v.ExitEpoch
		effectiveBalance := new(big.Int).SetUint64(v.EffectiveBalance)
		if activeCurrent {
			activeBalance.Add(activeBalance, effectiveBalance)
		}
		if v.Slashed {
			continue
		}
		if activeCurrent && ((cp[i]>>targetIdx)&1) == 1 {
			currentTarget.Add(currentTarget, effectiveBalance)
		}
		activePrevious := v.ActivationEpoch < currentEpoch && currentEpoch <= v.ExitEpoch
		if activePrevious && ((pp[i]>>targetIdx)&1) == 1 {
			prevTarget.Add(prevTarget, effectiveBalance)
		}
	}
	activeBalance, prevTarget, currentTarget = ensureLowerBound(activeBalance, prevTarget, currentTarget)
	return activeBalance, prevTarget, currentTarget, nil
}

func ensureLowerBound(activeCurrEpoch, prevTargetAttested, currTargetAttested *big.Int) (*big.Int, *big.Int, *big.Int) {
	ebi := new(big.Int).SetUint64(params.BeaconConfig().EffectiveBalanceIncrement)
	if ebi.Cmp(activeCurrEpoch) > 0 {
		activeCurrEpoch = ebi
	}
	if ebi.Cmp(prevTargetAttested) > 0 {
		prevTargetAttested = ebi
	}
	if ebi.Cmp(currTargetAttested) > 0 {
		currTargetAttested = ebi
	}
	return activeCurrEpoch, prevTargetAttested, currTargetAttested
}
