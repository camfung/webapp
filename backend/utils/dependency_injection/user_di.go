package dependency_injection

import (
	"github.com/pdftts/webapp/controllers"
	"github.com/pdftts/webapp/daos"
	daoInterfaces "github.com/pdftts/webapp/daos/interfaces"
	"github.com/pdftts/webapp/services"
	servInterfaces "github.com/pdftts/webapp/services/interfaces"
)

func InitUserDependencies() *controllers.UserController {
	var dao daoInterfaces.UserDaoInterface = daos.NewUserDao()
	var usageDao daoInterfaces.UsageDaoInterface = daos.NewUsageDao()
	var service servInterfaces.UserServiceInterface = services.NewUserService(dao, usageDao)
	return controllers.NewUserController(service)
}
