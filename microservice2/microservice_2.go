package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type responseMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

// getReversedString reverses the message string
func getReversedString(message string) (result string) {
	for _, v := range message {
		result = string(v) + result
	}
	return result
}

// getMessage gets the message in json and convert to go struct responseMessage
func getMessage(url string) (message responseMessage) {
	var failedMessage responseMessage
	resp, _ := http.Get(url)

	if resp == nil || resp.StatusCode != 200 {
		log.Printf("Error getting response from microservice 1. resp: %v\n", resp)
		failedMessage.ID = -1
		failedMessage.Message = "Could not get message from microservice 1 server. Please check if url is correct and server is up."
		return failedMessage
	}
	content, _ := io.ReadAll(resp.Body)
	// un-marshaling the json object to struct responseMessage
	err := json.Unmarshal(content, &message)
	if err != nil {
		log.Printf("Error while unmarshalling http response: %s\n", err.Error())
	}

	resp.Body.Close()
	return message
}

// microservice2 requests microservice1 to get message and
// reverse it.
func microservice2() {
	microservice1Url := "http://microservice1-service.default.svc.cluster.local/"
	message := getMessage(microservice1Url)

	log.Printf("Received Message from microservice 1 server: %s\n", message.Message)

	if message.ID != -1 {
		reversedMessage := getReversedString(message.Message)
		log.Printf("Reversed Message : %s\n", reversedMessage)
	}
	return
}
