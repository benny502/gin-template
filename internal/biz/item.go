package biz

import (
	"bookmark/internal/domain"
	"bookmark/internal/entity"
	cErr "bookmark/internal/pkg/error"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemRepo interface {
	ListAll(ctx *gin.Context) ([]*entity.Item, error)
	ListByClassId(ctx *gin.Context, classId int) ([]*entity.Item, error)
	Add(ctx *gin.Context, title string, url string, classId int, description string) (int, error)
	Update(ctx *gin.Context, id int, title string, url string, classId int, description string) error
	FindById(ctx *gin.Context, id int) (*entity.Item, error)
	Delete(ctx *gin.Context, id int) error
}

type ItemBiz struct {
	itemRepo  ItemRepo
	classRepo ClassRepo
}

func (i *ItemBiz) ListByClass(ctx *gin.Context) ([]*domain.Class, error) {
	class, err := i.classRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*domain.Class, 0)
	for _, v := range class {
		items, err := i.itemRepo.ListByClassId(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		classDomain := &domain.Class{
			ID:    v.ID,
			Icon:  v.Icon,
			Title: v.Title,
		}
		for _, item := range items {
			classDomain.Items = append(classDomain.Items, domain.Item{
				ID:          item.ID,
				Title:       item.Title,
				Url:         item.Url,
				ClassId:     item.ClassID,
				Description: item.Description,
			})
		}
		result = append(result, classDomain)

	}
	return result, nil
}

func (i *ItemBiz) Add(ctx *gin.Context, title string, url string, classId int, description string) (int, error) {
	return i.itemRepo.Add(ctx, title, url, classId, description)
}

func (i *ItemBiz) Update(ctx *gin.Context, id int, title string, url string, classId int, description string) error {
	return i.itemRepo.Update(ctx, id, title, url, classId, description)
}

func (i *ItemBiz) Delete(ctx *gin.Context, id int) error {
	return i.itemRepo.Delete(ctx, id)
}

func (i *ItemBiz) FindById(ctx *gin.Context, id int) (*domain.Item, error) {
	item, err := i.itemRepo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cErr.New(http.StatusOK, 404, "item not found")
		}
		return nil, err
	}
	return &domain.Item{
		ID:          item.ID,
		Title:       item.Title,
		Url:         item.Url,
		Description: item.Description,
		ClassId:     item.ClassID,
	}, nil
}

func NewItemBiz(itemRepo ItemRepo, classRepo ClassRepo) *ItemBiz {
	return &ItemBiz{
		itemRepo:  itemRepo,
		classRepo: classRepo,
	}
}
