package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Implementation of main provided by Chris K. as he was verying the correctness
// of the solution.
func main() {
	if len(os.Args) < 4 {
		fmt.Printf("USAGE: %s x T y\n", os.Args[0])
		os.Exit(1)
	}

	x, _ := strconv.Atoi(os.Args[1])
	T := []nucleotide(os.Args[2])
	y, _ := strconv.Atoi(os.Args[3])

	stream := &s{bufio.NewReader(os.Stdin)}

	processor := NewProcessor(T)
	processor.ProcessStream(stream, x, y)
}
