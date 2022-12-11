package smart_redis

import (
	"errors"
	"github.com/gly-hub/go-dandelion/database/redigo"
	"github.com/gly-hub/go-dandelion/database/redigo/mode/alone"
	"github.com/gly-hub/go-dandelion/database/redigo/mode/cluster"
	"github.com/gly-hub/go-dandelion/database/redigo/mode/sentinel"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Config struct {
	RedisType    string //cluster,alone,sentinel
	Network      string
	StartAddr    []string // Startup nodes
	Active       int
	Idle         int
	Auth         string
	ConnTimeout  time.Duration // Connection timeout
	ReadTimeout  time.Duration // Read timeout
	WriteTimeout time.Duration // Write timeout
	IdleTimeout  time.Duration
}
type SmartRedis struct {
	RedisType string
	Client    *redigo.Client
}

func NewSmartRedis(conf *Config) (*SmartRedis, error) {
	if conf.RedisType == "cluster" {
		client := cluster.NewClient(cluster.Nodes(conf.StartAddr),
			cluster.PoolOpts(redigo.Wait(true),
				redigo.MaxActive(conf.Active),
				redigo.MaxIdle(conf.Idle),
				redigo.IdleTimeout(conf.IdleTimeout)),
			cluster.DialOpts(redis.DialConnectTimeout(conf.ConnTimeout),
				redis.DialPassword(conf.Auth),
				redis.DialReadTimeout(conf.ReadTimeout),
				redis.DialWriteTimeout(conf.WriteTimeout)))
		return &SmartRedis{
			RedisType: conf.RedisType,
			Client:    client,
		}, nil
	} else if conf.RedisType == "alone" {
		client := alone.NewClient(alone.Addr(conf.StartAddr[0]),
			alone.PoolOpts(redigo.Wait(true),
				redigo.MaxActive(conf.Active),
				redigo.MaxIdle(conf.Idle),
				redigo.IdleTimeout(conf.IdleTimeout)),
			alone.DialOpts(redis.DialConnectTimeout(conf.ConnTimeout),
				redis.DialPassword(conf.Auth),
				redis.DialReadTimeout(conf.ReadTimeout),
				redis.DialWriteTimeout(conf.WriteTimeout)))
		return &SmartRedis{
			RedisType: conf.RedisType,
			Client:    client,
		}, nil
	} else if conf.RedisType == "sentinel" {
		client := sentinel.NewClient(sentinel.Addrs(conf.StartAddr),
			sentinel.PoolOpts(redigo.Wait(true),
				redigo.MaxActive(conf.Active),
				redigo.MaxIdle(conf.Idle),
				redigo.IdleTimeout(conf.IdleTimeout)),
			sentinel.DialOpts(redis.DialConnectTimeout(conf.ConnTimeout),
				redis.DialPassword(conf.Auth),
				redis.DialReadTimeout(conf.ReadTimeout),
				redis.DialWriteTimeout(conf.WriteTimeout)))
		return &SmartRedis{
			RedisType: conf.RedisType,
			Client:    client,
		}, nil
	}
	return nil, errors.New("redis no this mode")
}
