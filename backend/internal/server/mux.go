package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
)

// NewMux creates a new mux for the application.
func NewMux[T authorization.Ctx](mw *middleware.Interceptor[T]) *http.ServeMux {
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

	return mux
}
