package biz

import (
	"bookmark/internal/domain"
	"bookmark/internal/entity"
)

type ItemRepo interface {
	ListAll() ([]*entity.Item, error)
	ListByClassId(classId int) ([]*entity.Item, error)
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

func NewItemBiz(itemRepo ItemRepo, classRepo ClassRepo) *ItemBiz {
	return &ItemBiz{
		itemRepo:  itemRepo,
		classRepo: classRepo,
	}
}
