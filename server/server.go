package server

import (
	"context"
	"github.com/fanfaronDo/referral_system_api/config"
	"net/http"
)

type Server struct {
	serverHTTP http.Server
}

func NewServer(host string, cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		serverHTTP: http.Server{
			Addr:           host,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    cfg.Timeout,
			WriteTimeout:   cfg.Timeout,
		},
	}
}

func (s *Server) Start() error {
	return s.serverHTTP.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.serverHTTP.Shutdown(ctx)
}
