package biz

import (
	"bookmark/internal/domain"
	"bookmark/internal/entity"
	cErr "bookmark/internal/pkg/error"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type ItemRepo interface {
	ListAll() ([]*entity.Item, error)
	ListByClassId(classId int) ([]*entity.Item, error)
	Add(title string, url string, classId int, description string) (int, error)
	Update(id int, title string, url string, classId int, description string) error
	FindById(id int) (*entity.Item, error)
	Delete(id int) error
}

type ItemBiz struct {
	itemRepo  ItemRepo
	classRepo ClassRepo
}

func (i *ItemBiz) ListByClass() ([]*domain.Class, error) {
	class, err := i.classRepo.ListAll()
	if err != nil {
		return nil, err
	}
	result := make([]*domain.Class, 0)
	for _, v := range class {
		items, err := i.itemRepo.ListByClassId(v.ID)
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

func (i *ItemBiz) Add(title string, url string, classId int, description string) (int, error) {
	return i.itemRepo.Add(title, url, classId, description)
}

func (i *ItemBiz) Update(id int, title string, url string, classId int, description string) error {
	return i.itemRepo.Update(id, title, url, classId, description)
}

func (i *ItemBiz) Delete(id int) error {
	return i.itemRepo.Delete(id)
}

func (i *ItemBiz) FindById(id int) (*domain.Item, error) {
	item, err := i.itemRepo.FindById(id)
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
