package v1

import "bookmark/internal/api/request"

type Id struct {
	Id int `json:"id" form:"id" binding:"required"`
}

func (g *Id) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Id.required": "ID不能为空",
	}
}

type Item struct {
	Title       string `json:"title" binding:"required"`
	ClassId     int    `json:"class_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Url         string `json:"url" binding:"required,url"`
}

func (i *Item) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Title.required":       "标题不能为空",
		"ClassId.required":     "分类不能为空",
		"Description.required": "简介不能为空",
		"Url.required":         "URL不能为空",
		"Url.url":              "URL格式不正确",
	}
}

type ItemEdit struct {
	Id          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	ClassId     int    `json:"class_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Url         string `json:"url" binding:"required,url"`
}

func (i *ItemEdit) GetMessages() request.ValidatorMessages {
	return request.ValidatorMessages{
		"Id.required":          "id不能为空",
		"Title.required":       "标题不能为空",
		"ClassId.required":     "分类不能为空",
		"Description.required": "简介不能为空",
		"Url.required":         "URL不能为空",
		"Url.url":              "URL格式不正确",
	}
}
