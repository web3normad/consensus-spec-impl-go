package constants


type Slot uint64
type Epoch uint64
type Bytes1 [1]byte
type Bytes4 [4]byte

//Contants declaration
const (
	UINT64_MAX              uint64 = 1<<64 - 1
	UINT64_MAX_SQRT         uint64 = 4294967295
	GENESIS_SLOT            Slot   = 0
	GENESIS_EPOCH           Epoch  = 0
	FAR_FUTURE_EPOCH        Epoch  = 1<<64 - 1
	BASE_REWARDS_PER_EPOCH uint64 = 4
	DEPOSIT_CONTRACT_TREE_DEPTH uint64 = 32
	JUSTIFICATION_BITS_LENGTH uint64 = 4
	ENDIANNESS              string = "little"
)


// Withdrawal Prefixes
var (
	BLS_WITHDRAWAL_PREFIX                Bytes1 = [1]byte{0x00}
	ETH1_ADDRESS_WITHDRAWAL_PREFIX       Bytes1 = [1]byte{0x01}
)
