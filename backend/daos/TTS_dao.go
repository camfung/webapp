package daos

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/pdftts/webapp/models"
)

type TTSDao struct{}

func NewTTSDao() *TTSDao {
	return &TTSDao{}
}

func (dao TTSDao) PdfToText(pdfRequest models.PdfRequest) (string, error) {
	// Create a new pipe to read the file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the form
	part, err := writer.CreateFormFile("file", pdfRequest.File.Filename)
	if err != nil {
		return "", fmt.Errorf("error creating form file: %v", err)
	}

	// Open the file
	file, err := pdfRequest.File.Open()
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Copy the file content to the form
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("error copying file content: %v", err)
	}

	// Add fromPage field
	err = writer.WriteField("fromPage", pdfRequest.FromPage)
	if err != nil {
		return "", fmt.Errorf("error adding fromPage: %v", err)
	}

	// Add toPage field
	err = writer.WriteField("toPage", pdfRequest.ToPage)
	if err != nil {
		return "", fmt.Errorf("error adding toPage: %v", err)
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("error closing writer: %v", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", "http://localhost:8000/pdf-to-text", body)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set the content type
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer chickenTastesSoGood")

	// Create HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	// Check if the response status code is not 200
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned status code %d: %s", resp.StatusCode, string(responseBody))
	}

	return string(responseBody), nil
}
