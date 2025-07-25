package models

// Category 分类模型
type Category struct {
	ID   uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"unique;type:varchar(16)"`
	Icon string `gorm:"type:varchar(16)"`
}
