package types

type Slot uint64
type Epoch uint64
type CommitteeIndex uint64
type ValidatorIndex uint64
type Gwei uint64

// Fixed-size byte array definitions
type Bytes1 [1]byte
type Bytes4 [4]byte
type Bytes32 [32]byte
type Bytes48 [48]byte
type Bytes96 [96]byte

// Types mapped to their fixed-size arrays
type Root Bytes32
type Hash32 Bytes32
type Version Bytes4
type DomainType Bytes4
type ForkDigest Bytes4
type Domain Bytes32
type BLSPubkey Bytes48
type BLSSignature Bytes96



