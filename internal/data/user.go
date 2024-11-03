package data

import (
	"bookmark/internal/biz"
	"bookmark/internal/entity"
	"bookmark/internal/pkg/gosafe"
	"bookmark/internal/pkg/log"
)

type userRepo struct {
	data   *Data
	logger log.Logger
}

func (u *userRepo) FindUserByUsername(username string) (*entity.User, error) {
	errChan := make(chan error)
	resChan := make(chan *entity.User)
	gosafe.GoSafe(func() {
		var user entity.User
		err := u.data.db.Where("username = ? and is_delete = ?", username, 0).First(&user).Error
		if err != nil {
			errChan <- err
			return
		}
		resChan <- &user
	}, u.logger)
	select {
	case user := <-resChan:
		return user, nil
	case err := <-errChan:
		return nil, err
	}
}

func (u *userRepo) FindUserById(id int) (*entity.User, error) {
	resChan := make(chan *entity.User)
	errChan := make(chan error)
	gosafe.GoSafe(func() {
		var user entity.User
		err := u.data.db.Where("id = ? and is_delete = ?", id, 0).First(&user).Error
		if err != nil {
			errChan <- err
		}
		resChan <- &user
	}, u.logger)
	select {
	case user := <-resChan:
		return user, nil
	case err := <-errChan:
		return nil, err
	}
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{data: data, logger: logger}
}

var _ biz.UserRepo = (*userRepo)(nil)
