package main

import (
	"testing"

	"github.com/Knetic/govaluate"
)

func Benchmark_govaluate(b *testing.B) {
	env := create()

	// govaluate doesn't support maps accessor, so we can't use `full`.
	// Let's replace it with something more simple expression to parse
	expression, err := govaluate.NewEvaluableExpression(example)

	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		// We need to manually extract values from env,
		// govaluate can't do this for us, to we need to do it of EVERY iteration,
		// to simulate new requests with new parameters.
		params := make(map[string]interface{})
		params["Origin"] = env.Origin
		params["Country"] = env.Country
		params["Adults"] = env.Adults
		params["Value"] = env.Value

		out, err = expression.Evaluate(params)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
