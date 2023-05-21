package server

import (
	"airbnb-auth-be/internal/pkg/log"
	"context"
	"fmt"
	"net/http"
)

func (s *Server) Start() {
	log.Event(Instance, "starting http server...")

	s.server = &http.Server{
		Addr:    s.address,
		Handler: s.Options.Router,
	}

	if s.Creds.TlsCerts == nil {
		log.Event(Instance, fmt.Sprintf("listening on %s://%s", "http", s.address))
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(Instance, "failed to start http server", err)
		}
	} else {
		log.Event(Instance, fmt.Sprintf("listening on %s://%s", "https", s.address))
		if err := s.server.ListenAndServeTLS(s.Creds.PublicCert, s.Creds.PrivateKey); err != nil && err != http.ErrServerClosed {
			log.Fatal(Instance, "failed to start http server", err)
		}
	}
}

func (s *Server) Stop() {
	log.Event(Instance, "shutting down http server...")

	if err := s.server.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
		log.Fatal(Instance, "failed to shutting down http server", err)
	}

	log.Event(Instance, "http server has been shutted down")
}
