package interfaces

import "github.com/pdftts/webapp/models"

type EndpointDaoInterface interface {
	GetAllEndpointsDao() ([]models.Endpoint, error)
}
