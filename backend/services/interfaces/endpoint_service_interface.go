package interfaces

import (
	"github.com/pdftts/webapp/models"
)

type EndpointServiceInterface interface {
	GetAllEndpoints() ([]models.Endpoint, error)
}
