package types

type ChainConfig struct {
	Bech32Prefix string `toml:"bech32_prefix" yaml:"bech32_prefix"`
	RPCAddr      string `toml:"rpc_addr" yaml:"rpc_addr"`
	GRPCAddr     string `toml:"grpc_addr" yaml:"grpc_addr"`
	GasPrice     string `toml:"gas_price" yaml:"gas_price"`
}

type AccountConfig struct {
	Mnemonic string `toml:"mnemonic" yaml:"mnemonic"`
	HDPath   string `toml:"hd_path" yaml:"hd_path"`
}
