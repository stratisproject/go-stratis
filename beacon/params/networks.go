// Copyright 2024 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/common"
)

// //go:embed checkpoint_mainnet.hex
// var checkpointMainnet string

var (
	MainnetLightConfig = (&ChainConfig{
		GenesisValidatorsRoot: common.HexToHash("0xa73c6c40923a73d0ba772eb3791352c8f6cf42bd72c4677e9153d5a14de991e5"),
		GenesisTime:           1710353050,
		Checkpoint:            common.HexToHash(""),
	}).
		AddFork("GENESIS", 0, []byte{10, 0, 0, 0}).
		AddFork("ALTAIR", 0, []byte{11, 0, 0, 0}).
		AddFork("BELLATRIX", 0, []byte{12, 0, 0, 0}).
		AddFork("CAPELLA", 0, []byte{13, 0, 0, 0}).
		AddFork("DENEB", 0, []byte{14, 0, 0, 0}).
		AddFork("ELECTRA", 18446744073709551615, []byte{15, 0, 0, 0})

	AuroriaLightConfig = (&ChainConfig{
		GenesisValidatorsRoot: common.HexToHash("0xe50761b51cdc83be5f6f20114142d38fce566e1caf210150ab5f6b4b82845a4d"),
		GenesisTime:           1741097967,
		Checkpoint:            common.HexToHash(""),
	}).
		AddFork("GENESIS", 0, []byte{10, 0, 10, 20}).
		AddFork("ALTAIR", 0, []byte{11, 0, 10, 20}).
		AddFork("BELLATRIX", 0, []byte{12, 0, 10, 20}).
		AddFork("CAPELLA", 0, []byte{13, 0, 10, 20}).
		AddFork("DENEB", 0, []byte{14, 0, 10, 20}).
		AddFork("ELECTRA", 18446744073709551615, []byte{15, 0, 10, 20})
)
