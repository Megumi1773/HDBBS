package main

import (
	"Backend/config"
	"Backend/router"
	"fmt"
)

func main() {
	config.GetConfig()
	r := router.InitRouters()
	port := fmt.Sprintf(":%s", config.AppConfig.App.Port)
	if port == "" {
		port = ":8080"
	}
	_ = r.Run(port)
}
