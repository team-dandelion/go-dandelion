package tools

import (
	"github.com/gly-hub/go-dandelion/database/redigo"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gomodule/redigo/redis"
	"sync"
	"unsafe"
)

type EventType string

const (
	ExpireEvent EventType = "__keyevent@0__:expired"
)

type PSubscribeCallback func(pattern, channel, device string)

type ExpireKeyCallBack func(key string)

var (
	cbMap       sync.Map
	expireFuncs sync.Map
)

func PSubscribeListen(conn *redigo.Client) (err error) {
	err = conn.Subscribe(func(c redis.PubSubConn) (err error) {
		// 订阅过期事件
		err = c.PSubscribe(ExpireEvent)
		if err != nil {
			logger.Error("redis Subscribe error.")
			return err
		}
		cbMap.Store(string(ExpireEvent), timeoutEventCallBack)
		go func() {
			for {
				switch res := c.Receive().(type) {
				case redis.Message:
					pattern := &res.Pattern
					channel := &res.Channel
					message := (*string)(unsafe.Pointer(&res.Data))

					if cb, ok := cbMap.Load(*channel); ok {
						cb.(PSubscribeCallback)(*pattern, *channel, *message)
					}
				case redis.Subscription:
					logger.Info("%s: %s %d", res.Channel, res.Kind, res.Count)
				case error:
					logger.Error("error handle...")
					continue
				}
			}
		}()
		return
	})
	return err
}

func timeoutEventCallBack(pattern, channel, device string) {
	// 获取注册
	expireFuncs.Range(func(key, function interface{}) bool {
		if len(device) > len(key.(string)) && key == device[:len(key.(string))] {
			function.(ExpireKeyCallBack)(device)
		}
		return true
	})
}

func RegisterExpireFunc(keyPrefix string, function ExpireKeyCallBack) {
	expireFuncs.Store(keyPrefix, function)
}

func RegisterFunc(channel string, cb PSubscribeCallback) {
	cbMap.Store(channel, cb)
}
