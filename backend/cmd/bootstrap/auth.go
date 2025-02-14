package bootstrap

import (
	"context"
	"log"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"

	"github.com/DevicBruno/zitadel-task/backend/cmd/config"
)

// InitializeAuth initializes both ZITADEL authorization and middleware.
func InitializeAuth(ctx context.Context) *middleware.Interceptor[*oauth.IntrospectionContext] {
	authZ, err := authorization.New(ctx, zitadel.New(config.Config.Domain), oauth.DefaultAuthorization(config.Config.KeyPath))
	if err != nil {
		log.Fatalf("Failed to initialize ZITADEL authorization: %v", err)
	}

	return middleware.New(authZ)
}
