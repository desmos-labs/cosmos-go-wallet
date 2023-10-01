# Version 0.7.2
## Bug fixes
- Fixed a bug in the fee amount computation

# Version 0.7.1
## Bug fixes
- Fixed a bug in the gas simulation process

# Version 0.7.0
## Features
- Added the ability to specify a sequence number when building a transaction 
- Updated the `NewTransactionData` to accept a single message slice param instead of two separate params

## Bug fixes
- Fixed a typo inside an error message (thanks [@giorgionocera](https://github.com/giorgionocera))

## Dependencies
- Updated Cosmos SDK to `v0.47.3`
- Removed the dependency from the `cosmos-sdk/simapp` package

# Version 0.6.0
## Dependencies
- Updated Cosmos SDK to `v0.47.2`

# Version 0.5.1
## Bug fixes
- Fixed a bug in the gas and fees simulation process

# Version 0.5.0
## Features
- Set min TLS version to `1.2` when connecting to gRPC endpoints
- Removed `Wallet#SetGasPerMessage`. `TransactionData#WithFeeAuto` and `TransactionData#WithGasAuto` should be used instead

# Version 0.4.0
## Features
- Added `WithGasAuto` and `WithFeeAuto` to `TransactionData` in order to allow auto computing the gas and fee amount by simulating the transaction

# Version 0.3.0
## Dependencies
- Updated Cosmos SDK to `v0.45.7`

# Version 0.2.0
## Dependencies
- Updated Cosmos SDK to `v0.45.6`

# Version 0.1.0
This version is compatible with Cosmos SDK `v0.44.5`.