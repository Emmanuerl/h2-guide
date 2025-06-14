package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./assets")))

	server := &http.Server{
		Addr:    ":8443",
		Handler: mux,
	}

	log.Println("Starting HTTP/2 server on https://localhost:8443")
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
