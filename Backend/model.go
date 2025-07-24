package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model `json:"gorm.Model"`
	Username   string `gorm:"type:varchar(32);unique" json:"username"`
	Password   string `gorm:"type:varchar(32)" json:"password"`
	Avatar     string `gorm:"type:varchar(32)" json:"avatar"`
	Email      string `gorm:"type:varchar(32);unique" json:"email"`
	Phone      string `gorm:"type:varchar(32);unique" json:"phone"`
	Gender     uint8  `gorm:"type:tinyint(1); foreignKey:sex" json:"gender"`
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

type Category struct {
	ID           uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	CategoryName string `gorm:"type:varchar(16)"`
	CategoryIcon string `gorm:"type:varchar(16)"`
}

func (Category) TableName() string {
	return "category"
}
func GetAllCategory(c *gin.Context) {
	var category []Category
	if res := db.Find(&category); res.Error != nil {
		log.Fatalf("查询失败：%v", res.Error)
	}
	Respond(c, http.StatusOK, 200, "获取成功", category)
}
