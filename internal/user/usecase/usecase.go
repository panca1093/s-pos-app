package usecase

// IRepository defines the interface for product data persistence operations.
// It provides methods to perform CRUD operations on product entities.
type IRepository interface{}

type IUsecase struct {
	repo IRepository
}

func NewUsecase(r IRepository) IUsecase {
	return IUsecase{repo: r}
}
