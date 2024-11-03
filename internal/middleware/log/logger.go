package log

import (
	"bookmark/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

type Logger struct {
	log log.Logger
}

func (l *Logger) LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", l.log)
	}
}

func NewLogger(logger log.Logger) *Logger {
	return &Logger{
		log: logger,
	}
}

func WithContext(ctx *gin.Context) log.Logger {
	return ctx.MustGet("logger").(log.Logger)
}
