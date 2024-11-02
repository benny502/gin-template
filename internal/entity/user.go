// Code generated by sql2gorm. DO NOT EDIT.
package entity

import (
	"time"
)

// 系统用户表
type User struct {
	ID        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username  string    `gorm:"column:username;NOT NULL" json:"username"`
	Password  string    `gorm:"column:password;NOT NULL" json:"password"`
	RoleId    int       `gorm:"column:roleId;NOT NULL" json:"roleId"`
	Email     string    `gorm:"column:email;NOT NULL" json:"email"`
	Cellphone string    `gorm:"column:cellphone;NOT NULL" json:"cellphone"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`   // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`   // 更新时间
	Status    int       `gorm:"column:status;default:1" json:"status"` // 1为启用，0为禁用
	IsDelete  int       `gorm:"column:is_delete;default:0" json:"is_delete"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (m *User) TableName() string {
	return "m_user"
}