package main

import (
	"log"
	pdfcompresser "my-project/pdf-compresser"
	controller "my-project/ping"
	"net/http"
)

func main() {

	// Register routes
	http.HandleFunc("/ping", controller.Ping)

	http.HandleFunc("/compress-pdf", pdfcompresser.CompressPDF)

	// Start the server
	startServer()
}

func startServer() {
	// print out the server is going to start and show the time
	log.Println("Starting server....")

	// Create server at localhost:8080
	addr := ":8080"

	// Listen and serve
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
