package v1

import "bookmark/internal/api/request"

type Login struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password"  binding:"required"`
}

func (l *Login) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
	}
}

type GetInfo struct {
	ID int `json:"id" form:"id" binding:"required"`
}

func (g *GetInfo) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"ID.required": "用户ID不能为空",
	}
}
