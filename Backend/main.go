package main

import (
	"Backend/config"
	"Backend/routers"
	"fmt"
)

func main() {
	config.GetConfig()
	r := routers.InitRouters()

	port := fmt.Sprintf(":%s", config.AppConfig.App.Port)
	if port == "" {
		port = ":8080"
	}
	_ = r.Run(port)
}
