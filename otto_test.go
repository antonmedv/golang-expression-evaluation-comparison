package main

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func Benchmark_otto(b *testing.B) {
	env := create()

	vm := otto.New()

	script, err := vm.Compile("", full)
	if err != nil {
		b.Fatal(err)
	}

	_ = vm.Set("Segments", env.Segments)
	_ = vm.Set("Passengers", env.Passengers)
	_ = vm.Set("Country", env.Country)
	_ = vm.Set("Tickets", env.Tickets)

	var out otto.Value

	for n := 0; n < b.N; n++ {
		out, err = vm.Run(script)
	}

	if err != nil {
		b.Fatal(err)
	}
	ok, err := out.ToBoolean()
	if err != nil {
		b.Fatal(err)
	}
	if !ok {
		b.Fail()
	}
}
