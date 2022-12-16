package config

import (
	xtime "github.com/gly-hub/go-dandelion/tools/time"
)

type WDB struct {
	Type     string `json:"type" yaml:"type"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Name     string `json:"name" yaml:"name"`
}

type RDB struct {
	Type     string `json:"type" yaml:"type"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Name     string `json:"name" yaml:"name"`
}

type Redis struct {
	RedisType    string         `json:"redis_type" yaml:"redisType"` //cluster,alone,sentinel
	Network      string         `json:"network" yaml:"network"`
	StartAddr    []string       `json:"start_addr" yaml:"startAddr"` // Startup nodes
	Active       int            `json:"active" yaml:"active"`
	Idle         int            `json:"idle" yaml:"idle"`
	Auth         string         `json:"auth" yaml:"auth"`
	ConnTimeout  xtime.Duration `json:"conn_timeout" yaml:"connTimeout"`   // Connection timeout
	ReadTimeout  xtime.Duration `json:"read_timeout" yaml:"readTimeout"`   // Read timeout
	WriteTimeout xtime.Duration `json:"write_timeout" yaml:"writeTimeout"` // Write timeout
	IdleTimeout  xtime.Duration `json:"idle_timeout" yaml:"idleTimeout"`
}
