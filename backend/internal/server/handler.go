package server

import (
	"net/http"

	"github.com/rs/cors"
)

// NewHandler creates a new handler for the application.
func NewHandler(frontendURL string, mux http.Handler) http.Handler {
	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{frontendURL},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          true,
	})

	return c.Handler(mux)
}
