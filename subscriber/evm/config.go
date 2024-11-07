package evm

type Config struct {
	Chain     Chain      `yaml:"chain"`
	Contracts []Contract `yaml:"contracts"`
}

type Chain struct {
	RPC           string `yaml:"rpc"`
	ChainID       int    `yaml:"chain_id"`
	ScanInterval  int    `yaml:"scan_interval"`
	Confirmations int    `yaml:"confirmations"`
}

type Contract struct {
	Name      string `yaml:"name"`
	Address   string `yaml:"address"`
	ABI       string `yaml:"abi"`
	FromBlock int64  `yaml:"from_block"`
	ToBlock   int64  `yaml:"to_block"`
}
