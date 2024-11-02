package service

import (
	"bookmark/internal/api/response"
	"bookmark/internal/biz"
	"bookmark/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

type ClassService struct {
	logger log.Logger
	biz    *biz.ClassBiz
}

func (s *ClassService) List(c *gin.Context) {
	list, err := s.biz.List()
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	response.Success(c, list)
}

func NewClassService(logger log.Logger, biz *biz.ClassBiz) *ClassService {
	return &ClassService{
		logger: logger,
		biz:    biz,
	}
}
