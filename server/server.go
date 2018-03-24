package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// Server ...
type Server struct {
	instance *http.Server
	mux      *http.ServeMux
}

// New creates a new http server
func New(host string, port string) *Server {
	return &Server{
		instance: &http.Server{
			Addr: fmt.Sprintf("%s:%s", host, port),
		},
		mux: http.NewServeMux(),
	}
}

// RegisterHandler registers new handlers
func (s *Server) RegisterHandler(path string, handler http.HandlerFunc) {
	s.mux.Handle(path, handler)
}

// StartServer ...
func (s *Server) StartServer() {
	log.Printf("Starting server...\n")
	http.Handle("/", s.mux)
	go func() {
		// Graceful shutdown
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		sig := <-sigquit
		log.Printf("caught sig: %+v", sig)
		log.Printf("Gracefully shutting down server...")

		if err := s.instance.Shutdown(context.Background()); err != nil {
			log.Printf("Unable to shut down server: %v", err)
		} else {
			log.Println("Server stopped")
		}
	}()

	if err := s.instance.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
