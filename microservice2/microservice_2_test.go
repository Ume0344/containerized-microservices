package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReversedString(t *testing.T) {
	reversedMessage := getReversedString("Hello World")

	assert.Equal(t, reversedMessage, "dlroW olleH")
}
func TestGetMessageSuccessful(t *testing.T) {
	url := "http://localhost:5000"
	message := getMessage(url)

	assert.Equal(t, message.Message, "Hello World")
	assert.Equal(t, message.ID, 1)
}

func TestGetMessageFailed(t *testing.T) {
	url := "http://localhost:5001"
	message := getMessage(url)

	assert.Equal(t, message.Message, "Could not get message from microservice 1 server. Please check if url is correct and server is up.")
	assert.Equal(t, message.ID, -1)
}
