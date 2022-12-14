package rpcx

import (
	etcd_client "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
)

type RPCXClient struct {
	rpcClient *client.OneClient
	etcdDis   client.ServiceDiscovery
}

func NewRPCXClient(Model int) *RPCXClient {
	d, err := etcd_client.NewEtcdV3Discovery(rpcxBasePath, "", rpcxEtcdAddr, true, nil)
	if err != nil {
		return nil
	}
	var clientRpc *client.OneClient
	switch Model {
	case RPC_One2One:
		clientRpc = client.NewOneClient(client.Failtry, client.RoundRobin, d, client.DefaultOption)
	case RPC_One2All:
		clientRpc = client.NewOneClient(client.Failtry, client.RoundRobin, d, client.DefaultOption)
	case RPC_One2Rand:
		clientRpc = client.NewOneClient(client.Failover, client.RoundRobin, d, client.DefaultOption)
	}

	return &RPCXClient{
		rpcClient: clientRpc,
		etcdDis:   d,
	}
}
func (r *RPCXClient) GetClient() *client.OneClient {
	return r.rpcClient
}
