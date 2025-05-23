package repository

import "gorm.io/gorm"

type IRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *IRepository {
	return &IRepository{db}
}
