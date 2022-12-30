package rpcx

import (
	"context"
	"fmt"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"hash/crc32"
	"time"
)

func NodeId(addr string) uint32 {
	return crc32.ChecksumIEEE([]byte(addr))
}
func makeLocalName(serviceType string, nodeId uint32) string {
	return fmt.Sprintf("%s#%d", serviceType, nodeId)
}

type AuthCall func(ctx context.Context, token string) error
type BroadCall func(ctx context.Context, service string, content []byte)
type RPCxServer struct {
	s           *server.Server
	nodeId      uint32
	addr        string
	serviceType string
	authCall    AuthCall
	BRCall      BroadCall
}

/**
 * @Description:
 * @param serviceType   服务类型。如：recharge服务 或 account服务
 * @param addr 服务地址
 * @param rcvHandle  处理类
 * @return *RPCxServer
 */
func NewRpcxServer(serviceType string, addr string, rcvHandle interface{}) *RPCxServer {
	s := server.NewServer()
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		EtcdServers:    rpcxEtcdAddr,
		BasePath:       rpcxBasePath,
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		return nil
	}
	s.Plugins.Add(r)
	//服务名
	s.RegisterName(serviceType, rcvHandle, "")
	nodeId := NodeId(addr)
	s.RegisterName(makeLocalName(serviceType, nodeId), rcvHandle, "")
	s.Plugins.Add(&ServerLoggerPlugin{})

	rpcS := &RPCxServer{s: s, addr: addr, serviceType: serviceType, nodeId: nodeId}
	s.RegisterFunctionName(fmt.Sprintf("SER-%s", serviceType), "broadcast", rpcS.BroadMsg, "")
	s.AuthFunc = rpcS.auth

	return rpcS
}

func (r *RPCxServer) SetAuthCall(authCall AuthCall) {
	if authCall != nil {
		r.authCall = authCall

	}
}
func (r *RPCxServer) SetBroadCall(brCall BroadCall) {
	if brCall != nil {
		r.BRCall = brCall
	}
}
func (r *RPCxServer) BroadMsg(ctx context.Context, arg *BRMsgReq, reply *BRMsgRsp) error {
	fmt.Println("content====", string(arg.Content))
	if r.BRCall != nil {
		r.BRCall(ctx, arg.Service, arg.Content)
	}

	return nil
}
func (r *RPCxServer) GetNodeId() uint32 {
	return r.nodeId
}
func (r *RPCxServer) GetServiceType() string {
	return r.serviceType
}
func (r *RPCxServer) RegisterName(serviceName string, rcvr interface{}, metadata string) error {
	return r.s.RegisterName(serviceName, rcvr, metadata)
}
func (r *RPCxServer) auth(ctx context.Context, req *protocol.Message, token string) error {
	if r.authCall != nil {
		return r.authCall(ctx, token)
	}

	return nil
}

func (r *RPCxServer) Start() {
	port := ParsePort(r.addr)
	go func() {
		err := r.s.Serve("tcp", fmt.Sprintf("%s", port))
		if err != nil {
			panic(err)
		}
	}()

}

func ParsePort(addr string) (port string) {
	for pos, c := range addr {
		switch c {
		case ':':
			port = addr[pos:]
		}
	}
	return
}
