package misc

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	// ErrBadProStratisMasterNodeExtra is returned if a header doesn't support the StratisMasterNode fork on a
	// pro-fork client.
	ErrBadProStratisMasterNodeExtra = errors.New("bad StratisMasterNode pro-fork extra-data")

	// ErrBadNoStratisMasterNodeExtra is returned if a header does support the StratisMasterNode fork on a no-
	// fork client.
	ErrBadNoStratisMasterNodeExtra = errors.New("bad StratisMasterNode no-fork extra-data")
)

// VerifyStratisMasterNodeHeaderExtraData validates the extra-data field of a block header to
// ensure it conforms to StratisMasterNode hard-fork rules.
//
// StratisMasterNode hard-fork extension to the header validity:
//
//   - if the node is no-fork, do not accept blocks in the [fork, fork+10) range
//     with the fork specific extra-data set.
//   - if the node is pro-fork, require blocks in the specific range to have the
//     unique extra-data set.
func VerifyStratisMasterNodeHeaderExtraData(config *params.ChainConfig, header *types.Header) error {
	// Short circuit validation if the node doesn't care about the StratisMaster fork
	if config.StratisMasterNodeForkBlock == nil {
		return nil
	}
	// Make sure the block is within the fork's modified extra-data range
	limit := new(big.Int).Add(config.StratisMasterNodeForkBlock, params.StratisMasterNodeExtraRange)
	if header.Number.Cmp(config.StratisMasterNodeForkBlock) < 0 || header.Number.Cmp(limit) >= 0 {
		return nil
	}
	// Depending on whether we support or oppose the fork, validate the extra-data contents
	if config.StratisMasterNodeForkSupport {
		if !bytes.Equal(header.Extra, params.StratisMasterNodeBlockExtra) {
			return ErrBadProStratisMasterNodeExtra
		}
	} else {
		if bytes.Equal(header.Extra, params.StratisMasterNodeBlockExtra) {
			return ErrBadNoStratisMasterNodeExtra
		}
	}
	// All ok, header has the same extra-data we expect
	return nil
}

func VerifyStratisMasterNodeV2HeaderExtraData(config *params.ChainConfig, header *types.Header) error {
	// Short circuit validation if the node doesn't care about the StratisMaster fork V2
	if config.StratisMasterNodeForkV2Block == nil {
		return nil
	}
	// Make sure the block is within the fork's modified extra-data range
	limit := new(big.Int).Add(config.StratisMasterNodeForkV2Block, params.StratisMasterNodeExtraRange)
	if header.Number.Cmp(config.StratisMasterNodeForkV2Block) < 0 || header.Number.Cmp(limit) >= 0 {
		return nil
	}
	// Depending on whether we support or oppose the fork, validate the extra-data contents
	if config.StratisMasterNodeForkV2Support {
		if !bytes.Equal(header.Extra, params.StratisMasterNodeV2BlockExtra) {
			return ErrBadProStratisMasterNodeExtra
		}
	} else {
		if bytes.Equal(header.Extra, params.StratisMasterNodeV2BlockExtra) {
			return ErrBadNoStratisMasterNodeExtra
		}
	}
	// All ok, header has the same extra-data we expect
	return nil
}

// ApplyStratisMasterNodeHardFork modifies Stratis MasterNode contract code
func ApplyStratisMasterNodeHardFork(chainID *big.Int, statedb *state.StateDB) {
	// Replace MasterNode contract bytecode
	statedb.SetCode(params.StratisMasterNodeContract, hexutil.MustDecode(params.StratisMasterNodeContractBytecode))
}

// ApplyStratisMasterNodeHardForkV2 modifies Stratis MasterNode contract code
func ApplyStratisMasterNodeHardForkV2(chainID *big.Int, statedb *state.StateDB) {
	// Replace MasterNode contract bytecode
	statedb.SetCode(params.StratisMasterNodeContract, hexutil.MustDecode(params.StratisMasterNodeV2ContractBytecode))
}
