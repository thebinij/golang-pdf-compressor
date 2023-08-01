// pdf-compresser.go

package pdfcompresser

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func CompressPDF(w http.ResponseWriter, r *http.Request) {

	// Assuming the PDF file is uploaded as a form file named "pdfFile".
	file, _, err := r.FormFile("pdfFile")
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a temporary file to store the uploaded PDF
	tmpFile, err := os.CreateTemp("", "input.pdf")
	if err != nil {
		http.Error(w, "Failed to create temporary file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Copy the uploaded PDF to the temporary file
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		http.Error(w, "Failed to write temporary file", http.StatusInternalServerError)
		return
	}

	// Get the absolute path of the script directory
	scriptDir := "./script" // Adjust this path if the script directory is located elsewhere
	absScriptDir, err := filepath.Abs(scriptDir)
	if err != nil {
		http.Error(w, "Failed to get script directory path", http.StatusInternalServerError)
		return
	}

	// Execute the shrinkpdf.sh script directly without using a shell
	cmd := exec.Command(filepath.Join(absScriptDir, "shrinkpdf.sh"), "-r", "90", "-o", "compressed.pdf", tmpFile.Name())

	cmd.Dir = absScriptDir // Set the script directory as the working directory for the script

	if err := cmd.Run(); err != nil {
		http.Error(w, "Failed to compress PDF", http.StatusInternalServerError)
		log.Printf("Error compressing PDF: %v", err)
		return
	}

	compressedPDFName := "compressed.pdf"

	// Serve the compressed PDF as a response
	compressedFile, err := os.Open(filepath.Join(absScriptDir, compressedPDFName))
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
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		log.Printf("Error writing response: %v", err)
		return
	}
}
