package dao

import "gorm.io/gorm"

// ExampleRepo Example 数据访问层
type ExampleRepo struct {
	db *gorm.DB
}

// NewExampleRepo 提供者
func NewExampleRepo(db *gorm.DB) *ExampleRepo {
	return &ExampleRepo{db: db}
}