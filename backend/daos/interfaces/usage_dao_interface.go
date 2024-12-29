package interfaces

import "github.com/pdftts/webapp/models"

type UsageDaoInterface interface {
	GetUsageByUserId(userId int) ([]models.UserEndpointUsage, error)
}
