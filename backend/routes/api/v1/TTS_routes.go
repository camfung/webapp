package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pdftts/webapp/utils/dependency_injection"
)

func SetTTSRoutes(router *gin.RouterGroup) {
	controller := dependency_injection.InitTTSDependencies()

	group := router.Group("/tts")
	{
		group.POST("pdf-to-text", controller.PdfToText)
	}
}
