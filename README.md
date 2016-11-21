# Nucldeotide solution notes

There is a testing dependency on testify. If you don't have it installed run

```go get github.com/stretchr/testify```

The solution leverages a stream function that matches the signature

```nucleotide next(stream *s);```

The core logic is based on the Knuth-Morris-Pratt algorithm. To handle context printing 
a buffer was implemented to process the stream in chunks.

Some validation was skipped based on "Be aware" section of the problem as I took those as guarantees, 
but I would add checks for each if this was to see production (e.g. if T was empty).

A benchmark test was included that has 1638400 nucleotides defined in largesequence.txt.
To execute the benchmark run `go test --run=Benchmark* -bench=.`

# Known inefficiencies

The way I restricted the nucleotide values causes some conversion back to rune for printing. This could be better handled. 

The buffer size and rollover could be optimized, but I opted to leave them as is for now.

Stream could be a interface to make it easier to call the process function

I make the processor stateful rather a func as I felt the border array was sufficient setup to
warrant it. That being said, I could still add a helper func, but am not sure what the signature should be.
Perhaps it would be something like the following

``` func processStream(streamer func(stream *s) nucleotide, T []nucleotide, x, y int)) ```
