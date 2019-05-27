package main

import (
	"testing"

	"github.com/dop251/goja"
)

func Benchmark_goja(b *testing.B) {
	env := create()

	vm := goja.New()
	program, err := goja.Compile("", full, false)
	if err != nil {
		b.Fatal(err)
	}

	vm.Set("Segments", env.Segments)
	vm.Set("Passengers", env.Passengers)
	vm.Set("Country", env.Country)
	vm.Set("Tickets", env.Tickets)

	var out goja.Value

	for n := 0; n < b.N; n++ {
		out, err = vm.RunProgram(program)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.ToBoolean() {
		b.Fail()
	}
}
