PhaseZero Consensus Specification in Go (Test Implementation)
This is an experimental implementation and testing project for the PhaseZero consensus specification of Ethereum 2.0 in Go. The focus of this project is to define and test the constants, custom types, and core components essential to the consensus protocol, such as slots, epochs, validator indices, and more.

Key aspects of the project:

Testing constants (e.g., GENESIS_SLOT, BASE_REWARDS_PER_EPOCH) and their expected behavior in Go.
Defining and testing custom types like Slot, Epoch, ValidatorIndex, and cryptographic data types.
Implementing logic for validator rewards, penalties, and inactivity penalties.
Structuring the code for easy expansion towards a full beacon chain client.
This project serves as an initial step in building a beacon chain client and ensures that the core constants and data structures in the PhaseZero consensus spec work as expected in Go.
