# Tail call optimizations and Go

Here is just a little benchmark of how re-structuring a recursive function can
lead to better performance with tail call optimizations.

<https://www.youtube.com/watch?v=QWLyIBkBrl0> is a video comparing Rust and Go,
and the POST implementation where it used a Fibonacci recursive function to
simulate load. Two things though. First, the algorithms between the Go and
Rust versions weren't the same. And second, the recursion used by the
benchmark isn't written in a way such that the Go compiler can optimize the
recursion. Now, granted on that second point, the purpose of the recursive call
is to simulate some form of load - and optimizing would defeat the purpose.

To be clear, the resulting assessment of the video is correct - Rust has better
performance than Go, just not by the scales the author presented.

## Local results

    ↪ go test -bench=.
    goos: darwin
    goarch: arm64
    pkg: shifteleven.com/recursionbenchmark
    BenchmarkBasicRecursion-8                      4         312238250 ns/op
    BenchmarkExtraCaseRecursion-8                  7         150884179 ns/op
    BenchmarkTCORecursion-8                 33120693             35.36 ns/op
    PASS
    ok      shifteleven.com/recursionbenchmark      5.051s
