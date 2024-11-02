package data

import (
	"bookmark/internal/biz"
	"bookmark/internal/entity"
)

type userRepo struct {
	data *Data
}

func (u *userRepo) FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := u.data.db.Where("username = ? and is_delete = ?", username, 0).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) FindUserById(id int) (*entity.User, error) {
	var user entity.User
	err := u.data.db.Where("id = ? and is_delete = ?", id, 0).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{data: data}
}

var _ biz.UserRepo = (*userRepo)(nil)
