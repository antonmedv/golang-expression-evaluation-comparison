# Go expression evaluation comparison

I've created an [expression evaluation package](https://github.com/antonmedv/expr) 
and wanted to compare it performance against other similar project. So I created
this repository with benchmarks for various packages.

Benchmarks are run with the following specs:

```
goos: darwin
goarch: arm64
cpu: Apple M2
```

Benchmarks:

```
Benchmark_expr-8               	342251774	        70.27 ns/op
Benchmark_celgo-8              	263402224	        91.27 ns/op
Benchmark_goja-8               	164888170	        146.2 ns/op
Benchmark_govaluate-8          	157343320	        152.9 ns/op
Benchmark_otto-8               	63185932	        380.1 ns/op
Benchmark_gval-8               	58023624	        412.9 ns/op
Benchmark_evalfilter-8         	24184514	        990.6 ns/op
Benchmark_bexpr-8              	18741454	        1276 ns/op
Benchmark_starlark-8           	 9002794	        2653 ns/op
```

And additional benchmarks for some specific cases. 

Starts with:

```
Benchmark_expr_startswith-8    	192129724	       124.9 ns/op
Benchmark_celgo_startswith-8   	178077362	       134.8 ns/op
```

Custom function call:

```
Benchmark_expr_func-8          	260001452	        92.33 ns/op
Benchmark_celgo_func-8         	217867923	       110.2 ns/op
```

Map predicate:

```
Benchmark_expr_map-8           	 5089671	      4695 ns/op
Benchmark_celgo_map-8          	  982980	     24421 ns/op
```

## Usage

You can clone repo and run benchmarks yourself.

```bash
go test -bench=. -benchtime=20s
```
