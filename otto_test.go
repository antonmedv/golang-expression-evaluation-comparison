package main

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func Benchmark_otto(b *testing.B) {
	env := create()

	vm := otto.New()

	script, err := vm.Compile("", example)
	if err != nil {
		b.Fatal(err)
	}

	var out otto.Value

	for n := 0; n < b.N; n++ {
		// We need to set new params of every iteration,
		// to simulate new requests with new parameters.
		_ = vm.Set("Origin", env.Origin)
		_ = vm.Set("Country", env.Country)
		_ = vm.Set("Adults", env.Adults)
		_ = vm.Set("Value", env.Value)
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
