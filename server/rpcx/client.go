package rpcx

import (
	"fmt"
	conClient "github.com/rpcxio/rpcx-consul/client"
	etcdClient "github.com/rpcxio/rpcx-etcd/client"
	zkClient "github.com/rpcxio/rpcx-zookeeper/client"
	"github.com/smallnest/rpcx/client"
)

const (
	// FailFast 如果调用失败，立即返回错误
	FailFast FailRetryModel = iota + 1
	// FailOver 如果调用失败，重试其他服务器
	FailOver
	// FailTry 如果调用失败，重试当前服务器
	FailTry
)

const (
	// Random 随机
	Random BalanceModel = iota + 1
	// RoundRobin 轮询
	RoundRobin
	// ConsistentHash 一致性哈希
	ConsistentHash
	// NetworkQuality 网络质量
	NetworkQuality
)

type (
	// FailRetryModel 失败重试模式
	FailRetryModel int
	// BalanceModel 负载均衡模式
	BalanceModel int

	ClientConfig struct {
		ClientName      string
		BasePath        string
		RegisterPlugin  RegisterPluginType
		RegisterServers []string
		FailRetryModel  FailRetryModel
		BalanceModel    BalanceModel
		PoolSize        int
	}
	ClientPool struct {
		clientPool *client.OneClientPool
	}
)

// 100毫秒后内失败次数达到5次，熔断器打开
// func() client.Breaker { return client.NewConsecCircuitBreaker(5, 100*time.Millisecond) }
var customOption func() client.Option

func NewRPCClient(conf ClientConfig) (c *ClientPool, err error) {
	var discovery client.ServiceDiscovery
	switch conf.RegisterPlugin {
	case Multiple:
		var servers []*client.KVPair
		for _, v := range conf.RegisterServers {
			servers = append(servers, &client.KVPair{
				Key: v,
			})
		}
		discovery, err = client.NewMultipleServersDiscovery(servers)
	case ETCD:
		discovery, err = etcdClient.NewEtcdV3Discovery(conf.BasePath, "",
			conf.RegisterServers, true, nil)
	case ZK:
		discovery, err = zkClient.NewZookeeperDiscovery(conf.BasePath, "",
			conf.RegisterServers, nil)
	case Con:
		discovery, err = conClient.NewConsulDiscovery(conf.BasePath, "",
			conf.RegisterServers, nil)
	default:
		err = fmt.Errorf("not support register plugin: %s", conf.RegisterPlugin)
	}
	if err != nil {
		return nil, err
	}

	clientPool := client.NewOneClientPool(
		conf.PoolSize,
		failRetry(conf.FailRetryModel),
		balance(conf.BalanceModel),
		discovery,
		option())
	return &ClientPool{
		clientPool: clientPool,
	}, nil

}

func (c *ClientPool) Client(auth ...string) *client.OneClient {
	oneClient := c.clientPool.Get()
	plugins := client.NewPluginContainer()
	plugins.Add(&ClientLoggerPlugin{})
	oneClient.SetPlugins(plugins)
	if len(auth) > 0 {
		oneClient.Auth(auth[0])
	}
	return oneClient
}

func CustomOptions(f func() client.Option) {
	customOption = f
}

func option() client.Option {
	if customOption != nil {
		return customOption()
	}

	//opt := client.Option{
	//	Retries:             10, // 重试次数
	//	RPCPath:             share.DefaultRPCPath,
	//	TimeToDisallow:      time.Minute, // 30秒内不会对失败的服务器进行重试
	//	ConnectTimeout:      time.Second, // 连接超时
	//	BackupLatency:       10 * time.Millisecond,
	//	SerializeType:       protocol.MsgPack,
	//	CompressType:        protocol.None,
	//	TCPKeepAlivePeriod:  time.Minute,
	//	MaxWaitForHeartbeat: 30 * time.Second,
	//	BidirectionalBlock:  false,
	//}

	opt := client.DefaultOption

	return opt
}

func failRetry(model FailRetryModel) client.FailMode {
	var m client.FailMode
	switch model {
	case FailFast:
		m = client.Failfast
	case FailOver:
		m = client.Failover
	case FailTry:
		m = client.Failtry
	default:
		m = client.Failover
	}
	return m
}

func balance(model BalanceModel) client.SelectMode {
	var m client.SelectMode
	switch model {
	case Random:
		m = client.RandomSelect
	case RoundRobin:
		m = client.RoundRobin
	case ConsistentHash:
		m = client.ConsistentHash
	case NetworkQuality:
		m = client.WeightedICMP
	default:
		m = client.RoundRobin
	}
	return m
}
