package services

import (
	iDao "github.com/pdftts/webapp/daos/interfaces"
	"github.com/pdftts/webapp/models"
)

type TTSService struct {
	dao iDao.TTSDaoInterface
}

func NewTTSService(dao iDao.TTSDaoInterface) *TTSService {
	return &TTSService{dao: dao}
}

func (service TTSService) PdfToText(pdfRequest models.PdfRequest) (string, error) {
	return service.dao.PdfToText(pdfRequest)
}
