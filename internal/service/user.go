package service

import (
	"bookmark/internal/api/request"
	"bookmark/internal/api/response"
	v1 "bookmark/internal/api/v1"
	"bookmark/internal/biz"
	"bookmark/internal/config"
	"bookmark/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	logger   log.Logger
	userRepo *biz.UserBiz
}

func (s *UserService) Login(context *gin.Context) {

	var form v1.Login
	if err := context.ShouldBindBodyWithJSON(&form); err != nil {
		response.FailByErr(context, request.GetError(&form, err))
		return
	}

	u, err := s.userRepo.Login(form.Username, form.Password)
	if err != nil {
		response.FailByErr(context, err)
		return
	}

	response.Success(context, u)
}

func (s *UserService) GetInfo(context *gin.Context) {

	id := context.GetInt("id")
	u, err := s.userRepo.GetInfo(id)
	if err != nil {
		response.FailByErr(context, err)
		return
	}

	response.Success(context, u)
}

func NewUserService(logger log.Logger, userRepo *biz.UserBiz, conf *config.Configuration) *UserService {
	return &UserService{
		logger:   logger,
		userRepo: userRepo,
	}
}
