package main

import (
	"context"
	"log"

	"github.com/DevicBruno/zitadel-task/backend/cmd/bootstrap"
	"github.com/DevicBruno/zitadel-task/backend/cmd/config"
	"github.com/DevicBruno/zitadel-task/backend/internal/server"
)

func main() {
	bootstrap.NewConfig()

	ctx := context.Background()

	// Initialize ZITADEL authorization
	authZ := bootstrap.NewAuthZ(ctx)

	// Initialize the HTTP middleware
	mw := bootstrap.NewMiddleware(authZ)

	// Create a new mux router
	mux := server.NewMux(mw)

	// Wrap the mux with CORS middleware
	handler := server.NewHandler(config.Config.FrontendURL, mux)

	// Create a new server and start it
	srv := server.NewServer(handler, config.Config.ServerPort)

	log.Printf("Server starting on port %s", config.Config.ServerPort)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
