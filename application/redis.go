package application

import (
	"github.com/gly-hub/go-dandelion/config"
	"github.com/gly-hub/go-dandelion/database/redigo"
	"github.com/gly-hub/go-dandelion/database/redigo/smart_redis"
	timex "github.com/gly-hub/go-dandelion/tools/time"
)

var redis *smart_redis.SmartRedis

func initRedis() {
	if config.Conf.Redis == nil {
		return
	}
	var err error
	redis, err = smart_redis.NewSmartRedis(&smart_redis.Config{
		RedisType:    config.Conf.Redis.RedisType,
		Network:      config.Conf.Redis.Network,
		StartAddr:    config.Conf.Redis.StartAddr,
		Active:       config.Conf.Redis.Active,
		Idle:         config.Conf.Redis.Idle,
		Auth:         config.Conf.Redis.Auth,
		ConnTimeout:  timex.ParseDuration(config.Conf.Redis.ConnTimeout),
		ReadTimeout:  timex.ParseDuration(config.Conf.Redis.ReadTimeout),
		WriteTimeout: timex.ParseDuration(config.Conf.Redis.WriteTimeout),
		IdleTimeout:  timex.ParseDuration(config.Conf.Redis.IdleTimeout),
	})
	if err != nil {
		panic(err)
	}

	//err = tools.PSubscribeListen(redis.Client)
	if err != nil {
		panic(err)
	}
}

type Redis struct {
}

func (r *Redis) GetRedis() *redigo.Client {
	return redis.Client
}
