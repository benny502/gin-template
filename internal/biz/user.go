package biz

import (
	"bookmark/internal/config"
	"bookmark/internal/domain"
	"bookmark/internal/entity"
	"bookmark/internal/pkg/jwt"
	"errors"
	"time"
)

type UserRepo interface {
	FindUserById(id int) (*entity.User, error)
	FindUserByUsername(username string) (*entity.User, error)
}

type UserBiz struct {
	userRepo UserRepo
	conf     *config.Configuration
}

func (u *UserBiz) Login(username string, password string) (*domain.User, error) {
	user, err := u.userRepo.FindUserByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	return &domain.User{
		ID:       user.ID,
		Username: user.Username,
		Token:    jwt.GenerateToken(user.ID, u.conf.Jwt.JwtKey, u.conf.Jwt.Issuer, time.Duration(u.conf.Jwt.Expttl)*time.Second),
	}, nil
}

func (u *UserBiz) GetInfo(id int) (*domain.User, error) {
	user, err := u.userRepo.FindUserById(id)
	if err != nil {
		return nil, errors.New("用户不存在")
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