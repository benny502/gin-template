package router

import (
	"bookmark/internal/middleware"
	"bookmark/internal/pkg/log"
	"bookmark/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	logger log.Logger
	engine *gin.Engine
	auth   *middleware.Auth
	user   *service.UserService
}

func (r *Router) Register() {
	// // 登录
	r.engine.GET("/login", r.user.Login)

	user := r.engine.Group("/user").Use(r.auth.AuthMiddleware())
	user.GET("/info", r.user.GetInfo)

}

func NewRouter(engine *gin.Engine, auth *middleware.Auth, user *service.UserService, logger log.Logger) *Router {
	return &Router{
		engine: engine,
		logger: logger,
		user:   user,
		auth:   auth,
	}
}
