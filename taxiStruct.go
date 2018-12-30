package main

// Taxis creating a type Taxis
type Taxis []Taxi

// Taxi creating a type Taxi
type Taxi struct {
	State    string   `json:"state"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
	City     string   `json:"city"`
}

// Location creating a type Location
type Location struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
