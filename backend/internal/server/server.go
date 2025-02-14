package server

import (
	"net/http"
)

// Server represents the server for the application.
type Server struct {
	handler http.Handler
	port    string
}

// NewServer creates a new server for the application.
func NewServer(handler http.Handler, port string) *Server {
	return &Server{handler: handler, port: port}
}

// Start starts the server.
func (s *Server) Start() error {
	return http.ListenAndServe(s.port, s.handler)
}
