// Copyright 2015 The go-ethereum Authors
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

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://a484ac1e9ba411d3f2ed94c730342b199a20ae61148cd5272f0d192d716f2345829f7b5fe37c2c5fe10b353749e4bf80df576caea2efe5696110516943efdecc@157.90.164.232:30303",
	"enode://917a1b527b8fc1f7858471e0ab6292ca0169f35735de54a4717192176b326bffecf9444e7c7f47e4c1f2f6597ff71a4c7d36208406b2af1cac46430046891ab2@138.201.116.212:30303",
	"enode://dbc78d4dc93b5f907195fbfd9217a1cd6ab0f3ad12cf349bc134624cc70e8c78121586de0809ec5478c4c204fab508f6954221bbfdccb7d230f7635e503c49c0@37.27.28.0:30303",
	"enode://8db54b7ec2d5df88acf0217039594d8e6e2f68668521ebe2485af4fd3f37df99e0836d789ae0969e9d7e601fbe4ae5ed83fda41f1f2e032625f73690bd75bf94@31.220.102.203:30303",
	"enode://99120219b0c5cc34caa0e65b359e29ce0f6d3b7e2bc8f94bf59d892e6934a02bddec6f45e43f428fe92585b6573f2f3c1747d9021663300208a69c7b986e4fc4@94.72.121.31:30303",
	"enode://2e7bce2e75a6bd8258568a8a0768de8bd6a78ee31eac966a90b0959173704c509577b0224feb24fa02704b3b65527b3f34aad772c58c06b71f0d580bec3e4c58@109.123.234.55:30303",
	"enode://9424eab9c95b3584117e22902dabd880c2a2fbfee773404d0f9cfb5d87ffeafa67f3b368f43eba3fac306ab9299068227480ed50a144392ff335bbedb20536ff@84.247.158.2:30303",
	"enode://a997259bbb6076de2e967f90a9056b56f8d9e37911c502f366fbf14f62226570b106d36f9415f872e3037f95c1e6f051ca9dd3ed3b8b2c44fdda3547a3f03dac@46.250.243.161:30303",
}

// AuroriaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Auroria test network.
var AuroriaBootnodes = []string{
	"enode://46d68d753537d4396564199dc348dc9293a5926198cb83ee27beba3d2b1a60cc682c63114bec77686fb81a49fe1c3f9782374935636613fdd6b6f0874ed657bc@207.154.197.204:30303",
	"enode://e9e2beba4d197733a435a4b5827d032201f9511a6aceccbaf729dd5cb9f1007d9b5866eea514a007fe54bf94164783d344caeb44f810ae000ff4e95d455fa98d@144.126.200.138:30303",
	"enode://f8aab97719acebf6a5c99aa6f28d7d5cd4a25cbe0b183986872f78a9a140f27b63605d506c29a751ed2a7832250b21ebd3e6b7f8dece510668ef82ba116fbfc5@128.199.63.227:30303",
	"enode://dc9e5e5cf936a7b5bccb86b894e35141bbff9bee84840ad65baae6f51bf1a01122b5602772154ae3dc57fb51dadbf3a146705d44e7991bf84f078f4701e43cb0@138.197.124.43:30303",
}

var V5Bootnodes = []string{}

// const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
	// var net string
	// switch genesis {
	// case MainnetGenesisHash:
	// 	net = "mainnet"
	// case AuroriaGenesisHash:
	// 	net = "auroria"
	// default:
	// 	return ""
	// }
	// return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
