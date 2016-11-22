package main

import (
	"io"
)

const (
	newLine = '\n'
)

type s struct {
	stream io.RuneReader
}

/* I kept this signature to match the example, even though changes will need
 to be made if the hope is to test the algorithm using outside code. I think
 it might be preferrable to expose an interface in the example to make hooking
 a different stream source in more easy.

 e.g.
 type s interface {
	next() nucleotide
 }
*/
func next(stream *s) nucleotide {
	char, _, err := stream.stream.ReadRune()
	if err != nil {
		if err == io.EOF {
			return ε
		}
		// If another error occurred it is unrecoverable
		panic("Unexpected error occurred")
	}
	// Handle cases where the input could be a file. Since this is a stream
	// of data, we enforce that a newline is the end of the stream.
	if char == newLine {
		return ε
	}
	nucl, err := New(char)
	if err != nil {
		// If this isn't a valid nucleotide value there is a fundamental
		// problem and our assumptions are invalid
		panic(err.Error())
	}
	return nucl
}
