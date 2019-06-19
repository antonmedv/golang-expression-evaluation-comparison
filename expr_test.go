package main

import (
	"testing"

	"github.com/antonmedv/expr"
)

func Benchmark_expr(b *testing.B) {
	params := createParams()

	program, err := expr.Compile(example, expr.Env(params))
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

func Benchmark_expr_startswith(b *testing.B) {
	params := map[string]interface{}{
		"name":  "/groups/foo/bar",
		"group": "foo",
	}

	program, err := expr.Compile(`name startsWith "/groups/" + "foo"`, expr.Env(params))
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
