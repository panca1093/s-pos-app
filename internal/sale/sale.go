package sale

import "time"

// IUsecase defines the interface for user entities.
type IUsecase interface{}

// Sale is a concrete implementation of the IUser interface.
type (
	Sale struct {
		ID      uint    `db:"id" json:"id"`
		UserID  uint    `db:"user_id" json:"user_id"`
		Status  string  `db:"status" json:"status"`
		Payment string  `db:"payment" json:"payment"`
		Total   float64 `db:"total" json:"total"`
		Item    []Item

		CreatedAt time.Time `db:"created_at" json:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	}

	Item struct {
		ID        uint    `db:"id" json:"id"`
		SaleID    uint    `db:"sale_id" json:"sale_id"`
		ProductID uint    `db:"product_id" json:"product_id"`
		Quantity  uint    `db:"quantity" json:"quantity"`
		Price     float64 `db:"price" json:"price"`
		Subtotal  float64 `db:"subtotal" json:"subtotal"`

		CreatedAt time.Time `db:"created_at" json:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	}
)

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
