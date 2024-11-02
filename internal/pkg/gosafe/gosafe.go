package gosafe

import "bookmark/internal/pkg/log"

func GoSafe(fn func(), logger log.Logger) {
	go RunSafe(fn, logger)
}

func RunSafe(fn func(), logger log.Logger) {
	defer Recovery(logger)

	fn()
}
