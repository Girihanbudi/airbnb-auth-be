package server

import (
	"airbnb-auth-be/internal/pkg/log"
	"context"
	"fmt"
	"net/http"
)

func (s *Server) Start() error {
	log.Event(Instance, "starting server...")
	s.address = fmt.Sprintf("%s:%s", s.Host, s.Port)

	s.server = &http.Server{
		Addr:    s.address,
		Handler: s.Options.Router,
	}

	log.Event(Instance, fmt.Sprintf("listening on %s", s.address))
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	log.Event(Instance, "shutting down server...")
	if err := s.server.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
