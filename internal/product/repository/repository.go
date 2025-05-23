package repository

import (
	"context"
	"github.com/s-pos-app/internal/product"
	"gorm.io/gorm"
)

type IRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *IRepository {
	return &IRepository{db}
}

func (r *IRepository) GetByID(ctx context.Context, id uint64) (product.IProduct, error) {
	var products product.IProduct

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&products).Error
	return products, err
}

func (r *IRepository) GetAll(ctx context.Context) ([]product.IProduct, error) {
	var products []product.IProduct

	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err
}

func (r *IRepository) Create(ctx context.Context, p *product.IProduct) error {
	return r.db.WithContext(ctx).Create(p).Error
}
