package main

import (
	"fmt"
)

func main() {
	InitConfig()
	InitDB()
	_ = db.AutoMigrate(&User{})
	_ = db.AutoMigrate(&Gender{})
	_ = db.AutoMigrate(&Category{})
	r := InitRouters()

	port := fmt.Sprintf(":%s", appConfig.Server.Port)
	if port == "" {
		port = ":8080"
	}
	_ = r.Run(port)
}
