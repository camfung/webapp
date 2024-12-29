package dependency_injection

import (
	"github.com/pdftts/webapp/controllers"
	"github.com/pdftts/webapp/daos"
	iDao "github.com/pdftts/webapp/daos/interfaces"
	"github.com/pdftts/webapp/services"
	iServices "github.com/pdftts/webapp/services/interfaces"
)

func InitAuthDependencies() controllers.AuthController {
	var userDao iDao.UserDaoInterface = daos.NewUserDao()
	var usageDao iDao.UsageDaoInterface = daos.NewUsageDao()
	var userService iServices.UserServiceInterface = services.NewUserService(userDao, usageDao)
	var authDao iDao.AuthDaoInterface = daos.NewAuthDao()
	var authService iServices.AuthServiceInterface = services.NewAuthService(authDao, userService)
	var authController controllers.AuthController = *controllers.NewAuthController(authService, userService)
	return authController
}
