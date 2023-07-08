package rpcx

import (
	"context"
	"fmt"
	"github.com/rcrowley/go-metrics"
	consulPlugin "github.com/rpcxio/rpcx-consul/serverplugin"
	etcdPlugin "github.com/rpcxio/rpcx-etcd/serverplugin"
	zookeeperPlugin "github.com/rpcxio/rpcx-zookeeper/serverplugin"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"hash/crc32"
	"time"
)

const (
	ETCD RegisterPluginType = "etcd"
	ZK   RegisterPluginType = "zookeeper"
	Con  RegisterPluginType = "consul"
)

type (
	RegisterPluginType string

	AuthFunc func(ctx context.Context, token string) error

	ServerConfig struct {
		ServerName      string
		Addr            string
		BasePath        string
		RegisterPlugin  RegisterPluginType
		RegisterServers []string
		Handle          interface{}
	}

	Server struct {
		s      *server.Server
		nodeId uint32
		addr   string
		Name   string
		// 认证回调
		authFunc AuthFunc
	}
)

func NewRPCServer(conf ServerConfig, plugins ...server.Plugin) (rpc *Server, err error) {
	// init server
	s := server.NewServer()
	// add plugin
	var plugin server.Plugin
	switch conf.RegisterPlugin {
	case ETCD:
		plugin, err = EtcdV3Register(conf)
	case ZK:
		plugin, err = ZooKeeperRegister(conf)
	case Con:
		plugin, err = ConsulRegister(conf)
	default:
		err = fmt.Errorf("not support register plugin: %s", conf.RegisterPlugin)
	}
	if err != nil {
		return nil, err
	}
	s.Plugins.Add(plugin)
	// register server name
	if err = s.RegisterName(conf.ServerName, conf.Handle, ""); err != nil {
		return nil, err
	}
	id := nodeId(conf.Addr)
	if err = s.RegisterName(makeLocalName(conf.ServerName, id), conf.Handle, ""); err != nil {
		return nil, err
	}

	// add middleware
	s.Plugins.Add(&ServerLoggerPlugin{})

	for _, p := range plugins {
		s.Plugins.Add(p)
	}

	// new rpc server
	rpc = &Server{
		s:      s,
		nodeId: id,
		addr:   conf.Addr,
		Name:   conf.ServerName,
	}
	// auth
	s.AuthFunc = rpc.auth

	return rpc, nil
}

// RegisterAuthFunc 注册认证回调
func (s *Server) RegisterAuthFunc(authFunc AuthFunc) {
	s.authFunc = authFunc
}

// Start 启动服务
func (s *Server) Start() {
	port := parsePort(s.addr)
	go func() {
		err := s.s.Serve("tcp", fmt.Sprintf("%s", port))
		if err != nil {
			panic(err)
		}
	}()
}

func (s *Server) auth(ctx context.Context, req *protocol.Message, token string) error {
	if s.authFunc != nil {
		return s.authFunc(ctx, token)
	}

	return nil
}

func EtcdV3Register(conf ServerConfig) (server.Plugin, error) {
	s := &etcdPlugin.EtcdV3RegisterPlugin{
		ServiceAddress: serverAddress(conf.Addr),
		EtcdServers:    conf.RegisterServers,
		BasePath:       conf.BasePath,
		UpdateInterval: time.Minute,
	}
	if err := s.Start(); err != nil {
		return nil, err
	}
	return s, nil
}

func ZooKeeperRegister(conf ServerConfig) (server.Plugin, error) {
	s := &zookeeperPlugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   serverAddress(conf.Addr),
		ZooKeeperServers: conf.RegisterServers,
		BasePath:         conf.BasePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	if err := s.Start(); err != nil {
		return nil, err
	}
	return s, nil
}

func ConsulRegister(conf ServerConfig) (server.Plugin, error) {
	s := &consulPlugin.ConsulRegisterPlugin{
		ServiceAddress: serverAddress(conf.Addr),
		ConsulServers:  conf.RegisterServers,
		BasePath:       conf.BasePath,
		UpdateInterval: time.Minute,
	}
	if err := s.Start(); err != nil {
		return nil, err
	}
	return s, nil
}

func serverAddress(addr string) string {
	return "tcp@" + addr
}

// nodeId 生成节点ID
func nodeId(addr string) uint32 {
	return crc32.ChecksumIEEE([]byte(addr))
}

// makeLocalName 生成本地服务名
func makeLocalName(serviceName string, nodeId uint32) string {
	return fmt.Sprintf("%s#%d", serviceName, nodeId)
}

func parsePort(addr string) (port string) {
	for pos, c := range addr {
		switch c {
		case ':':
			port = addr[pos:]
		}
	}
	return
}
