# Go Compress/Decompress algorithm benchmark results

Testing Golang's built in compression Zlib, Flate, Gzip and LZW CPU/Memory performance.  
System AMD Ryzen 3600 with 32GB of DDR4 RAM.

```
$ ~/g/s/g/c/go-compress-bench> go test -bench=. -benchmem -cpuprofile profile.out
pi size 100002
goos: darwin
goarch: amd64
BenchmarkZlibCompress-12       	1000000000	         0.394 ns/op	       0 B/op	       0 allocs/op
BenchmarkZlibDecompress-12     	1000000000	         0.000424 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlateCompress-12      	1000000000	         0.388 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlateDecompress-12    	1000000000	         0.000298 ns/op	       0 B/op	       0 allocs/op
BenchmarkGzipCompress-12       	1000000000	         0.391 ns/op	       0 B/op	       0 allocs/op
BenchmarkGzipDecompress-12     	1000000000	         0.000487 ns/op	       0 B/op	       0 allocs/op
BenchmarkLZWCompress-12        	1000000000	         0.141 ns/op	       0 B/op	       0 allocs/op
BenchmarkLZWDecompress-12      	1000000000	         0.000215 ns/op	       0 B/op	       0 allocs/op
PASS
tested 2000 times each
 zlib avg size 46967 46pc
flate avg size 46961 46pc
 gzip avg size 46979 46pc
  lzw avg size 50548 50pc
ok  	_/~/go/src/github.com/comomac/go-compress-bench     25.196s
$ ~/g/s/g/c/go-compress-bench>
```

## CPU profile
![cpu profile](https://raw.githubusercontent.com/comomac/go-compress-bench/master/profile001-cpu.gif)

## Memory profile
![memory profile](https://raw.githubusercontent.com/comomac/go-compress-bench/master/profile002-mem.gif)

## Block profile
![block profile](https://raw.githubusercontent.com/comomac/go-compress-bench/master/profile003-block.gif)
