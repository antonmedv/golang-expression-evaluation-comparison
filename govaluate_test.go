package main

import (
	"testing"

	"github.com/Knetic/govaluate"
)

func Benchmark_govaluate(b *testing.B) {
	params := createParams()

	expression, err := govaluate.NewEvaluableExpression(example)

	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		out, err = expression.Evaluate(params)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
