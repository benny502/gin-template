package pkg

import (
	"bookmark/internal/pkg/log"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(log.NewLogger, log.NewWriter)
