# Go expression evaluation comparison

I've created an [expression evaluation package](https://github.com/antonmedv/expr) and wanted to compare it performance
against other similar project. So I created this this repository with benchmarks for various packages.

Here is results:

```
Benchmark_expr-8        	100000000	       259 ns/op
Benchmark_celgo-8       	100000000	       344 ns/op
Benchmark_govaluate-8   	100000000	       376 ns/op
Benchmark_goja-8        	100000000	       388 ns/op
Benchmark_bexpr-8       	30000000	       776 ns/op
Benchmark_otto-8        	30000000	      1031 ns/op
Benchmark_gval-8        	 3000000	      8011 ns/op
```

## Usage

You can clone repo and run benchmarks yourself.

```bash
go test -bench=. -benchtime=20s
```
