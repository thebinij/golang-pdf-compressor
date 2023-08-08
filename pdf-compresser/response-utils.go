// response-utils.go

package pdfcompresser

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func handleError(w http.ResponseWriter, status int, message string, err error) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
	sendJSONResponse(w, status, message)
}

func sendJSONResponse(w http.ResponseWriter, status int, message string) {
	response := JsonResponse{
		Status:  status,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(response)
}
