package product

import (
	"context"
	"time"
)

type IProduct struct {
	ID          uint64    `db:"id" json:"id"`
	Pictures    string    `db:"pictures" json:"pictures"`
	Category    string    `db:"category" json:"category"`
	Description string    `db:"description" json:"description"`
	Name        string    `db:"name" json:"name"`
	Price       float64   `db:"price" json:"price"`
	Stock       int       `db:"stock" json:"stock"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// IUsecase defines the business logic interface for product operations.
// It provides methods to handle product-related use cases in the application.
type IUsecase interface {
	GetAll(ctx context.Context) ([]IProduct, error)
	Create(ctx context.Context, p IProduct) (IProduct, error)
	GetByID(ctx context.Context, id uint64) (IProduct, error)
}

var defaultUsecase IUsecase

func Init(uc IUsecase) {
	// Initialize the default use case with a concrete implementation.
	// This is where you would typically set up your repository and other dependencies.
	defaultUsecase = uc
}

// GetDefaultUsecase returns the default use case implementation for product operations.
// It is used to access the product use case methods without needing to create a new instance.
func GetDefaultUsecase() IUsecase {
	return defaultUsecase
}
