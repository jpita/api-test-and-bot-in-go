package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxisNotEmpty(t *testing.T) {
	taxis := getTaxis(baseURL)
	assert.NotEmpty(t, taxis)
}

func TestTaxisByCityNotEmpty(t *testing.T) {
	taxis := getTaxis(baseURL)
	city := taxis[0].City
	taxis = getTaxisByCity(city)
	assert.NotEmpty(t, taxis)
}

func TestTaxisByNameNotEmpty(t *testing.T) {
	taxis := getTaxis(baseURL)
	city := taxis[0].City
	name := taxis[0].Name
	taxi := getTaxiByName(city + "/" + name)
	assert.NotEmpty(t, taxi)
}

func TestTaxiBooking(t *testing.T) {
	taxis := getTaxis(baseURL)
	for _, taxi := range taxis {
		if taxi.State == "free" {
			fmt.Println("Name   = ", taxi.Name)
			assert.Equal(t, 200, bookTaxi(taxi.City, taxi.Name))
			return
		}
	}
	assert.FailNow(t, "something went wrong with the booking of the taxi")
}
