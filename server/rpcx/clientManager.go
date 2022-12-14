package rpcx

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

var manager *CManager
var once sync.Once

type CManager struct {
	clients map[int]*RPCXClient
}

var rpcxBasePath string
var rpcxEtcdAddr []string

func SetBase(basePath string, etcdAddr []string) {
	rpcxBasePath = basePath
	rpcxEtcdAddr = etcdAddr
	once.Do(func() {
		manager = newManger()
	})
}
func newManger() *CManager {
	manager := &CManager{clients: make(map[int]*RPCXClient)}
	for i := RPC_One2One; i < RPC_MAX; i++ {
		manager.clients[i] = NewRPCXClient(i)
	}
	return manager
}
func SetAuth(token string, model int) {
	if model < RPC_One2One || model >= RPC_MAX {
		return
	}
	manager.clients[model].GetClient().Auth(token)
}

//点对点发送，ServiceType,nodeId,methodName
func One2OneCall(ctx context.Context, ServiceType string, nodeId uint32, method string, args interface{}, reply interface{}) error {
	svrName := makeLocalName(ServiceType, nodeId)
	cli := manager.clients[RPC_One2One].GetClient()
	return cli.Call(ctx, svrName, method, args, reply)
}

//点对随机
func One2RandCall(ctx context.Context, serviceName string, method string, args interface{}, reply interface{}) error {
	cli := manager.clients[RPC_One2Rand]
	return cli.GetClient().Call(ctx, serviceName, method, args, reply)
}

func BroadcastMsg(serviceName string, msgID string, msg interface{}) error {
	val, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	cli := manager.clients[RPC_One2All]
	req := &BRMsgReq{
		Service: serviceName,
		MsgId:   msgID,
		Content: val,
	}
	rsp := &BRMsgRsp{}
	return cli.GetClient().Broadcast(context.Background(), fmt.Sprintf("SER-%s", serviceName), "broadcast", req, rsp)
}
