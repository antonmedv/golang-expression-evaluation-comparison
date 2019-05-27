package main

type Segment struct {
	Origin      string
	Destination string
}

type Passengers struct {
	Adults   int
	Children int
	Infants  int
}

type Price struct {
	Currency string
	Value    int
}

type Ticket struct {
	Prices map[string]*Price
}

type Env struct {
	Segments   []*Segment
	Passengers *Passengers
	Tickets    []*Ticket
	Country    string
}

const full = `(Segments[0].Origin == "MOW" || Country == "RU") && (Tickets[0].Prices["oneway"].Value >= 100 || Passengers.Adults == 1)`
const simple = `(Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)`

func create() *Env {
	return &Env{
		Segments: []*Segment{
			{
				Origin:      "MOW",
				Destination: "LED",
			},
		},
		Passengers: &Passengers{
			Adults: 1,
		},
		Tickets: []*Ticket{
			{
				Prices: map[string]*Price{
					"oneway": {
						Currency: "RUB",
						Value:    100,
					},
				},
			},
		},
		Country: "RU",
	}
}
