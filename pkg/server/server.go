package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	inner http.Server
}

func NewServer(addr string, router *mux.Router) *Server {
	return &Server{
		inner: http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *Server) ListenAndServe() error {
	if err := s.inner.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.inner.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
