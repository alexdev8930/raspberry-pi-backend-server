package main

import (
	"log"
	"net/http"

	httptransport "rpi-backend.local/internal/transport/http"
)

func main() {
	handler := httptransport.NewHandler()
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.Routes(),
	}

	log.Println("server listening on :8080")
	log.Fatal(server.ListenAndServe())
}
