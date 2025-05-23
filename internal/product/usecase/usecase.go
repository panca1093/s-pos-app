package usecase

import (
	"context"
	"github.com/s-pos-app/internal/product"
)

// IRepository defines the interface for product data persistence operations.
// It provides methods to perform CRUD operations on product entities.
type IRepository interface {
	// GetByID retrieves a product by its ID from the repository.
	GetByID(ctx context.Context, id uint64) (product.IProduct, error)
	// GetAll retrieves all available products from the repository.
	GetAll(ctx context.Context) ([]product.IProduct, error)
	// Create stores a new product in the repository and returns the created product.
	Create(ctx context.Context, product product.IProduct) (product.IProduct, error)
}

type IUsecase struct {
	repo IRepository
}

func NewUsecase(r IRepository) IUsecase {
	return IUsecase{repo: r}
}

func (u *IUsecase) GetAll(ctx context.Context) ([]product.IProduct, error) {
	return u.repo.GetAll(ctx)
}

func (u *IUsecase) Create(ctx context.Context, p product.IProduct) (product.IProduct, error) {
	return u.repo.Create(ctx, p)
}

func (u *IUsecase) GetByID(ctx context.Context, id uint64) (product.IProduct, error) {
	return u.repo.GetByID(ctx, id)
}
