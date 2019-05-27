package main

import (
	"testing"

	"github.com/antonmedv/expr/compiler"
	"github.com/antonmedv/expr/parser"
	"github.com/antonmedv/expr/vm"
)

func Benchmark_expr(b *testing.B) {
	env := create()

	tree, err := parser.Parse(example)
	if err != nil {
		b.Fatal(err)
	}

	program, err := compiler.Compile(tree)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	for n := 0; n < b.N; n++ {
		out, err = vm.Run(program, env)
	}

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
