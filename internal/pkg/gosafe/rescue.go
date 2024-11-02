package gosafe

import (
	"bookmark/internal/pkg/log"
	"runtime"
)

func Recovery(logger log.Logger, cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		buf := make([]byte, 64<<10) //nolint:gomnd
		n := runtime.Stack(buf, false)
		buf = buf[:n]
		logger.Errorf("%v:\n%s\n", p, buf)
	}

}
