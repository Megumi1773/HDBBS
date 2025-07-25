package models

import "gorm.io/gorm"

// User 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"unique not null;type:varchar(255)" json:"username"`
	Password string `gorm:"not null;type:varchar(255)" json:"password"`
	Avatar   string `gorm:"not null;type:varchar(255)" json:"avatar"`
	NickName string `gorm:"not null;type:varchar(255)" json:"nickname"`
}
type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UpdateUser struct {
	NickName string `gorm:"not null;type:varchar(255)" json:"nickname,omitempty"`
	Avatar   string `gorm:"not null;type:varchar(255)" json:"avatar,omitempty"`
}
