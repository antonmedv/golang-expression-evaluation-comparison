# Go expression evaluation comparison

I've created an [expression evaluation package](https://github.com/antonmedv/expr) and wanted to compare it performance
against other similar project. So I created this this repository with benchmarks for various packages.

Here is results (lower is better):

```
Benchmark_expr-8               	10000000	       244 ns/op
Benchmark_celgo-8              	 5000000	       422 ns/op
Benchmark_govaluate-8          	 3000000	       423 ns/op
Benchmark_goja-8               	 3000000	       464 ns/op
Benchmark_bexpr-8              	 2000000	       802 ns/op
Benchmark_otto-8               	 1000000	      1336 ns/op
Benchmark_starlark-8           	  200000	      7885 ns/op
Benchmark_gval-8               	  200000	      9770 ns/op
```

## Usage

You can clone repo and run benchmarks yourself.

```bash
go test -bench=. -benchtime=20s
```
