package cluster

import (
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/gly-hub/go-dandelion/database/redigo"
)

type options struct {
	nodes    []string
	waitTime time.Duration
	poolOpts []redigo.PoolOption
	dialOpts []redis.DialOption
}

type OptFunc func(opts *options)

func Nodes(value []string) OptFunc {
	return func(opts *options) {
		opts.nodes = value
	}
}

func WaitTime(value time.Duration) OptFunc {
	return func(opts *options) {
		opts.waitTime = value
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
