package main

import (
	"reflect"
	"testing"

	"github.com/Knetic/govaluate"
)

func Benchmark_govaluate(b *testing.B) {
	env := create()

	// govaluate doesn't support maps accessor, so we can't use `full`.
	// Let's replace it with something more simple expression to parse
	expression, err := govaluate.NewEvaluableExpression(simple)

	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		// We need to manually extract values from env,
		// govaluate can't do this for us, to we need to do it of EVERY iteration.
		// Also we going to use reflection here, as other libraries  will use it as well,
		// and we want to test all libraries without any tricks.
		params := make(map[string]interface{})
		params["Origin"] = reflect.ValueOf(env).Elem().FieldByName("Segments").Index(0).Elem().FieldByName("Origin").Interface()
		params["Country"] = reflect.ValueOf(env).Elem().FieldByName("Country").Interface()
		params["Adults"] = reflect.ValueOf(env).Elem().FieldByName("Passengers").Elem().FieldByName("Adults").Interface()
		params["Value"] = reflect.ValueOf(env).Elem().FieldByName("Tickets").Index(0).Elem().FieldByName("Prices").MapIndex(reflect.ValueOf("oneway")).Elem().FieldByName("Value").Interface()

		out, err = expression.Evaluate(params)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
