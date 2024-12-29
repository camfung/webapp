package interfaces

import "github.com/pdftts/webapp/models"

type TTSServiceInterface interface {
	PdfToText(models.PdfRequest) (string, error)
}
