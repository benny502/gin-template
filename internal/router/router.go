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
	cors   *middleware.Cors
	user   *service.UserService
	class  *service.ClassService
	item   *service.ItemService
}

func (r *Router) Register() {
	// 登录
	r.engine.Use(r.cors.CorsMiddleware())
	api := r.engine.Group("/api")
	{
		api.POST("/login", r.user.Login)

		user := api.Group("/user").Use(r.auth.AuthMiddleware())
		{
			user.GET("/info", r.user.GetInfo)
		}

		class := api.Group("/class")
		{
			class.GET("/list", r.class.List)
			class.GET("/items", r.item.ListByClass)
		}

	}

}

func NewRouter(engine *gin.Engine, auth *middleware.Auth, cors *middleware.Cors, class *service.ClassService, item *service.ItemService, user *service.UserService, logger log.Logger) *Router {
	return &Router{
		engine: engine,
		logger: logger,
		user:   user,
		class:  class,
		item:   item,
		auth:   auth,
		cors:   cors,
	}
}
