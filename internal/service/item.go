package service

import (
	"bookmark/internal/api/response"
	"bookmark/internal/biz"
	"bookmark/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

type ItemService struct {
	logger  log.Logger
	itemBiz *biz.ItemBiz
}

func (i *ItemService) ListByClass(c *gin.Context) {
	list, err := i.itemBiz.ListByClass()
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	response.Success(c, list)
}

func NewItemService(logger log.Logger, itemBiz *biz.ItemBiz) *ItemService {
	return &ItemService{logger: logger, itemBiz: itemBiz}
}
