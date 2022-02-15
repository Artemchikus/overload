package main

import (
	"encoding/json"
	"fmt"
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
		requestNumber++
		log.Printf("Getting request number: %v", requestNumber)
		w.WriteHeader(http.StatusOK)
	}
}

func handleReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Reseting number of requests")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"number_of_requests": fmt.Sprint(requestNumber)})
		requestNumber = 0
	}
}
