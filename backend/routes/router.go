package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pdftts/webapp/daos"
	iDao "github.com/pdftts/webapp/daos/interfaces"
	"github.com/pdftts/webapp/middlewares"
	v1 "github.com/pdftts/webapp/routes/api/v1"
	"github.com/pdftts/webapp/services"
	iServices "github.com/pdftts/webapp/services/interfaces"
	"github.com/pdftts/webapp/utils/database"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Setup middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORS())

	// Setup public routes for v1
	v1RouterGroup := router.Group("/api/v1")
	{
		v1.SetAuthRoutes(v1RouterGroup)
	}

	var userDao iDao.UserDaoInterface = daos.NewUserDao()
	var usageDao iDao.UsageDaoInterface = daos.NewUsageDao()
	var userService iServices.UserServiceInterface = services.NewUserService(userDao, usageDao)
	var authDao iDao.AuthDaoInterface = daos.NewAuthDao()
	var authService iServices.AuthServiceInterface = services.NewAuthService(authDao, userService)

	// Setup private routes (requires authentication)
	privateRouterGroup := v1RouterGroup.Group("")
	privateRouterGroup.Use(middlewares.Auth(authService))
	privateRouterGroup.Use() // Add usage tracking middleware
	{
		v1.SetUserRoutes(privateRouterGroup)
		v1.SetEndpointRoutes(privateRouterGroup)

		// Setup routes that count towards api usage total
		usageTrackingRouterGroup := privateRouterGroup.Group("")
		usageTrackingRouterGroup.Use(middlewares.UsageTrackingMiddleware(database.GetInstance()))
		{
			v1.SetTTSRoutes(usageTrackingRouterGroup)
		}
	}

	return router
}
