package config

type HttpServer struct {
	Port int32 `json:"port" yaml:"port"`
}

type RpcServer struct {
	ServerName string   `json:"server_name" yaml:"serverName"`
	BasePath   string   `json:"base_path" yaml:"basePath"`
	Etcd       []string `json:"etcd" yaml:"etcd"`
	Addr       string   `json:"addr" yaml:"addr"`
	Port       int      `json:"port" yaml:"port"`
	Model      int      `json:"model" yaml:"model"`
}
