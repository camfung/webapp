package models

import "mime/multipart"

type PdfRequest struct {
	File     *multipart.FileHeader `json:"file"`     // File uploaded by the user.
	FromPage string                `json:"fromPage"` // The starting page for processing.
	ToPage   string                `json:"toPage"`   // The ending page for processing.
}
