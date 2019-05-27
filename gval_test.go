package main

import (
	"testing"

	"github.com/PaesslerAG/gval"
)

func Benchmark_gval(b *testing.B) {
	env := create()

	var out interface{}
	var err error

	for n := 0; n < b.N; n++ {
		out, err = gval.Evaluate(example, env)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
