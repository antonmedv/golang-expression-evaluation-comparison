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

	program, err := expr.Compile(`name startsWith "/groups/" + group`, expr.Env(params))
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

func Benchmark_expr_func(b *testing.B) {
	join := expr.Function(
		"join",
		func(params ...interface{}) (interface{}, error) {
			return params[0].(string) + params[1].(string), nil
		},
	)

	program, err := expr.Compile(`join("hello", ", world")`, join)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, nil)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(string) != "hello, world" {
		b.Fail()
	}
}

func Benchmark_expr_map(b *testing.B) {
	env := map[string]interface{}{
		"array": createRange(1, 100),
	}
	program, err := expr.Compile(`map(array, # * 2)`, expr.Env(env))
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = expr.Run(program, env)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.([]interface{})[0] != 2 {
		b.Fail()
	}
}
