package alone

import (
	"github.com/gomodule/redigo/redis"

	"github.com/gly-hub/go-dandelion/database/redigo"
)

type aloneMode struct{ pool redis.Pool }

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
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", opts.addr, opts.dialOpts...)
		},
	}
	for _, poolOptFunc := range opts.poolOpts {
		poolOptFunc(&pool)
	}
	return &aloneMode{pool: pool}
}

func NewClient(optFuncs ...OptFunc) *redigo.Client {
	return redigo.New(New(optFuncs...))
}
