package config

type DB struct {
	DBType        string   `json:"db_type" yaml:"dbType"`
	MaxOpenConn   int      `json:"max_open_conn" yaml:"maxOpenConn"`
	MaxIdleConn   int      `json:"max_idle_conn" yaml:"maxIdleConn"`
	MaxLifeTime   int      `json:"max_life_time" yaml:"maxLifeTime"`
	MaxIdleTime   int      `json:"max_idle_time" yaml:"maxIdleTime"`
	Level         int      `json:"level" yaml:"level"`
	SlowThreshold string   `json:"slow_threshold" yaml:"slowThreshold"`
	Master        *Master  `json:"master" yaml:"master"`
	Slaves        []*Slave `json:"slaves" yaml:"slaves"`
}

type Master struct {
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Database string `json:"data_base" yaml:"database"`
}

type Slave struct {
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Database string `json:"data_base" yaml:"database"`
}

type Redis struct {
	RedisType    string   `json:"redis_type" yaml:"redisType"` //cluster,alone,sentinel
	StartAddr    []string `json:"start_addr" yaml:"startAddr"` // Startup nodes
	Active       int      `json:"active" yaml:"active"`
	Idle         int      `json:"idle" yaml:"idle"`
	Auth         string   `json:"auth" yaml:"auth"`
	ConnTimeout  string   `json:"conn_timeout" yaml:"connTimeout"`   // Connection timeout
	ReadTimeout  string   `json:"read_timeout" yaml:"readTimeout"`   // Read timeout
	WriteTimeout string   `json:"write_timeout" yaml:"writeTimeout"` // Write timeout
	IdleTimeout  string   `json:"idle_timeout" yaml:"idleTimeout"`
}
