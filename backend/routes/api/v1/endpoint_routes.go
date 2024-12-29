package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pdftts/webapp/utils/dependency_injection"
)

func SetEndpointRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitEndpointDependencies()

	group := router.Group("/endpoint")
	{
		group.GET("", controller.GetAllEndpoints)
	}
}
