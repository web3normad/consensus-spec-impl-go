package main

import (
	"fmt"
	"consensus-impl-go/phaseZero/constants" 
	"consensus-impl-go/phaseZero/types"    
)



func main() {
	var slot types.Slot = 29219
	var root types.Root
	var signature types.BLSSignature

	// Example: Assigning values
	copy(root[:], []byte("example_merkle_root_256_bit"))
	copy(signature[:], []byte("example_bls_signature_96_bytes"))

	fmt.Printf("Slot: %d\n", slot)
	fmt.Printf("Root: %x\n", root)
	fmt.Printf("Signature: %x\n", signature)

	// Using constants from the constants package
	fmt.Printf("UINT64_MAX: %d\n", constants.UINT64_MAX)
    fmt.Printf("UINT64_MAX_SQRT: %d\n", constants.UINT64_MAX_SQRT)
    fmt.Printf("GENESIS_SLOT: %d\n", constants.GENESIS_SLOT)
    fmt.Printf("GENESIS_EPOCH: %d\n", constants.GENESIS_EPOCH)
    fmt.Printf("FAR_FUTURE_EPOCH: %d\n", constants.FAR_FUTURE_EPOCH)
    fmt.Printf("BASE_REWARDS_PER_EPOCH: %d\n", constants.BASE_REWARDS_PER_EPOCH)
    fmt.Printf("DEPOSIT_CONTRACT_TREE_DEPTH: %d\n", constants.DEPOSIT_CONTRACT_TREE_DEPTH)
    fmt.Printf("JUSTIFICATION_BITS_LENGTH: %d\n", constants.JUSTIFICATION_BITS_LENGTH)
    fmt.Printf("ENDIANNESS: %s\n", constants.ENDIANNESS)
	fmt.Printf("BLS_WITHDRAWAL_PREFIX : %s\n", constants.BLS_WITHDRAWAL_PREFIX)
	fmt.Printf("ETH1_ADDRESS_WITHDRAWAL_PREFIX  : %s\n", constants.ETH1_ADDRESS_WITHDRAWAL_PREFIX )
}
