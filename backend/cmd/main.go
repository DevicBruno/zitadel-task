package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

// Configuration constants
const (
	domain     = "bruno-devic-interview-task-instance-ujvch7.us1.zitadel.cloud"
	keyPath    = "key.json"
	serverPort = ":8080"
)

// Response represents the API response structure
type Response struct {
	Message string `json:"message"`
}

func jsonResponse(w http.ResponseWriter, resp any, status int) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func main() {
	ctx := context.Background()

	// Initialize ZITADEL authorization
	authZ, err := authorization.New(ctx, zitadel.New(domain), oauth.DefaultAuthorization(keyPath))
	if err != nil {
		log.Fatalf("Failed to initialize ZITADEL authorization: %v", err)
	}

	// Initialize the HTTP middleware
	mw := middleware.New(authZ)

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          true,
	})

	// Create a new mux router
	mux := http.NewServeMux()

	// Public endpoint
	mux.Handle("/api/public", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := jsonResponse(w, Response{Message: "This is a public endpoint - Hello!"}, http.StatusOK)
			if err != nil {
				log.Printf("error writing response: %v", err)
			}
		}))

	// Protected endpoint - requires authentication
	mux.Handle("/api/private", mw.RequireAuthorization()(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authCtx := mw.Context(r.Context())
			log.Printf("user accessed private endpoint - user with ID: %s", authCtx.UserID())

			err := jsonResponse(w, Response{
				Message: fmt.Sprintf("This is a protected endpoint - You are authenticated! Welcome user with ID: %s", authCtx.UserID()),
			}, http.StatusOK)
			if err != nil {
				log.Printf("error writing response: %v", err)
			}
		})))

	// Wrap the mux with CORS middleware
	handler := c.Handler(mux)

	// Start the server
	log.Printf("Server starting on port %s", serverPort)
	if err := http.ListenAndServe(serverPort, handler); err != nil {
		log.Fatal(err)
	}
}
