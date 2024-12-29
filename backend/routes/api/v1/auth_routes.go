package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pdftts/webapp/middlewares"
	"github.com/pdftts/webapp/utils/dependency_injection"
)

func SetAuthRoutes(router *gin.RouterGroup) {
	authController := dependency_injection.InitAuthDependencies()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.RegisterUser)
		authGroup.POST("/login", authController.LoginUser)
		authGroup.GET("/logout", authController.LogoutUser)
		authGroup.GET("/test", middlewares.Auth(authController.Service), authController.TestAuthMiddleware)
	}
}
