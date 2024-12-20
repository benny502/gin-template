package biz

import (
	"bookmark/internal/config"
	"bookmark/internal/domain"
	"bookmark/internal/entity"
	cErr "bookmark/internal/pkg/error"
	"bookmark/internal/pkg/jwt"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindUserById(ctx *gin.Context, id int) (*entity.User, error)
	FindUserByUsername(ctx *gin.Context, username string) (*entity.User, error)
}

type UserBiz struct {
	userRepo UserRepo
	conf     *config.Configuration
}

func (u *UserBiz) Login(ctx *gin.Context, username string, password string) (*domain.User, error) {
	user, err := u.userRepo.FindUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cErr.New(http.StatusOK, 404, "用户不存在")
		}
		return nil, err
	}
	if user.Password != password {
		return nil, cErr.New(http.StatusOK, 422, "密码错误")
	}
	return &domain.User{
		ID:       user.ID,
		Username: user.Username,
		Token:    jwt.GenerateToken(user.ID, u.conf.Jwt.JwtKey, u.conf.Jwt.Issuer, time.Duration(u.conf.Jwt.Expttl)*time.Second),
	}, nil
}

func (u *UserBiz) GetInfo(ctx *gin.Context, id int) (*domain.User, error) {
	user, err := u.userRepo.FindUserById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cErr.New(http.StatusOK, 404, "用户不存在")
		}
	}
	return &domain.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func NewUserBiz(userRepo UserRepo, conf *config.Configuration) *UserBiz {
	return &UserBiz{
		userRepo: userRepo,
		conf:     conf,
	}
}
