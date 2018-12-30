package main

import (
	"fmt"
	"log"

	"github.com/yanzay/tbot"
)

func main() {
	bot, err := tbot.NewServer("633097808:AAEioeQiYXDgM0GtxB_4OPXMYS9cbqSLUhQ")
	if err != nil {
		log.Fatal(err)
	}
	bot.HandleFunc("/getAllTaxis", getAllTaxisByBot)
	bot.HandleFunc("/bookAvailableTaxi", bookAnyTaxiByBot)
	bot.ListenAndServe()
}

func getAllTaxisByBot(m *tbot.Message) {
	taxis := getTaxis(baseURL)
	for _, taxi := range taxis {
		m.Replyf("\n_______________")
		m.Replyf("\nState = " + taxi.State)
		m.Replyf("Name   = " + taxi.Name)
		m.Replyf("Location   = Lon:" + fmt.Sprintf("%f", taxi.Location.Lon) + ", Lat: " + fmt.Sprintf("%f", taxi.Location.Lat))
		m.Replyf("City   = " + taxi.City)
		m.Replyf("\n_______________")
	}
}

func bookAnyTaxiByBot(m *tbot.Message) {
	taxis := getTaxis(baseURL)
	for _, taxi := range taxis {
		if taxi.State == "free" {
			m.Replyf("Trying to book taxi: " + taxi.Name)
			if bookTaxi(taxi.City, taxi.Name) == 200 {
				m.Replyf("Taxi " + taxi.Name + " booked for you!")
			} else {
				m.Replyf("Something went wrong, please try again later")
			}
			break
		}
	}
}
