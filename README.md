# Test for QA position

## Considerations
This is my first time using Golang. 

My understanding of the language and its building tools is very limited so when you try running the code on your machine you might encounter some issues. 

I chose Golang because I've been wanting to try it for a long time and its popularity is on the rise.

Feedback is welcomed if you encounter any problems.

## Requirements
* [dep](https://github.com/golang/dep)

The vendor directory is gitignored, to ensure that the dependencies exist in the local
machine please run dep before proceeding with any of the below.

```console
dep ensure
```

## API Tests

In order to run the tests you need to execute the following command inside the root folder 

`go test -v -race`

### Test explanation

#### `TestTaxisNotEmpty`

This test checks that the call to `http://130.211.103.134:4000/api/v1/taxis` doesn't return an empty json.

#### `TestTaxisByCityNotEmpty`

This test checks that the call to `http://130.211.103.134:4000/api/v1/taxis/[CITY]` doesn't return an empty json.

It gets the first city from the `/taxis/` endpoint


#### `TestTaxisByNameNotEmpty`

This test checks that the call to `http://130.211.103.134:4000/api/v1/taxis/[CITY]/[NAME]` doesn't return an empty json.

It gets the first city and taxi name from the `/taxis/` endpoint

#### `TestTaxiBooking`

This test tries to book the first free taxi in the array from `/taxis/`.

## Telegram bot

I built a telegram bot that allows us to get the free taxis and book a taxi.

In order to put the bot running, you need to run the following commands from the root folder:

`go build .`

and 

`./api-test-and-bot-in-go` 

Then you need to open telegram on your phone or browser and search for the bot named `bookTaxiPita_bot`.

You can use the `/help` command to get the available commands.

Right now the implemented commands are:

* `/getAllTaxis` - prints all the taxis returned by the API with their data (State, Name, Location, City)
* `/bookAvailableTaxi` - tries to book the first available taxi (with State = Free)


![telegram bot](https://i.imgur.com/kvejwU0.jpg)

### Improvements that can be made
* book a taxi by name or city or location
* get available taxis by location or city 
* error handling: most of the cases I'm just assuming the type of variable returned or that everything goes well, I should protect the code for when something goes wrong
