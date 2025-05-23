package repository

import (
	"github.com/s-pos-app/internal/user"
	"gorm.io/gorm"
)

type IRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *IRepository {
	return &IRepository{db}
}

func (r *IRepository) FindByUsername(username string) (*user.IUser, error) {
	var u user.IUser
	if err := r.db.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *IRepository) Create(u *user.IUser) error {
	return r.db.Create(u).Error
}
