package main

import (
	"net/http"
	"log"
)

var requestNumber int

func main() {
	http.HandleFunc("/test", handleTest())
	log.Printf("Starting API server on port %s", ":8080")
	http.ListenAndServe(":8080", nil)
}

func handleTest() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Getting request number: %v", requestNumber)
		requestNumber++
		w.WriteHeader(http.StatusOK)
	}
}
