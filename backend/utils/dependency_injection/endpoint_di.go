package dependency_injection

import (
	"github.com/pdftts/webapp/controllers"
	"github.com/pdftts/webapp/daos"
	daoInterfaces "github.com/pdftts/webapp/daos/interfaces"
	"github.com/pdftts/webapp/services"
	servInterfaces "github.com/pdftts/webapp/services/interfaces"
)

func InitEndpointDependencies() *controllers.EndpointController {
	var dao daoInterfaces.EndpointDaoInterface = daos.NewEndpointDao()
	var service servInterfaces.EndpointServiceInterface = services.NewEndpointService(dao)

	return controllers.NewEndpointController(service)
}
