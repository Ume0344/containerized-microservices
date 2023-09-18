package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleFunc(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(handleFunc))

	resp, _ := http.Get(testServer.URL)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but getting %d instead\n", http.StatusOK, resp.StatusCode)
	}

	var msg message

	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &msg)

	assert.Equal(t, msg.Message, "Hello World")
	assert.Equal(t, msg.ID, 1)
}
