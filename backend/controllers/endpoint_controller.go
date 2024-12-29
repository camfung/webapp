package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pdftts/webapp/services/interfaces"
)

type EndpointController struct {
	service interfaces.EndpointServiceInterface
}

func NewEndpointController(service interfaces.EndpointServiceInterface) *EndpointController {
	return &EndpointController{
		service: service,
	}
}

func (contr *EndpointController) GetAllEndpoints(c *gin.Context) {
	// call the Service
	content, err := contr.service.GetAllEndpoints()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to get endpoints. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}
