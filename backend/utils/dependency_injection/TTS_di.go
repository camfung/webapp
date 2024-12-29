package dependency_injection

import (
	"github.com/pdftts/webapp/controllers"
	"github.com/pdftts/webapp/daos"
	daoInterfaces "github.com/pdftts/webapp/daos/interfaces"
	"github.com/pdftts/webapp/services"
	servInterfaces "github.com/pdftts/webapp/services/interfaces"
)

func InitTTSDependencies() *controllers.TTSController {
	var dao daoInterfaces.TTSDaoInterface = daos.NewTTSDao()
	var service servInterfaces.TTSServiceInterface = services.NewTTSService(dao)

	return controllers.NewTTSController(service)
}
