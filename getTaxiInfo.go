package main

import (
	"encoding/json"
	"fmt"
)

var baseURL = "http://130.211.103.134:4000/api/v1/taxis/"

func getTaxisByCity(city string) Taxis {
	//url to get taxis by city
	url := baseURL + city

	// Returns the array of taxis
	return getTaxis(url)
}

func getTaxiByName(cityAndName string) Taxis {
	//url to get taxis by city and name
	url := baseURL + cityAndName

	// Returns the array of taxis
	return getTaxis(url)
}

func getTaxis(url string) Taxis {
	resp := doGetCall(url)
	var taxis Taxis
	// Use json.Decode for reading streams of JSON data
	// fallsback to getOneTaxi if in the body of the response there is only
	// one taxi
	if err := json.NewDecoder(resp.Body).Decode(&taxis); err != nil {
		return Taxis{getOneTaxi(url)}
	}
	// Closing of the body
	defer resp.Body.Close()

	return taxis
}

func getOneTaxi(url string) Taxi {
	resp := doGetCall(url)

	var taxi Taxi
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&taxi); err != nil {
		fmt.Println("something went wrong: ")
	}
	// Closing of the body
	defer resp.Body.Close()

	return taxi
}

func bookTaxi(city string, name string) int {
	url := baseURL + city + "/" + name
	var body = `{"state":"hired"}`
	resp := doPostCall(url, body)

	// Closing of the body
	defer resp.Body.Close()

	return resp.StatusCode
}
