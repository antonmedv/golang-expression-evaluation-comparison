package main

import (
	"testing"

	"github.com/antonmedv/expr"
)

func Benchmark_expr(b *testing.B) {
	env := create()

	script, err := expr.Parse(example)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		out, err = expr.Run(script, env)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
