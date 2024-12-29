package services

import (
	iDao "github.com/pdftts/webapp/daos/interfaces"
	"github.com/pdftts/webapp/models"
)

type EndpointService struct {
	dao iDao.EndpointDaoInterface
}

func NewEndpointService(dao iDao.EndpointDaoInterface) *EndpointService {
	return &EndpointService{dao: dao}
}

func (service *EndpointService) GetAllEndpoints() ([]models.Endpoint, error) {
	return service.dao.GetAllEndpointsDao()
}
