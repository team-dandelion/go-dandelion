package alone

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

var client = NewClient()
var message = "hello world"
var channel = "test-channel"

func BenchmarkAloneMode_Exec(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res, err := client.String(func(c redis.Conn) (res interface{}, err error) {
				return c.Do("ECHO", message)
			})
			assert.Nil(b, err)
			assert.Equal(b, message, res)
		}
	})
}

func BenchmarkAloneMode_Sub(b *testing.B) {
	var counter int32
	var notifyChan = make(chan struct{})
	go client.Subscribe(func(c redis.PubSubConn) (err error) {
		if err = c.Subscribe(channel); err != nil {
			b.Error(err)
		}
		for {
			switch msg := c.ReceiveWithTimeout(0).(type) {
			case redis.Subscription:
				notifyChan <- struct{}{}
			case redis.Message:
				atomic.AddInt32(&counter, -1)
				assert.EqualValues(b, message, msg.Data)
			case error:
				b.Errorf("receive failed, err = %s", msg)
			}
		}
	})
	<-notifyChan
	go func() {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, err := client.Execute(func(c redis.Conn) (res interface{}, err error) {
					return c.Do("PUBLISH", channel, message)
				})
				assert.Nil(b, err)
				atomic.AddInt32(&counter, 1)
			}
		})
		notifyChan <- struct{}{}
	}()
	<-notifyChan
	b.StopTimer()
	time.Sleep(time.Second)
	assert.Equal(b, int32(0), counter)
}
