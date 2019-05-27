package main

import (
	"reflect"
	"testing"

	"github.com/hashicorp/go-bexpr"
)

type bexprValue struct {
	Origin  string `bexpr:"Origin"`
	Country string `bexpr:"Country"`
	Adults  int    `bexpr:"Adults"`
	Value   int    `bexpr:"Value"`
}

func Benchmark_bexpr(b *testing.B) {
	env := create()

	// bexpr does not support operating on types with cyclical structures.
	// So, let's use something more simple. Also replace operators and parentheses as bexpr can't parse them correctly. So sad :(
	eval, err := bexpr.CreateEvaluatorForType(
		`foo.Origin == "MOW" and foo.Country == "RU" and foo.Value == 100 and foo.Adults == 1`, nil, (map[string]*bexprValue)(nil))
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		// We need to manually extract values from env,
		// bexpr can't do this for us, to we need to do it of EVERY iteration.
		// Also we going to use reflection here, as other libraries  will use it as well,
		// and we want to test all libraries without any tricks.
		params := make(map[string]*bexprValue)
		params["foo"] = &bexprValue{
			Origin:  reflect.ValueOf(env).Elem().FieldByName("Segments").Index(0).Elem().FieldByName("Origin").Interface().(string),
			Country: reflect.ValueOf(env).Elem().FieldByName("Country").Interface().(string),
			Adults:  reflect.ValueOf(env).Elem().FieldByName("Passengers").Elem().FieldByName("Adults").Interface().(int),
			Value:   reflect.ValueOf(env).Elem().FieldByName("Tickets").Index(0).Elem().FieldByName("Prices").MapIndex(reflect.ValueOf("oneway")).Elem().FieldByName("Value").Interface().(int),
		}
		out, err = eval.Evaluate(params)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
