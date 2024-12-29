package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pdftts/webapp/utils/dependency_injection"
)

func SetUserRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitUserDependencies()

	group := router.Group("/user")
	{
		group.GET("", controller.GetAllUsersHandler)
		group.GET("/:id", controller.GetUserHandler)
		group.GET("/current", controller.GetCurrentUserHandler)
		group.GET("/:id/usage", controller.GetUserUsageHandler)
		//group.POST("/", controller.CreateUserHandler)
		//group.PUT("/:id", controller.UpdateUserHandler)
		//group.PUT("/current", controller.UpdateCurrentUserHandler)
		//group.DELETE("/:id", controller.DeleteUserHandler)
	}
}
