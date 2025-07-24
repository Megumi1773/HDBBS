package main

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	DataBase struct {
		Url      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Dbname   string `yaml:"dbname"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
}

var appConfig *AppConfig

func InitConfig() *AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}
	appConfig = &AppConfig{}
	if err := viper.Unmarshal(appConfig); err != nil {
		log.Fatalf("Error parsing config file, %v", err)
	}
	return appConfig
}
