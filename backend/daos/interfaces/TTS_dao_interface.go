package interfaces

import "github.com/pdftts/webapp/models"

type TTSDaoInterface interface {
	PdfToText(pdfRequest models.PdfRequest) (string, error)
}
