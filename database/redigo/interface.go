package redigo

import (
	"fmt"
	"io"
	"time"

	"github.com/gomodule/redigo/redis"
)

type ModeInterface interface {
	io.Closer
	fmt.Stringer
	GetConn() redis.Conn
	NewConn() (redis.Conn, error)
}

// DefaultDialOpts 默认连接配置
func DefaultDialOpts() []redis.DialOption {
	return []redis.DialOption{
		// 使用Go默认心跳间隔
		redis.DialKeepAlive(time.Second * 15),
		redis.DialConnectTimeout(time.Second),
		redis.DialReadTimeout(time.Second * 3),
		redis.DialWriteTimeout(time.Second * 3),
	}
}

// DefaultPoolOpts 默认连接池配置
func DefaultPoolOpts() []PoolOption {
	return []PoolOption{
		Wait(true), MaxIdle(5),
		TestOnBorrow(func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		}),
	}
}
