package user

import "time"

// IUsecase defines the interface for user entities.
type IUsecase interface {
	// Register creates a new user.
	Register(user *IUser) error
	// Login authenticates a user and returns a JWT token.
	Login(username, password string) (*IUser, string, error)
}

// IUser is a concrete implementation of the IUser interface.
type IUser struct {
	ID         uint   `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Username   string `db:"username" json:"username"`
	Password   string `db:"password" json:"-"`
	Role       string `db:"role" json:"role"`
	Email      string `db:"email" json:"email"`
	Phone      string `db:"phone" json:"phone"`
	Address    string `db:"address" json:"address"`
	ProfilePic string `db:"profile_pic" json:"profile_pic"`
	IsActive   bool   `db:"is_active" json:"is_active"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
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
