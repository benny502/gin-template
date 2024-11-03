package middleware

import (
	"bookmark/internal/middleware/auth"
	"bookmark/internal/middleware/cors"
	"bookmark/internal/middleware/log"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(auth.NewAuth, cors.NewCors, log.NewLogger)
