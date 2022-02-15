package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

var requestNumber int

const port = ":8000"

func main() {
	http.HandleFunc("/test", handleTest())
	http.HandleFunc("/reset", handleReset())
	log.Printf("Starting API server on port %s", port)
	http.ListenAndServe(port, nil)
}

func handleTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestNumber++

		log.Printf("Getting request number: %v, %f",
			requestNumber,
			math.Pow(float64(requestNumber), float64(requestNumber)))
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
