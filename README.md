# Go expression evaluation comparison

```
Benchmark_expr-8        	100000000	       259 ns/op
Benchmark_celgo-8       	100000000	       344 ns/op
Benchmark_govaluate-8   	100000000	       376 ns/op
Benchmark_goja-8        	100000000	       388 ns/op
Benchmark_bexpr-8       	30000000	       776 ns/op
Benchmark_otto-8        	30000000	      1031 ns/op
Benchmark_gval-8        	3000000	          8011 ns/op
```

## Usage

```bash
go test -bench=. -benchtime=20s
```
