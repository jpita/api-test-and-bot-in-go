package main

import (
	"log"
	"net/http"
	"strings"
)

var client = &http.Client{}

// Helper function to do POST calls
func doPostCall(url, body string) *http.Response {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	return resp
}

// Helper function to do GET calls
func doGetCall(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	return resp
}
