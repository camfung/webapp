package daos

import (
	"github.com/pdftts/webapp/models"
	"github.com/pdftts/webapp/utils/database"
)

type EndpointDao struct{}

func NewEndpointDao() *EndpointDao {
	return &EndpointDao{}
}

func (dao *EndpointDao) GetAllEndpointsDao() ([]models.Endpoint, error) {
	db := database.GetInstance()

	var usage []models.Endpoint

	if err := db.Find(&usage).Error; err != nil {
		return nil, err
	}

	return usage, nil
}
