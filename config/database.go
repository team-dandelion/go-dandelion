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

// Mongo MongoDB配置
type Mongo struct {
	MaxPoolSize            int      `json:"max_pool_size" yaml:"maxPoolSize"`
	MinPoolSize            int      `json:"min_pool_size" yaml:"minPoolSize"`
	MaxConnIdleTime        int      `json:"max_conn_idle_time" yaml:"maxConnIdleTime"`
	ConnectTimeout         string   `json:"connect_timeout" yaml:"connectTimeout"`
	SocketTimeout          string   `json:"socket_timeout" yaml:"socketTimeout"`
	ServerSelectionTimeout string   `json:"server_selection_timeout" yaml:"serverSelectionTimeout"`
	LogLevel               int      `json:"log_level" yaml:"logLevel"`
	SlowThreshold          string   `json:"slow_threshold" yaml:"slowThreshold"`
	Hosts                  []string `json:"hosts" yaml:"hosts"`                    // MongoDB副本集主机列表
	Database               string   `json:"database" yaml:"database"`              // 数据库名
	Username               string   `json:"username" yaml:"username"`              // 用户名
	Password               string   `json:"password" yaml:"password"`              // 密码
	AuthDB                 string   `json:"auth_db" yaml:"authDB"`                 // 认证数据库
	ReplicaSet             string   `json:"replica_set" yaml:"replicaSet"`         // 副本集名称（可选）
	ReadPreference         string   `json:"read_preference" yaml:"readPreference"` // 读偏好：primary, primaryPreferred, secondary, secondaryPreferred, nearest
}
