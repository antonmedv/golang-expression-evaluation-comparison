package main

import (
	"testing"

	"github.com/hashicorp/go-bexpr"
)

func Benchmark_bexpr(b *testing.B) {
	p := createParams()
	params := Params{
		Origin:  p["Origin"].(string),
		Country: p["Country"].(string),
		Value:   p["Value"].(int),
		Adults:  p["Adults"].(int),
	}

	// Replace operators and parentheses as bexpr can't parse them correctly. So sad :(
	eval, err := bexpr.CreateEvaluatorForType(
		`Origin == "MOW" and Country == "RU" and Value == 100 and Adults == 1`,
		nil,
		Params{},
	)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		out, err = eval.Evaluate(params)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
