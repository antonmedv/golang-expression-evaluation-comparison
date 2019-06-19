package main

import (
	"testing"

	"github.com/dop251/goja"
)

func Benchmark_goja(b *testing.B) {
	params := createParams()

	vm := goja.New()
	program, err := goja.Compile("", example, false)
	if err != nil {
		b.Fatal(err)
	}

	vm.Set("Origin", params["Origin"])
	vm.Set("Country", params["Country"])
	vm.Set("Adults", params["Adults"])
	vm.Set("Value", params["Value"])

	var out goja.Value

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.RunProgram(program)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.ToBoolean() {
		b.Fail()
	}
}
