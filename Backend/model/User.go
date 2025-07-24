package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"gorm.Model"`
	Username   string `gorm:"type:varchar(32);unique" json:"username"`
	Password   string `gorm:"type:varchar(32)" json:"password"`
	Avatar     string `gorm:"type:varchar(32)" json:"avatar"`
	Email      string `gorm:"type:varchar(32);unique" json:"email"`
	Phone      string `gorm:"type:varchar(32);unique" json:"phone"`
	Gender     uint8  `gorm:"type:tinyint(1)" json:"gender"`
	Status     uint8  `gorm:"type:tinyint(1)" json:"status"`
}

func (*User) TableName() string {
	return "users"
}

type Gender struct {
	gorm.Model
	Sex string `gorm:"type:varchar(8)"`
}

func (*Gender) TableName() string {
	return "genders"
}
