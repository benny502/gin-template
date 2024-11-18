package service

import (
	"bookmark/internal/api/request"
	"bookmark/internal/api/response"
	v1 "bookmark/internal/api/v1"
	"bookmark/internal/biz"
	"bookmark/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

type ItemService struct {
	logger  log.Logger
	itemBiz *biz.ItemBiz
}

func (i *ItemService) ListByClass(c *gin.Context) {
	list, err := i.itemBiz.ListByClass(c)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	response.Success(c, list)
}

func (i *ItemService) Add(c *gin.Context) {
	var item v1.Item
	err := request.ShouldBindWithJSON(c, &item)
	if err != nil {
		response.FailByErr(c, request.GetError(&item, err))
		return
	}
	id, err := i.itemBiz.Add(c, item.Title, item.Url, item.ClassId, item.Description)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	response.Success(c, id)
}

func (i *ItemService) Update(c *gin.Context) {
	var item v1.ItemEdit
	err := request.ShouldBindWithJSON(c, &item)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	err = i.itemBiz.Update(c, item.Id, item.Title, item.Url, item.ClassId, item.Description)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	response.Success(c, &response.EmptyBody{})
}

func (i *ItemService) Delete(c *gin.Context) {
	var req v1.Id
	err := request.ShouldBindWithJSON(c, &req)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	err = i.itemBiz.Delete(c, req.Id)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	response.Success(c, &response.EmptyBody{})
}

func (i *ItemService) Get(c *gin.Context) {
	var req v1.Id
	err := request.ShouldBindWithQuery(c, &req)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	item, err := i.itemBiz.FindById(c, req.Id)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	response.Success(c, item)
}

func NewItemService(logger log.Logger, itemBiz *biz.ItemBiz) *ItemService {
	return &ItemService{logger: logger, itemBiz: itemBiz}
}
