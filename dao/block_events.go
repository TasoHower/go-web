package dao

import (
	"time"

	"gorm.io/gorm"
)

// 更新方式：insert & update
type Demo struct {
	ID                 uint       `gorm:"column:id;primaryKey;->"`
	CreatedAt          time.Time  `gorm:"column:created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at"`
}

const demoTableName = "table_name"

func (c *Demo) TableName() string {
	return demoTableName
}

func CreateHolder(db *gorm.DB, holder Demo) (Demo, error) {
	err := db.Model(Demo{}).Create(&holder).Error
	return holder, err
}
