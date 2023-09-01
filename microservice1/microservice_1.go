package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type message struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

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
	return
}

// serverMicroservice1 writes message in response to http request
func serverMicroservice1() {
	http.HandleFunc("/", handleFunc)

	log.Println("Starting the server at http://localhost:5000/")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Printf("Error while initializing microservice1 server: %s\n", err.Error())
		return
	}
	return
}
