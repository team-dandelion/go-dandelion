package cluster

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
	"github.com/team-dandelion/go-dandelion/database/redigo"
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
			"127.0.0.1:6380", "127.0.0.1:6381", "127.0.0.1:6382",
			"127.0.0.1:6383", "127.0.0.1:6384", "127.0.0.1:6385",
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

	if terr := rc.Refresh(); terr != nil {
		fmt.Println(terr)
	}
	return &clusterMode{rc: rc}
}

func NewClient(optFuncs ...OptFunc) *redigo.Client {
	return redigo.New(New(optFuncs...))
}
