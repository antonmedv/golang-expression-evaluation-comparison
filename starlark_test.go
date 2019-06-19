package main

import (
	"testing"

	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

func Benchmark_starlark(b *testing.B) {
	thread := &starlark.Thread{Name: "example"}
	predeclared := starlark.StringDict{
		"greeting": starlark.String("hello"),
		"Origin":   starlark.String("MOW"),
		"Country":  starlark.String("RU"),
		"Adults":   starlark.MakeInt(1),
		"Value":    starlark.MakeInt(100),
	}

	expr, err := syntax.ParseExpr("example.star", `(Origin == "MOW" or Country == "RU") and (Value >= 100 or Adults == 1)`, syntax.RetainComments)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = starlark.EvalExpr(thread, expr, predeclared)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(starlark.Bool).Truth() {
		b.Fail()
	}
}
