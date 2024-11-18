package data

import (
	"bookmark/internal/biz"
	"bookmark/internal/entity"
	"bookmark/internal/pkg/gosafe"
	"bookmark/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

type itemRepo struct {
	data   *Data
	logger log.Logger
}

func (i *itemRepo) ListAll(ctx *gin.Context) ([]*entity.Item, error) {
	resChan := make(chan []*entity.Item)
	errChan := make(chan error)
	gosafe.GoSafe(ctx, func(ctx *gin.Context) {
		result := make([]*entity.Item, 0)
		err := i.data.db.Where("is_delete = ?", 0).Find(&result).Error
		if err != nil {
			errChan <- err
		}
		resChan <- result
	})
	select {
	case result := <-resChan:
		return result, nil
	case err := <-errChan:
		return nil, err
	}
}

func (i *itemRepo) ListByClassId(ctx *gin.Context, classId int) ([]*entity.Item, error) {
	resChan := make(chan []*entity.Item)
	errChan := make(chan error)
	gosafe.GoSafe(ctx, func(ctx *gin.Context) {
		result := make([]*entity.Item, 0)
		err := i.data.db.Where("class_id = ? and is_delete = ?", classId, 0).Find(&result).Error
		if err != nil {
			errChan <- err
		}
		resChan <- result
	})
	select {
	case result := <-resChan:
		return result, nil
	case err := <-errChan:
		return nil, err
	}
}

func (i *itemRepo) Add(ctx *gin.Context, title string, url string, classId int, description string) (int, error) {
	resChan := make(chan int)
	errChan := make(chan error)
	gosafe.GoSafe(ctx, func(ctx *gin.Context) {
		item := &entity.Item{
			Title:       title,
			Url:         url,
			ClassID:     classId,
			Description: description,
		}
		err := i.data.db.Create(item).Error
		if err != nil {
			errChan <- err
			return
		}
		resChan <- item.ID
	})
	select {
	case id := <-resChan:
		return id, nil
	case err := <-errChan:
		return 0, err
	}
}

func (i *itemRepo) FindById(ctx *gin.Context, id int) (*entity.Item, error) {
	resChan := make(chan *entity.Item)
	errChan := make(chan error)
	gosafe.GoSafe(ctx, func(ctx *gin.Context) {
		var item entity.Item
		err := i.data.db.Where("id = ? and is_delete = ?", id, 0).First(&item).Error
		if err != nil {
			errChan <- err
			return
		}
		resChan <- &item
	})
	select {
	case item := <-resChan:
		return item, nil
	case err := <-errChan:
		return nil, err
	}
}

func (i *itemRepo) Update(ctx *gin.Context, id int, title string, url string, classId int, description string) error {
	errChan := make(chan error)
	gosafe.GoSafe(ctx, func(ctx *gin.Context) {
		errChan <- i.data.db.Save(&entity.Item{
			ID:          id,
			Title:       title,
			Url:         url,
			ClassID:     classId,
			Description: description,
		}).Error
	})
	return <-errChan
}

func (i *itemRepo) Delete(ctx *gin.Context, id int) error {
	errChan := make(chan error)
	gosafe.GoSafe(ctx, func(*gin.Context) {
		var item entity.Item
		errChan <- i.data.db.Model(&item).Where("id = ?", id).Update("is_delete", 1).Error
	})
	return <-errChan
}

func NewItemRepo(data *Data, logger log.Logger) biz.ItemRepo {
	return &itemRepo{
		data:   data,
		logger: logger,
	}
}

var _ biz.ItemRepo = (*itemRepo)(nil)
