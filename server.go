package gopitman

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	hhtpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.hhtpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.hhtpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.hhtpServer.Shutdown(ctx)
}
