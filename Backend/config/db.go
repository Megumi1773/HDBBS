package config

import (
	"Backend/global"
	"Backend/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func initDB() {
	var err error
	c := AppConfig.Database
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.Username, c.Password, c.Url, c.Port, c.Dbname)
	if db, err = gorm.Open(mysql.Open(dns), &gorm.Config{SkipDefaultTransaction: true}); err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)
	global.Db = db
	migrateDB()
}
func migrateDB() {
	if err := global.Db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}
	if err := global.Db.AutoMigrate(&models.Category{}); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}
}
