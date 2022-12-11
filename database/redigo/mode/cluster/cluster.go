package cluster

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"

	"github.com/gly-hub/go-dandelion/database/redigo"
)

type clusterMode struct {
	rc *redisc.Cluster
}

func (cm *clusterMode) GetConn() redis.Conn {
	return cm.rc.Get()
}

func (cm *clusterMode) NewConn() (redis.Conn, error) {
	return cm.rc.Dial()
}

func (cm *clusterMode) Close() error {
	return cm.rc.Close()
}

func (cm *clusterMode) String() string { return "cluster" }

func New(optFuncs ...OptFunc) redigo.ModeInterface {
	opts := options{
		nodes: []string{
			"127.0.0.1:30001", "127.0.0.1:30002", "127.0.0.1:30003",
			"127.0.0.1:30004", "127.0.0.1:30005", "127.0.0.1:30006",
		},
		dialOpts: redigo.DefaultDialOpts(),
		poolOpts: redigo.DefaultPoolOpts(),
	}
	for _, optFunc := range optFuncs {
		optFunc(&opts)
	}
	rc := &redisc.Cluster{
		StartupNodes: opts.nodes,
		DialOptions:  opts.dialOpts,
		PoolWaitTime: opts.waitTime,
		CreatePool: func(address string, options ...redis.DialOption) (*redis.Pool, error) {
			pool := &redis.Pool{
				Dial: func() (redis.Conn, error) {
					return redis.Dial("tcp", address, options...)
				},
			}
			for _, poolOptFunc := range opts.poolOpts {
				poolOptFunc(pool)
			}
			return pool, nil
		},
	}
	return &clusterMode{rc: rc}
}

func NewClient(optFuncs ...OptFunc) *redigo.Client {
	return redigo.New(New(optFuncs...))
}
