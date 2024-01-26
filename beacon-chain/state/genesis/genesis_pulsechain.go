//go:build !noMainnetGenesis
// +build !noMainnetGenesis

package genesis

import (
	_ "embed"

	"github.com/prysmaticlabs/prysm/v5/config/params"
)

var (
	//go:embed pulsechain.ssz.snappy
	pulseChainRawSSZCompressed []byte // 466Kb
)

func init() {
	embeddedStates[params.PulseChainName] = &pulseChainRawSSZCompressed
}
