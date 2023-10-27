package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
)

type message struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

// Create a var of prometheus counter
var httpRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "total_http_requests",
		Help: "Number of http requests.",
	},
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	var helloMessage = message{
		ID:      1,
		Message: "Hello World",
	}

	jsonMessage, err := json.Marshal(helloMessage)
	if err != nil {
		log.Printf("Error marshalling the message: %s\n", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(jsonMessage))

	if err != nil {
		log.Printf("Error writing response to http request: %s\n", err.Error())
	}

	// Increment the counter whenever microservice1 is accessed
	httpRequests.Inc()
	return
}

// Register the httRequest counter
func init() {
	prometheus.Register(httpRequests)
}


// serverMicroservice1 writes message in response to http request
func serverMicroservice1() {
	http.HandleFunc("/", handleFunc)

	// Prometheus endpoint: http://localhost:5000/prometheus
	http.Handle("/prometheus", promhttp.Handler())

	log.Println("Starting the server at http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Printf("Error while initializing microservice1 server: %s\n", err.Error())
		return
	}
	return
}
