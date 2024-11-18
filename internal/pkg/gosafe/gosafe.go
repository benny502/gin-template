package gosafe

import (
	"github.com/gin-gonic/gin"
)

func GoSafe(ctx *gin.Context, fn func(ctx *gin.Context)) {
	go RunSafe(ctx, fn)
}

func RunSafe(ctx *gin.Context, fn func(ctx *gin.Context)) {
	defer Recovery(ctx)

	fn(ctx)
}
