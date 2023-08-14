package main

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

func Benchmark_celgo(b *testing.B) {
	params := createParams()

	env, err := cel.NewEnv(
		cel.Variable("Origin", cel.StringType),
		cel.Variable("Country", cel.StringType),
		cel.Variable("Value", cel.IntType),
		cel.Variable("Adults", cel.IntType),
	)
	if err != nil {
		b.Fatal(err)
	}

	parsed, issues := env.Parse(example)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("parse error: %s", issues.Err())
	}
	checked, issues := env.Check(parsed)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("type-check error: %s", issues.Err())
	}
	prg, err := env.Program(checked)
	if err != nil {
		b.Fatalf("program construction error: %s", err)
	}

	var out ref.Val

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _, err = prg.Eval(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.Value().(bool) {
		b.Fail()
	}
}

func Benchmark_celgo_startswith(b *testing.B) {
	params := map[string]interface{}{
		"name":  "/groups/foo/bar",
		"group": "foo",
	}

	env, err := cel.NewEnv(
		cel.Variable("name", cel.StringType),
		cel.Variable("group", cel.StringType),
	)
	if err != nil {
		b.Fatal(err)
	}

	parsed, issues := env.Parse(`name.startsWith("/groups/" + group)`)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("parse error: %s", issues.Err())
	}
	checked, issues := env.Check(parsed)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("type-check error: %s", issues.Err())
	}
	prg, err := env.Program(checked)
	if err != nil {
		b.Fatalf("program construction error: %s", err)
	}

	var out ref.Val

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _, err = prg.Eval(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.Value().(bool) {
		b.Fail()
	}
}

func Benchmark_celgo_func(b *testing.B) {
	params := map[string]interface{}{}

	env, err := cel.NewEnv(
		cel.Function("join",
			cel.Overload("join_string_string",
				[]*cel.Type{cel.StringType, cel.StringType},
				cel.StringType,
				cel.BinaryBinding(func(lhs, rhs ref.Val) ref.Val {
					return types.String(lhs.Value().(string) + rhs.Value().(string))
				}))))
	if err != nil {
		b.Fatal(err)
	}

	parsed, issues := env.Parse(`join("hello", ", world")`)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("parse error: %s", issues.Err())
	}
	checked, issues := env.Check(parsed)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("type-check error: %s", issues.Err())
	}
	prg, err := env.Program(checked)
	if err != nil {
		b.Fatalf("program construction error: %s", err)
	}

	var out ref.Val

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _, err = prg.Eval(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.Value().(string) != "hello, world" {
		b.Fail()
	}
}

func Benchmark_celgo_map(b *testing.B) {
	params := map[string]interface{}{
		"array": createRange(1, 100),
	}

	env, err := cel.NewEnv(
		cel.Variable("array", cel.ListType(cel.IntType)),
	)
	if err != nil {
		b.Fatal(err)
	}

	parsed, issues := env.Parse(`array.map(x, x * 2)`)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("parse error: %s", issues.Err())
	}
	checked, issues := env.Check(parsed)
	if issues != nil && issues.Err() != nil {
		b.Fatalf("type-check error: %s", issues.Err())
	}
	prg, err := env.Program(checked)
	if err != nil {
		b.Fatalf("program construction error: %s", err)
	}

	var out ref.Val

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _, err = prg.Eval(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.Value().([]ref.Val)[0].Value().(int64) != 2 {
		b.Fail()
	}
}
