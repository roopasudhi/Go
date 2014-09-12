package exercise

import (
	"net"
	"net/http"
	"strings"
)
const (
	GET    = "GET"
	PUT    = "PUT"
	POST   = "POST"
	DELETE = "DELETE"
)
type HttpServer struct {
	port     string
	listener net.Listener
	server   *http.Server
	mux      *http.ServeMux
	exit     bool
}

func NewHttpServer(port string) *HttpServer {
	return &HttpServer{port: port,
		exit: true,
		server: &http.Server{
			Handler: nil,
		},
		mux: http.NewServeMux(),
	}
}

func (s *HttpServer) ListenAndServe() {
	if s.port == "" {
		return
	}
	s.exit = false
	var err error
	s.listener, err = net.Listen("tcp", s.port)
	if err != nil {
		panic(err)
	}
	err = s.server.Serve(s.listener)
	if !strings.Contains(err.Error(), "closed") {
		panic(err)
	}
}

func (s *HttpServer) Stop() {
	s.exit = true
	s.server = nil
	s.listener.Close()
}

func (s *HttpServer) AddHandler(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
	s.server.Handler = s.mux
}
