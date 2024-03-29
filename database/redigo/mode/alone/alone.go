package alone

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/team-dandelion/go-dandelion/database/redigo"
	"github.com/team-dandelion/go-dandelion/database/redigo/logger"
)

type aloneMode struct{ pool *redis.Pool }

func (am *aloneMode) GetConn() redis.Conn {
	return am.pool.Get()
}

func (am *aloneMode) NewConn() (redis.Conn, error) {
	return am.pool.Dial()
}

func (am *aloneMode) Close() error {
	return am.pool.Close()
}

func (am *aloneMode) String() string { return "alone" }

func New(optFuncs ...OptFunc) redigo.ModeInterface {
	opts := options{
		addr:     "127.0.0.1:6379",
		dialOpts: redigo.DefaultDialOpts(),
		poolOpts: redigo.DefaultPoolOpts(),
	}
	for _, optFunc := range optFuncs {
		optFunc(&opts)
	}
	pool := &redis.Pool{
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			conn, err := redis.Dial("tcp", opts.addr, opts.dialOpts...)
			if err != nil {
				return nil, err
			}
			return &logger.LoggingConn{Conn: conn}, nil
		},
	}
	for _, poolOptFunc := range opts.poolOpts {
		poolOptFunc(pool)
	}
	return &aloneMode{pool: pool}
}

func NewClient(optFuncs ...OptFunc) *redigo.Client {
	return redigo.New(New(optFuncs...))
}
