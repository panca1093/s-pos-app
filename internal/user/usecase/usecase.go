package usecase

import (
	"errors"
	"github.com/s-pos-app/internal/user"
	"github.com/s-pos-app/internal/utilities/security"
)

// IRepository defines the interface for product data persistence operations.
// It provides methods to perform CRUD operations on product entities.
type IRepository interface {
	// FindByUsername retrieves a user by their username.
	FindByUsername(username string) (*user.IUser, error)
	// Create saves a new user to the database.
	Create(user *user.IUser) error
}

type IUsecase struct {
	repo      IRepository
	jwtSecret string
}

func NewUsecase(r IRepository, secret string) IUsecase {
	return IUsecase{repo: r, jwtSecret: secret}
}

func (u *IUsecase) Register(user *user.IUser) error {
	// Check if the username already exists
	existingUser, err := u.repo.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already exists")
	}

	// Hash the password before saving`
	hashed, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	return u.repo.Create(user)
}

func (u *IUsecase) Login(username, password string) (*user.IUser, string, error) {
	// Find the user by username
	usr, err := u.repo.FindByUsername(username)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	// Check if the password is correct
	if !security.CheckPasswordHash(password, usr.Password) {
		return nil, "", errors.New("invalid credentials")
	}
	// Check if the user is active
	if !usr.IsActive {
		return nil, "", errors.New("user is not active")
	}

	// Generate JWT token
	token, err := security.GenerateJWT(usr.ID, usr.Role, usr.Username, u.jwtSecret)
	if err != nil {
		return nil, "", err
	}

	return usr, token, nil
}
