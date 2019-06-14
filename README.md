# Go expression evaluation comparison

I've created an [expression evaluation package](https://github.com/antonmedv/expr) and wanted to compare it performance
against other similar project. So I created this this repository with benchmarks for various packages.

Here is results (lower is better):

```
Benchmark_expr-8                50000000               274 ns/op
Benchmark_celgo-8               30000000               425 ns/op
Benchmark_govaluate-8           30000000               464 ns/op
Benchmark_goja-8                30000000               472 ns/op
Benchmark_bexpr-8               20000000               998 ns/op
Benchmark_otto-8                10000000              1279 ns/op
Benchmark_gval-8                 1000000             10491 ns/op
```

## Usage

You can clone repo and run benchmarks yourself.

```bash
go test -bench=. -benchtime=20s
```
