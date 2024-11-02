package data

import (
	"bookmark/internal/biz"
	"bookmark/internal/entity"
)

type itemRepo struct {
	data *Data
}

func (i *itemRepo) ListAll() ([]*entity.Item, error) {
	result := make([]*entity.Item, 0)
	err := i.data.db.Where("is_delete = ?", 0).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *itemRepo) ListByClassId(classId int) ([]*entity.Item, error) {
	result := make([]*entity.Item, 0)
	err := i.data.db.Where("class_id = ? and is_delete = ?", classId, 0).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewItemRepo(data *Data) biz.ItemRepo {
	return &itemRepo{
		data: data,
	}
}

var _ biz.ItemRepo = (*itemRepo)(nil)
