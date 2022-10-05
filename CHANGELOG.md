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