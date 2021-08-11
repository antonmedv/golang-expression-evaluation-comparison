# Go expression evaluation comparison

I've created an [expression evaluation package](https://github.com/antonmedv/expr) and wanted to compare it performance
against other similar project. So I created this this repository with benchmarks for various packages.

Here is results (lower is better):

Basic Benchmark:
```
Benchmark_expr-16                       184196986              134.2 ns/op
Benchmark_celgo-16                      122053126              187.4 ns/op
Benchmark_govaluate-16                  77529223               295.3 ns/op
Benchmark_goja-16                       79449454               331.2 ns/op
Benchmark_gval-16                       30330537               744.4 ns/op
Benchmark_otto-16                       30256912               860.9 ns/op
Benchmark_evalfilter-16                 13407027               1780 ns/op
Benchmark_bexpr-16                      10721992               2309 ns/op
Benchmark_starlark-16                    4600719               5295 ns/op
```

StartsWith Benchmark:
```
Benchmark_expr_startswith-16            81022359               280.1 ns/op
Benchmark_celgo_startswith-16           66076286               327.7 ns/op
```

## Usage

You can clone repo and run benchmarks yourself.

```bash
go test -bench=. -benchtime=20s
```
