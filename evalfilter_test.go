package main

import (
	"testing"

	"github.com/skx/evalfilter/v2"
)

func Benchmark_evalfilter(b *testing.B) {

	var ret bool
	var err error

	params := createParams()

	// Script we run has to be modified a little, to turn it into a
	// filter with a boolean return-value.
	src := `if ( (Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1) ) { return true; } else { return false; }`

	eval := evalfilter.New(src)

	err = eval.Prepare()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ret, err = eval.Run(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !ret {
		b.Fail()
	}
}
