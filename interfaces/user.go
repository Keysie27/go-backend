package interfaces

import "github.com/Keysie27/go-backend/models"

type UsersRepo interface {
	GetUserById(int64) (*models.User, error)
	CreateUser(*models.User) (int64, error)
	GetUserAddresses(int64) ([]*models.Address, error)
	UpdateUser(*models.User) error
	GetAllUsers() ([]*models.User, error)
}
