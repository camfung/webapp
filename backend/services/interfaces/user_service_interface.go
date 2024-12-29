package interfaces

import "github.com/pdftts/webapp/models"

type UserServiceInterface interface {
	GetAllUsers(includeDeleted bool, full bool) ([]models.User, error)
	GetUser(id int, includeDeleted bool, full bool) (*models.User, error)
	GetUserByEmail(email string, includeDeleted bool, full bool) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(updatedUser *models.User) (*models.User, error)
	DeleteUser(id int, softDelete bool) error
	GetUserUsageByUserId(userId int) ([]models.UserEndpointUsage, error)
}
