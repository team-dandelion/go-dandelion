package gorm

import (
	"time"
)

const (
	GoDandelionConnChange = "go_dandelion_conn_change"
)

type ChangeType int

const (
	Created ChangeType = iota + 1
	Updated
	Deleted
)

type AppConfig struct {
	AppKey   string
	DBMaster *Master
	DBSlave  []*Slave
}

type Config struct {
	DBType        string        `json:"db_type"`
	MaxOpenConn   int           `json:"max_open_conn"`
	MaxIdleConn   int           `json:"max_idle_conn"`
	MaxLifeTime   int           `json:"max_life_time"`
	MaxIdleTime   int           `json:"max_idle_time"`
	Level         int           `json:"level"`
	SlowThreshold time.Duration `json:"slow_threshold"`
	Master        *Master       `json:"master"`
	Slaves        []*Slave      `json:"slaves"`
}

type Master struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type Slave struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type AppDBMessage struct {
	AppKey     string
	ChangeType ChangeType
}
