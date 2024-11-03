package auth

import (
	"bookmark/internal/api/response"
	"bookmark/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	auth "bookmark/internal/pkg/jwt"

	cErr "bookmark/internal/pkg/error"
)

type Auth struct {
	conf *config.Configuration
}

func (a *Auth) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.FailByErr(c, cErr.Unauthorized("missing Authorization header"))
			return
		}
		//解析token
		token, err := jwt.ParseWithClaims(tokenStr, &auth.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(a.conf.Jwt.JwtKey), nil
		})
		if err != nil {
			response.FailByErr(c, cErr.Unauthorized("登录授权已失效"))
			return
		}

		claims := token.Claims.(*auth.CustomClaims)

		c.Set("id", claims.ID)
		c.Set("token", token)
	}
}

func NewAuth(conf *config.Configuration) *Auth {
	return &Auth{conf: conf}
}
