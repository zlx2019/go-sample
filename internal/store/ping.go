package store

import "gorm.io/gorm"

type PingStore struct {
	db *gorm.DB
}

// NewPingStore 提供者
func NewPingStore(db *gorm.DB) *PingStore {
	return &PingStore{db: db}
}
