package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

// Server represents the server for the application.
type Server struct {
	handler    http.Handler
	port       string
	httpServer *http.Server
}

// NewServer creates a new server for the application.
func NewServer(handler http.Handler, port string) *Server {
	httpServer := &http.Server{
		Handler:           handler,
		Addr:              port,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &Server{
		handler:    handler,
		port:       port,
		httpServer: httpServer,
	}
}

// Start starts the server with graceful shutdown handling.
func (s *Server) Start(ctx context.Context) {
	serverErrs := make(chan error, 1)
	defer close(serverErrs)

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", s.port)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrs <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		err := s.httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("http server shutdown error: %v", err)
		} else {
			log.Printf("http server stopped gracefully")
		}

	case err := <-serverErrs:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("http server listen and serve finished with closed error: %v", err)

			shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			err = s.httpServer.Shutdown(shutdownCtx)
			if err != nil {
				log.Printf("http server shutdown error: %v", err)
			} else {
				log.Printf("http server stopped gracefully")
			}
		} else {
			log.Printf("http server listen and serve finished with closed error: %v", err)
		}
	}
}
