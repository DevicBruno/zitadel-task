package main

import (
	"context"

	"github.com/DevicBruno/zitadel-task/backend/cmd/bootstrap"
	"github.com/DevicBruno/zitadel-task/backend/cmd/config"
	"github.com/DevicBruno/zitadel-task/backend/internal/server"
)

func main() {
	bootstrap.NewConfig()

	ctx := context.Background()

	// Initialize ZITADEL authorization and middleware
	mw := bootstrap.InitializeAuth(ctx)

	// Create a new mux router
	mux := server.NewMux(mw)

	// Wrap the mux with CORS middleware
	handler := server.NewHandler(config.Config.FrontendURL, mux)

	// Create a new server and start it
	srv := server.NewServer(handler, config.Config.ServerPort)
	srv.Start(ctx)
}
