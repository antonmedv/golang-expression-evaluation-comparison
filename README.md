# Go expression evaluation comparison

I've created an [expression evaluation package](https://github.com/antonmedv/expr) 
and wanted to compare it performance against other similar project. So I created
this repository with benchmarks for various packages.

Benchmarks are run with the following specs:

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
```

Benchmarks (lower is better):

```
Benchmark_expr-8               	137284434	       175.0 ns/op
Benchmark_celgo-8              	124539046	       193.3 ns/op
Benchmark_govaluate-8          	78848359	       302.6 ns/op
Benchmark_goja-8               	63050496	       387.7 ns/op
Benchmark_gval-8               	31351183	       767.4 ns/op
Benchmark_otto-8               	29288571	       825.0 ns/op
Benchmark_evalfilter-8         	10204525	      2330 ns/op
Benchmark_bexpr-8              	 8534502	      2811 ns/op
Benchmark_starlark-8           	 4184384	      5728 ns/op
```

## Usage

You can clone repo and run benchmarks yourself.

```bash
go test -bench=. -benchtime=20s
```
