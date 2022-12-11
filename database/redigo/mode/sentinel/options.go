package sentinel

import (
	"github.com/gomodule/redigo/redis"

	"github.com/gly-hub/go-dandelion/database/redigo"
)

type options struct {
	addrs            []string
	masterName       string
	poolOpts         []redigo.PoolOption
	dialOpts         []redis.DialOption
	sentinelDialOpts []redis.DialOption
}

type OptFunc func(opts *options)

func Addrs(value []string) OptFunc {
	return func(opts *options) {
		opts.addrs = value
	}
}

func MasterName(value string) OptFunc {
	return func(opts *options) {
		opts.masterName = value
	}
}

func PoolOpts(value ...redigo.PoolOption) OptFunc {
	return func(opts *options) {
		for _, poolOpt := range value {
			opts.poolOpts = append(opts.poolOpts, poolOpt)
		}
	}
}

func DialOpts(value ...redis.DialOption) OptFunc {
	return func(opts *options) {
		for _, dialOpt := range value {
			opts.dialOpts = append(opts.dialOpts, dialOpt)
		}
	}
}

func DialSentinelOpts(value ...redis.DialOption) OptFunc {
	return func(opts *options) {
		for _, dialOpt := range value {
			opts.sentinelDialOpts = append(opts.sentinelDialOpts, dialOpt)
		}
	}
}
