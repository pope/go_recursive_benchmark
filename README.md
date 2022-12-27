# TCO and Go

Here is just a little benchmark of how re-structuring a recursive function can
lead to better performance with tail call optimizations.

I was watching https://www.youtube.com/watch?v=QWLyIBkBrl0 - a video comparing
Rust and Go, and the POST implementation where it used a Fibonacci recursive
function to simulate load. I felt the implementation could be off where who
knows what optimizations are being done by the Rust compiler backend that we
could easily restructure in Go that would result in TCO.

To be clear, I still thing the resulting assessment of the video is correct -
Rust is more performant than Go. But it's a silly video and I wanted to do a
silly rebuttal :D

## Local Results

```
↪ go test -bench=.
goos: darwin
goarch: arm64
pkg: shifteleven.com/recursionbenchmark
BenchmarkBasicRecursion-8                      4         312238250 ns/op
BenchmarkExtraCaseRecursion-8                  7         150884179 ns/op
BenchmarkTCORecursion-8                 33120693                35.36 ns/op
PASS
ok      shifteleven.com/recursionbenchmark      5.051s
```

