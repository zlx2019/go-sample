package store

import "gorm.io/gorm"

// ExampleStore Example 数据访问层
type ExampleStore struct {
	db *gorm.DB
}

// NewExampleStore 提供者
func NewExampleStore(db *gorm.DB) *ExampleStore {
	return &ExampleStore{db: db}
}
