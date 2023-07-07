package analysis

import (
	"net"
	"net/http"
	_ "net/http/pprof"
)

type Server struct {
	Port string
}

func (s *Server) RunServer() {
	listener, _ := net.Listen("tcp", net.JoinHostPort("", s.Port))
	_ = http.Serve(listener, nil)
}
