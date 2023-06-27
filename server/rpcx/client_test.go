package rpcx

import (
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCustomOptions(t *testing.T) {
	CustomOptions(func() client.Option {
		return client.Option{
			Retries:            5,                // 重试次数
			TimeToDisallow:     time.Minute,      // 30秒内不会对失败的服务器进行重试
			ConnectTimeout:     3 * time.Second,  // 连接超时
			IdleTimeout:        10 * time.Second, // 最大空闲时间
			BackupLatency:      10 * time.Millisecond,
			SerializeType:      protocol.MsgPack,
			CompressType:       protocol.None,
			TCPKeepAlivePeriod: time.Minute,
		}
	})
	opt := client.Option{
		Retries:            5,                // 重试次数
		TimeToDisallow:     time.Minute,      // 30秒内不会对失败的服务器进行重试
		ConnectTimeout:     3 * time.Second,  // 连接超时
		IdleTimeout:        10 * time.Second, // 最大空闲时间
		BackupLatency:      10 * time.Millisecond,
		SerializeType:      protocol.MsgPack,
		CompressType:       protocol.None,
		TCPKeepAlivePeriod: time.Minute,
	}
	got := option()
	assert.Equal(t, opt, got)
}

func TestNewRPCClient(t *testing.T) {
	gotC, err := NewRPCClient(ClientConfig{
		ClientName:      "testServer",
		BasePath:        "test",
		RegisterPlugin:  "etcc",
		RegisterServers: []string{},
		FailRetryModel:  2,
		BalanceModel:    2,
		PoolSize:        1,
	})
	assert.NotEmpty(t, err)
	assert.Nil(t, gotC)
}

func Test_balance(t *testing.T) {
	model := balance(RoundRobin)
	assert.Equal(t, client.RoundRobin, model)
}

func Test_failRetry(t *testing.T) {
	assert.Equal(t, client.Failtry, failRetry(FailTry))
}
