package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	c := appConfig.DataBase
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.Username, c.Password, c.Url, c.Port, c.Dbname)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}
}
