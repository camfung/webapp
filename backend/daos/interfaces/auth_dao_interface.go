package interfaces

import "github.com/pdftts/webapp/models"

type AuthDaoInterface interface {
	RegisterUser(user models.User) (*models.User, error)
}
