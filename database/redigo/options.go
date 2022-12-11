package redigo

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// PoolOption 连接池配置函数
type PoolOption func(pool *redis.Pool)

// Wait 连接池枯竭是是否阻塞等待
func Wait(value bool) PoolOption {
	return func(pool *redis.Pool) {
		pool.Wait = value
	}
}

// MaxIdle 最多允许空闲的连接数
func MaxIdle(value int) PoolOption {
	return func(pool *redis.Pool) {
		pool.MaxIdle = value
	}
}

// MaxActive 最多允许的连接数, 0 则无限制
func MaxActive(value int) PoolOption {
	return func(pool *redis.Pool) {
		pool.MaxActive = value
	}
}

// IdleTimeout 最大空闲时长, 超过后关闭连接
func IdleTimeout(value time.Duration) PoolOption {
	return func(pool *redis.Pool) {
		pool.IdleTimeout = value
	}
}

// MaxConnLifetime 连接生命周期, 超过后关闭连接
func MaxConnLifetime(value time.Duration) PoolOption {
	return func(pool *redis.Pool) {
		pool.MaxConnLifetime = value
	}
}

// TestOnBorrow 健康检查, 检测连接是否可用
func TestOnBorrow(value func(c redis.Conn, t time.Time) (err error)) PoolOption {
	return func(pool *redis.Pool) {
		pool.TestOnBorrow = value
	}
}
