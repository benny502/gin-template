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
	Realname  string    `gorm:"column:realname" json:"realname"`
	RoleId    int       `gorm:"column:roleId;NOT NULL" json:"roleId"`
	Email     string    `gorm:"column:email;NOT NULL" json:"email"`
	Cellphone string    `gorm:"column:cellphone;NOT NULL" json:"cellphone"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`     // 创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`     // 更新时间
	Status    int       `gorm:"column:status;default:1" json:"status"` // 1为启用，0为禁用
	IsDel     int       `gorm:"column:isDel;default:0" json:"isDel"`
	Name      string    `gorm:"column:name" json:"name"`
	Company   string    `gorm:"column:company" json:"company"`
}

func (m *User) TableName() string {
	return "m_user"
}