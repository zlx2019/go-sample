package dao

import "gorm.io/gorm"

type PingRepo struct {
	db *gorm.DB
}

// NewPingRepo 提供者
func NewPingRepo(db *gorm.DB) *PingRepo {
	return &PingRepo{db: db}
}
