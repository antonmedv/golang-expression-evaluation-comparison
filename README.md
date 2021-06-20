# Go expression evaluation comparison

I've created an [expression evaluation package](https://github.com/antonmedv/expr) and wanted to compare it performance
against other similar project. So I created this this repository with benchmarks for various packages.

Here is results (lower is better):

```
Benchmark_expr-24                	152140300	       158.2 ns/op
Benchmark_celgo-24               	100000000	       212.1 ns/op
Benchmark_expr_startswith-24     	 81988958	       292.4 ns/op
Benchmark_govaluate-24           	 76476085	       319.9 ns/op
Benchmark_goja-24                	 73179843	       330.7 ns/op
Benchmark_celgo_startswith-24    	 71623069	       336.4 ns/op
Benchmark_otto-24                	 30886207	       770.0 ns/op
Benchmark_gval-24                	 31006279	       774.4 ns/op
Benchmark_evalfilter-24          	 14328974	        1671 ns/op
Benchmark_bexpr-24               	  9910032	        2422 ns/op
Benchmark_starlark-24            	  4702758	        5082 ns/op
```

## Usage

You can clone repo and run benchmarks yourself.

```bash
go test -bench=. -benchtime=20s
```
