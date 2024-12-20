package biz

import (
	"bookmark/internal/domain"
	"bookmark/internal/entity"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClassRepo interface {
	ListAll(ctx *gin.Context) ([]*entity.Class, error)
}

type ClassBiz struct {
	classRepo ClassRepo
}

func (c *ClassBiz) List(ctx *gin.Context) ([]*domain.Class, error) {
	list, err := c.classRepo.ListAll(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return make([]*domain.Class, 0), nil
		}
		return nil, err
	}
	result := make([]*domain.Class, 0)
	for _, v := range list {
		result = append(result, &domain.Class{
			ID:    v.ID,
			Title: v.Title,
			Icon:  v.Icon,
		})
	}
	return result, nil
}

func NewClassBiz(classRepo ClassRepo) *ClassBiz {
	return &ClassBiz{
		classRepo: classRepo,
	}
}
