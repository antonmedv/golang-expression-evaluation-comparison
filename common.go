package main

type Env struct {
	Origin  string
	Country string
	Value   int64
	Adults  int64
}

const example = `(Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)`

func create() *Env {
	return &Env{
		Origin:  "MOW",
		Country: "RU",
		Value:   100,
		Adults:  1,
	}
}
