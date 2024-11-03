package router

import (
	"bookmark/internal/pkg/log"
	"bookmark/internal/service"

	"bookmark/internal/middleware/auth"
	"bookmark/internal/middleware/cache"
	"bookmark/internal/middleware/cors"
	logm "bookmark/internal/middleware/log"

	"github.com/gin-gonic/gin"
)

type Router struct {
	logger log.Logger
	engine *gin.Engine
	auth   *auth.Auth
	cors   *cors.Cors
	cache  *cache.Cache
	user   *service.UserService
	class  *service.ClassService
	item   *service.ItemService
	logm   *logm.Logger
}

func (r *Router) Register() {
	// 登录
	r.engine.Use(r.cors.CorsMiddleware())
	api := r.engine.Group("/api")
	{
		api.Use(r.cache.CacheMiddleware())
		api.Use(r.logm.LoggerMiddleware())
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

		item := api.Group("/item").Use(r.auth.AuthMiddleware())
		{
			item.POST("/add", r.item.Add)
			item.POST("/update", r.item.Update)
			item.POST("/delete", r.item.Delete)
			item.GET("/get", r.item.Get)
		}

	}

}

func NewRouter(engine *gin.Engine, auth *auth.Auth, cors *cors.Cors, logm *logm.Logger, cache *cache.Cache, class *service.ClassService, item *service.ItemService, user *service.UserService, logger log.Logger) *Router {
	return &Router{
		engine: engine,
		logger: logger,
		user:   user,
		class:  class,
		item:   item,
		auth:   auth,
		cors:   cors,
		cache:  cache,
		logm:   logm,
	}
}
