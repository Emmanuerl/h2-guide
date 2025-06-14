package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./assets")))

	server := &http.Server{
		Addr:    ":9443",
		Handler: mux,
	}

	log.Println("Starting HTTP/1.1 server on http://localhost:9443")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
