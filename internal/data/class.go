package data

import (
	"bookmark/internal/biz"
	"bookmark/internal/entity"
)

type classRepo struct {
	data *Data
}

func (c *classRepo) ListAll() ([]*entity.Class, error) {
	var class []*entity.Class
	err := c.data.db.Where("is_delete = ?", 0).Find(&class).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}

func NewClassRepo(data *Data) biz.ClassRepo {
	return &classRepo{
		data: data,
	}
}
