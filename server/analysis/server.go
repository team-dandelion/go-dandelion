package analysis

import (
	"github.com/spf13/cast"
	"net"
	"net/http"
	_ "net/http/pprof"
)

type Server struct {
	Port       int32
	Prometheus *Prometheus
}

type Config struct {
	Port        int32
	Prometheus  bool
	ServiceName string
}

func New(config Config) *Server {
	if config.Port == 0 {
		return nil
	}
	prom := NewPrometheus(config.ServiceName)
	return &Server{
		Port:       config.Port,
		Prometheus: prom,
	}
}

func (s *Server) RunServer() {
	http.Handle("/metrics", s.Prometheus.HandlerFunc())
	listener, _ := net.Listen("tcp", net.JoinHostPort("", cast.ToString(s.Port)))
	_ = http.Serve(listener, nil)
}
