package main

import (
	"log"
	"net/http"
)

var requestNumber int

func main() {
	http.HandleFunc("/test", handleTest())
	http.HandleFunc("/reset", handleReset())
	log.Printf("Starting API server on port %s", ":8080")
	http.ListenAndServe(":8080", nil)
}

func handleTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Getting request number: %v", requestNumber)
		requestNumber++
		w.WriteHeader(http.StatusOK)
	}
}

func handleReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Reseting number of requests")
		requestNumber = 0
		w.WriteHeader(http.StatusOK)
	}
}
