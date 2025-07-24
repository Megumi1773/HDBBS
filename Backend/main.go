package main

import (
	"Backend/model"
	"fmt"
)

func main() {
	InitConfig()
	fmt.Println(appConfig)
	InitDB()
	fmt.Println(db)
	_ = db.AutoMigrate(&model.User{})
}
