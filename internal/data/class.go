package data

import (
	"bookmark/internal/biz"
	"bookmark/internal/entity"
	"bookmark/internal/pkg/gosafe"
	"bookmark/internal/pkg/log"
)

type classRepo struct {
	data   *Data
	logger log.Logger
}

func (c *classRepo) ListAll() ([]*entity.Class, error) {
	errChan := make(chan error)
	resChan := make(chan []*entity.Class)
	gosafe.GoSafe(func() {
		result := make([]*entity.Class, 0)
		err := c.data.db.Where("is_delete = ?", 0).Find(&result).Error
		if err != nil {
			errChan <- err
		}
		resChan <- result
	}, c.logger)
	select {
	case result := <-resChan:
		return result, nil
	case err := <-errChan:
		return nil, err
	}
}

func NewClassRepo(data *Data, logger log.Logger) biz.ClassRepo {
	return &classRepo{
		data:   data,
		logger: logger,
	}
}
