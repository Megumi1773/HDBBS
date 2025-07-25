package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique not null;type:varchar(16)" json:"username"`
	Password string `gorm:"not null;type:varchar(16)" json:"password"`
}
