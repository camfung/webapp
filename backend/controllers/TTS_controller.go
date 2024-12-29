package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdftts/webapp/models"
	"github.com/pdftts/webapp/services/interfaces"
)

type TTSController struct {
	service interfaces.TTSServiceInterface
}

func NewTTSController(service interfaces.TTSServiceInterface) *TTSController {
	return &TTSController{
		service: service,
	}
}

func (contr *TTSController) PdfToText(c *gin.Context) {
	// call the Service
	file, error := c.FormFile("file")
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file found under key 'file'",
		})
		return
	}

	pdfRequest := models.PdfRequest{
		File:     file,
		FromPage: c.DefaultPostForm("fromPage", "1"),
		ToPage:   c.DefaultPostForm("toPage", "1"),
	}

	content, err := contr.service.PdfToText(pdfRequest)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to get endpoints. Error: " + err.Error(),
		})
		return
	}

	c.JSON(200, content)
}
