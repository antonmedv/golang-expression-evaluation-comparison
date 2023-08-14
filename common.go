package main

const example = `(Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)`

func createParams() map[string]interface{} {
	params := make(map[string]interface{})
	params["Origin"] = "MOW"
	params["Country"] = "RU"
	params["Adults"] = 1
	params["Value"] = 100
	return params
}

type Params struct {
	Origin  string
	Country string
	Value   int
	Adults  int
}

func createRange(from, to int) []int {
	array := make([]int, to-from+1)
	for i := range array {
		array[i] = i + from
	}
	return array
}
