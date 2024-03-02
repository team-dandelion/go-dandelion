package smart_redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestNewSmartRedis(t *testing.T) {
	client, err := NewSmartRedis(&Config{
		RedisType: "cluster",
		StartAddr: []string{
			"127.0.0.1:6385", "127.0.0.1:6380", "127.0.0.1:6381", "127.0.0.1:6382", "127.0.0.1:6383", "127.0.0.1:6384",
		},
		Active:       10,
		Idle:         10,
		Auth:         "gaoliangyong",
		ConnTimeout:  100 * time.Millisecond,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 600 * time.Millisecond,
		IdleTimeout:  600 * time.Millisecond,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Client.Close()
	_, err = client.Client.Execute(func(c redis.Conn) (res interface{}, err error) {

		return c.Do("set", "fff0", "1234")
	})
	fmt.Println(err)
}
