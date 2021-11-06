package types

type ChainConfig struct {
	RPCAddr  string `toml:"rpc_addr"`
	GRPCAddr string `toml:"grpc_addr"`
	GasPrice string `toml:"gas_price"`
}

type AccountConfig struct {
	Bech32Prefix string `toml:"bech32_prefix"`
	Mnemonic     string `toml:"mnemonic"`
	HDPath       string `toml:"hd_path"`
}
