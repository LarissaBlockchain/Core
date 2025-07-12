package misc

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var (
	// AddressBlockingForkBlock Define the hard fork block number when blocking becomes active
	AddressBlockingForkBlock = uint64(3672246) // Set your fork block number

	// BlockedAddresses List of blocked addresses
	BlockedAddresses = map[common.Address]bool{
		common.HexToAddress("0x5CcCcb6d334197c7C4ba94E7873d0ef11381CD4e"): true, // Xeggex Address 1
		common.HexToAddress("0x4648451b5F87FF8F0F7D622bD40574bb97E25980"): true, // Xeggex Address 2
	}
)

// ValidateTransaction checks if transaction involves blocked addresses
func ValidateTransaction(tx *types.Transaction, blockNumber uint64, config *params.ChainConfig) error {
	// Only apply blocking after the fork block
	if blockNumber < AddressBlockingForkBlock {
		return nil
	}

	// Check if sender is blocked
	if tx.To() != nil {
		if IsAddressBlocked(*tx.To()) {
			return errors.New("transaction to blocked address rejected")
		}
	}

	// Note: Sender validation is handled in state transition
	return nil
}

// IsAddressBlocked checks if an address is in the blocked list
func IsAddressBlocked(addr common.Address) bool {
	return BlockedAddresses[addr]
}
