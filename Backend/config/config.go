package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Url         string `yaml:"url"`
		Port        string `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Dbname      string `yaml:"dbname"`
		MaxIdleConn int    `yaml:"max_idle_conn"`
		MaxOpenConn int    `yaml:"max_open_conn"`
	} `yaml:"database"`
	Jwt struct {
		Key string `yaml:"key"`
		Exp int    `yaml:"exp"`
	}
}

var AppConfig *Config

func GetConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	AppConfig = &Config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cofing reading err: %v", err)
	}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("cofing unmarshal err: %v", err)
	}
	initDB()
}
