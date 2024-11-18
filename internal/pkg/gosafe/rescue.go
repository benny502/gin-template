package gosafe

import (
	"bookmark/internal/middleware/log"
	"runtime"

	"github.com/gin-gonic/gin"
)

func Recovery(ctx *gin.Context, cleanups ...func()) {
	logger := log.FromContext(ctx)

	for _, cleanup := range cleanups {
		cleanup()
	}

	if logger != nil {
		if p := recover(); p != nil {
			buf := make([]byte, 64<<10) //nolint:gomnd
			n := runtime.Stack(buf, false)
			buf = buf[:n]
			logger.Errorf("%v:\n%s\n", p, buf)
		}
	}

}
