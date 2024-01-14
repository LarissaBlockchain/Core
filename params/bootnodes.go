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
	"enode://41a8f8da1af713cb8bff1502db493ed92f67b2a6de4223b4ae851ff0d942aa283cbdc6be028574d4a0a0cb7ec7e5c1d037bbc93014359132838e6d648c02a1bc@rpc.larissa.network:30303",
	"enode://9d4b32ae9462e40e37e171caa7d498a4cc834a1432b94b31a5c7e9ae4ee1c2ad68763d103762f611ecf44635ac9b79fb1787de1f42c270750e6ec60f6480377a@pool.larissa.network:30303",
}

// NeptuneBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Neptune test network.
var NeptuneBootnodes = []string{
	"enode://882c6ba3d51b578ce497b2946769a961fa5f763d037f43558aeba97447a609eee95d3a37e3cc5d6d74b5e5acce49d683084df3708f56075b32db1cc663226828@rpc.larissa.network:30303",
	"enode://22beef86a2d544386ae7ad8221a30d55a2d93508b542b916c79524e8840099fe33fde7cc9985037e5a64be4a7969678cf26b26e33d6556dadc2c1f6b3f783624@pool.larissa.network:30303",
}

var V5Bootnodes = []string{}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case NeptuneGenesisHash:
		net = "neptune"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
