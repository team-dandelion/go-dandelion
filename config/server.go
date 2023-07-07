package config

type HttpServer struct {
	Port int32 `json:"port" yaml:"port"`
}

type RpcServer struct {
	ServerName      string   `json:"server_name" yaml:"serverName"`
	BasePath        string   `json:"base_path" yaml:"basePath"`
	RegisterPlugin  string   `json:"register_plugin" yaml:"registerPlugin"`
	RegisterServers []string `json:"register_servers" yaml:"registerServers"`
	Addr            string   `json:"addr" yaml:"addr"`
	Port            int      `json:"port" yaml:"port"`
}

type RpcClient struct {
	ClientName      string   `json:"client_name" yaml:"clientName"`
	BasePath        string   `json:"base_path" yaml:"basePath"`
	RegisterPlugin  string   `json:"register_plugin" yaml:"registerPlugin"`
	RegisterServers []string `json:"register_servers" yaml:"registerServers"`
	FailRetryModel  int      `json:"fail_retry_model" yaml:"failRetryModel"`
	BalanceModel    int      `json:"balance_model" yaml:"balanceModel"`
	PoolSize        int      `json:"pool_size" yaml:"poolSize"`
}

type AnalysisServer struct {
	AnalysisName string `json:"analysis_name" yaml:"analysisName"`
	Pprof        int32  `json:"pprof" yaml:"pprof"`
	Prometheus   bool   `json:"prometheus" yaml:"prometheus"`
}
