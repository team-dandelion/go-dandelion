package mongo

import (
	"time"
)

const (
	GoDandelionMongoConnChange = "go_dandelion_mongo_conn_change"
	MongoDB                    = "mongodb"
)

type ChangeType int

const (
	Created ChangeType = iota + 1
	Updated
	Deleted
)

type AppConfig struct {
	AppKey     string
	MongoHosts []string // MongoDB副本集主机列表
	Database   string   // 数据库名
	Username   string   // 用户名
	Password   string   // 密码
	AuthDB     string   // 认证数据库
	ReplicaSet string   // 副本集名称（可选）
}

type Config struct {
	MaxPoolSize            uint64        `json:"max_pool_size"`
	MinPoolSize            uint64        `json:"min_pool_size"`
	MaxConnIdleTime        time.Duration `json:"max_conn_idle_time"`
	ConnectTimeout         time.Duration `json:"connect_timeout"`
	SocketTimeout          time.Duration `json:"socket_timeout"`
	ServerSelectionTimeout time.Duration `json:"server_selection_timeout"`
	Hosts                  []string      `json:"hosts"`           // MongoDB副本集主机列表
	Database               string        `json:"database"`        // 数据库名
	Username               string        `json:"username"`        // 用户名
	Password               string        `json:"password"`        // 密码
	AuthDB                 string        `json:"auth_db"`         // 认证数据库
	ReplicaSet             string        `json:"replica_set"`     // 副本集名称（可选）
	ReadPreference         string        `json:"read_preference"` // 读偏好
	LogLevel               int           `json:"log_level"`
	SlowThreshold          time.Duration `json:"slow_threshold"`
}

type AppMongoMessage struct {
	AppKey     string
	ChangeType ChangeType
}
