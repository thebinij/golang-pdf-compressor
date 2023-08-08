// handler.go

package pdfcompresser

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	scriptDir         = "./script"
	compressedPDFName = "compressed.pdf"
)

func CompressPDF(w http.ResponseWriter, r *http.Request) {
	// Log incoming request details
	log.Printf("Incoming request from %s for path %s", r.RemoteAddr, r.URL.Path)

	// Check if the HTTP method is POST
	if r.Method != http.MethodPost {
		sendJSONResponse(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	// Read the uploaded PDF file from the request
	file, header, err := r.FormFile("pdf_file")
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to read uploaded file", err)
		return
	}
	defer file.Close()

	// Ensure the uploaded file is a PDF
	if header.Header.Get("Content-Type") != "application/pdf" {
		handleError(w, http.StatusBadRequest, "Uploaded file should be PDF", nil)
		return
	}

	// Create a temporary file to store the uploaded PDF
	tmpFile, err := os.CreateTemp("", "input.pdf")
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to create temporary file", err)
		return
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Copy the uploaded file's content to the temporary file
	if _, err := io.Copy(tmpFile, file); err != nil {
		handleError(w, http.StatusBadRequest, "Failed to write temporary file", err)
		return
	}

	// Get the absolute path of the script directory
	absScriptDir, err := filepath.Abs(scriptDir)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Failed to get script directory path", err)
		return
	}

	// Get the DPI value from the request form
	dpi, err := getDPIValue(r.FormValue("dpi"))
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid DPI value", err)
		return
	}

	grayscale, err := getGrayscaleValue(r.FormValue("grayscale"))
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid grayscale value", err)
		return
	}

	if err := executeShrinkPDF(absScriptDir, dpi, grayscale, tmpFile.Name()); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to compress PDF", err)
		return
	}

	serveCompressedPDF(w, absScriptDir)
}

func serveCompressedPDF(w http.ResponseWriter, absScriptDir string) {
	// Serve the compressed PDF as a response
	compressedFilePath := filepath.Join(absScriptDir, compressedPDFName)
	compressedFile, err := os.Open(compressedFilePath)
	if err != nil {
		http.Error(w, "Failed to open compressed PDF", http.StatusInternalServerError)
		log.Printf("Error opening compressed PDF: %v", err)
		return
	}
	defer compressedFile.Close()

	// Set the appropriate headers for the response
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename="+compressedPDFName)

	// Copy the file content to the response writer
	_, err = io.Copy(w, compressedFile)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, "Failed to write responseF")
		log.Printf("Error writing response: %v", err)
		return
	}
}
