package main

import (
	"testing"

	"github.com/dop251/goja"
)

func Benchmark_goja(b *testing.B) {
	env := create()

	vm := goja.New()
	program, err := goja.Compile("", example, false)
	if err != nil {
		b.Fatal(err)
	}

	var out goja.Value

	for n := 0; n < b.N; n++ {
		// We need to set new params of every iteration,
		// to simulate new requests with new parameters.
		vm.Set("Origin", env.Origin)
		vm.Set("Country", env.Country)
		vm.Set("Adults", env.Adults)
		vm.Set("Value", env.Value)
		out, err = vm.RunProgram(program)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.ToBoolean() {
		b.Fail()
	}
}
