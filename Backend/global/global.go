package global

import "gorm.io/gorm"

// Db 全局数据库对象
var (
	Db *gorm.DB
)
