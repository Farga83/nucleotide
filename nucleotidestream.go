package main

import (
	"fmt"
)

// StreamProcessor defines interface for processing a stream
type StreamProcessor interface {
	ProcessStream(stream *s, x, y int)
}

// struct to implement the StreamProcessor interface
type processor struct {
	pattern []nucleotide
	borders []int
}

// NewProcessor constructs a StreamProcessor capable of processing a stream
// for a given nucleotide pattern
func NewProcessor(T []nucleotide) StreamProcessor {
	borders := preProcess(T)
	return processor{pattern: T, borders: borders}
}

// ProcessStream searches a stream for the nucleotide pattern. It utilizes the
// Knuth-Morris-Pratt algorithm for searching and leverages a buffer to handle
// printing of context information.
func (p processor) ProcessStream(stream *s, x, y int) {
	if x < 0 || y < 0 {
		return
	}
	bufferPosition, shiftDistance := 0, 0
	buffer := makeBuffer(x, y, len(p.pattern))
	size, stopCondition := loadBuffer(buffer, 0, stream)
	for bufferPosition < size {
		// Figure out how much of the prefix we are currently matching
		for shiftDistance >= 0 && buffer[bufferPosition] != p.pattern[shiftDistance] {
			shiftDistance = p.borders[shiftDistance]
		}
		shiftDistance++
		bufferPosition++
		// Perform the roll over prior to condition check in case the match
		// is at the edge of a buffer and we need to pull the trailing
		// characters
		if bufferPosition == size && !stopCondition {
			offset := x
			if shiftDistance == len(p.pattern) {
				// We are at the edge of the buffer so include the pattern in what
				// we copy over
				offset += len(p.pattern)
			} else {
				// rollover any nucleotides that are part of an existing match search
				offset += shiftDistance
			}
			bufferPosition = offset
			size, stopCondition = loadBuffer(buffer, offset, stream)
		}
		if shiftDistance == len(p.pattern) {
			printPatternAndContext(buffer, p.pattern, bufferPosition, x, y, size)
			shiftDistance = p.borders[shiftDistance]
			bufferPosition = bufferPosition - len(p.pattern) + 1
			// if j is greater than zero we are in the process of a a match
			// and should skip more known nucleotides
			if shiftDistance > 0 {
				bufferPosition += shiftDistance
			}
		}
	}
}

// There is some additional cost converting types: rune -> nucleotide -> rune
// which could be removed if the nucleotide type was left as rune
func printPatternAndContext(buffer []nucleotide, pattern []nucleotide, endingIndex, x, y, size int) {
	startingPrefix := 0
	endingPosition := size
	var prefix []nucleotide
	var suffix []nucleotide
	position := endingIndex - len(pattern) - x
	if position > 0 {
		startingPrefix = position
	}
	prefix = buffer[startingPrefix : endingIndex-len(pattern)]
	if endingIndex+y < endingPosition {
		endingPosition = endingIndex + y
	}
	suffix = buffer[endingIndex:endingPosition]
	totalOut := make([]rune, (len(prefix) + len(pattern) + len(suffix)))
	marker := 0
	for _, val := range prefix {
		totalOut[marker] = rune(val)
		marker++
	}
	for _, val := range pattern {
		totalOut[marker] = rune(val)
		marker++
	}
	for _, val := range suffix {
		totalOut[marker] = rune(val)
		marker++
	}
	fmt.Printf("%s\n", string(totalOut))
}

// knp algorithm to handle processing of the shift distance array
func preProcess(pattern []nucleotide) []int {
	i, j := 0, -1
	borders := make([]int, len(pattern)+1)
	borders[i] = j
	for i < len(pattern) {
		for j >= 0 && pattern[i] != pattern[j] {
			j = borders[j]
		}
		i++
		j++
		borders[i] = j
	}
	return borders
}

func makeBuffer(leadingCharacters, trailingCharacters, patternLength int) []nucleotide {
	// This should be further refined to consider the pre/post fix characters and their
	// respective sizes and how buffer copying will impact overall performance
	bufferSize := 4 * (patternLength + leadingCharacters + trailingCharacters)
	if bufferSize < 256 {
		bufferSize = 256
	}
	return make([]nucleotide, bufferSize)
}

func loadBuffer(buffer []nucleotide, preserve int, stream *s) (int, bool) {
	var n nucleotide
	// Copy the end elements to preserve for context printing
	for i := preserve; i > 0; i-- {
		buffer[preserve-i] = buffer[len(buffer)-i]
	}
	size := preserve
	for i := preserve; i < len(buffer); i++ {
		n = next(stream)
		if n == ε {
			return size, true
		}
		buffer[i] = n
		size++
	}
	return size, false
}
