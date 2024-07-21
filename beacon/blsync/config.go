// Copyright 2022 The go-ethereum Authors
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

package blsync

import (
	"github.com/ethereum/go-ethereum/beacon/types"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/urfave/cli/v2"
)

// lightClientConfig contains beacon light client configuration
type lightClientConfig struct {
	*types.ChainConfig
	Checkpoint common.Hash
}

var (
	MainnetConfig = lightClientConfig{
		ChainConfig: (&types.ChainConfig{
			GenesisValidatorsRoot: common.HexToHash("0xa73c6c40923a73d0ba772eb3791352c8f6cf42bd72c4677e9153d5a14de991e5"),
			GenesisTime:           1710353050,
		}).
			AddFork("GENESIS", 0, []byte{10, 0, 0, 0}).
			AddFork("ALTAIR", 0, []byte{11, 0, 0, 0}).
			AddFork("BELLATRIX", 0, []byte{12, 0, 0, 0}).
			AddFork("CAPELLA", 0, []byte{13, 0, 0, 0}).
			AddFork("DENEB", 0, []byte{14, 0, 0, 0}),
		Checkpoint: common.HexToHash("0xb2f98efe7340a8e325cebb8797c7f78f12f814448b996fa51332db81de169c4c"),
	}

	AuroriaConfig = lightClientConfig{
		ChainConfig: (&types.ChainConfig{
			GenesisValidatorsRoot: common.HexToHash("0x0d2139e6c8e64a17096a416e2850a95e944129d32017a0e790072ed94a24e10e"),
			GenesisTime:           1707830378,
		}).
			AddFork("GENESIS", 0, []byte{10, 0, 10, 20}).
			AddFork("ALTAIR", 0, []byte{11, 0, 10, 20}).
			AddFork("BELLATRIX", 0, []byte{12, 0, 10, 20}).
			AddFork("CAPELLA", 0, []byte{13, 0, 10, 20}).
			AddFork("DENEB", 0, []byte{14, 0, 10, 20}),
		Checkpoint: common.HexToHash("0x6f7bd1a7b18d22f70b432fdea55b0549b8e710f57e28e08c10181ac17aa425b1"),
	}
)

func makeChainConfig(ctx *cli.Context) lightClientConfig {
	var config lightClientConfig
	customConfig := ctx.IsSet(utils.BeaconConfigFlag.Name)
	utils.CheckExclusive(ctx, utils.MainnetFlag, utils.AuroriaFlag, utils.BeaconConfigFlag)
	switch {
	case ctx.Bool(utils.MainnetFlag.Name):
		config = MainnetConfig
	case ctx.Bool(utils.AuroriaFlag.Name):
		config = AuroriaConfig
	default:
		if !customConfig {
			config = MainnetConfig
		}
	}
	// Genesis root and time should always be specified together with custom chain config
	if customConfig {
		if !ctx.IsSet(utils.BeaconGenesisRootFlag.Name) {
			utils.Fatalf("Custom beacon chain config is specified but genesis root is missing")
		}
		if !ctx.IsSet(utils.BeaconGenesisTimeFlag.Name) {
			utils.Fatalf("Custom beacon chain config is specified but genesis time is missing")
		}
		if !ctx.IsSet(utils.BeaconCheckpointFlag.Name) {
			utils.Fatalf("Custom beacon chain config is specified but checkpoint is missing")
		}
		config.ChainConfig = &types.ChainConfig{
			GenesisTime: ctx.Uint64(utils.BeaconGenesisTimeFlag.Name),
		}
		if c, err := hexutil.Decode(ctx.String(utils.BeaconGenesisRootFlag.Name)); err == nil && len(c) <= 32 {
			copy(config.GenesisValidatorsRoot[:len(c)], c)
		} else {
			utils.Fatalf("Invalid hex string", "beacon.genesis.gvroot", ctx.String(utils.BeaconGenesisRootFlag.Name), "error", err)
		}
		if err := config.ChainConfig.LoadForks(ctx.String(utils.BeaconConfigFlag.Name)); err != nil {
			utils.Fatalf("Could not load beacon chain config file", "file name", ctx.String(utils.BeaconConfigFlag.Name), "error", err)
		}
	} else {
		if ctx.IsSet(utils.BeaconGenesisRootFlag.Name) {
			utils.Fatalf("Genesis root is specified but custom beacon chain config is missing")
		}
		if ctx.IsSet(utils.BeaconGenesisTimeFlag.Name) {
			utils.Fatalf("Genesis time is specified but custom beacon chain config is missing")
		}
	}
	// Checkpoint is required with custom chain config and is optional with pre-defined config
	if ctx.IsSet(utils.BeaconCheckpointFlag.Name) {
		if c, err := hexutil.Decode(ctx.String(utils.BeaconCheckpointFlag.Name)); err == nil && len(c) <= 32 {
			copy(config.Checkpoint[:len(c)], c)
		} else {
			utils.Fatalf("Invalid hex string", "beacon.checkpoint", ctx.String(utils.BeaconCheckpointFlag.Name), "error", err)
		}
	}
	return config
}
