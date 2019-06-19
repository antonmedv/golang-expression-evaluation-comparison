package main

import (
	"testing"

	"github.com/PaesslerAG/gval"
)

func Benchmark_gval(b *testing.B) {
	params := createParams()

	var out interface{}
	var err error

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = gval.Evaluate(example, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
