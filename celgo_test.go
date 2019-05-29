package main

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types/ref"
)

func Benchmark_celgo(b *testing.B) {
	params := createParams()

	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewIdent("Origin", decls.String, nil),
			decls.NewIdent("Country", decls.String, nil),
			decls.NewIdent("Value", decls.Int, nil),
			decls.NewIdent("Adults", decls.Int, nil),
		),
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

	for n := 0; n < b.N; n++ {
		out, _, err = prg.Eval(params)
	}

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
		cel.Declarations(
			decls.NewIdent("name", decls.String, nil),
			decls.NewIdent("group", decls.String, nil),
		),
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

	for n := 0; n < b.N; n++ {
		out, _, err = prg.Eval(params)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.Value().(bool) {
		b.Fail()
	}
}
