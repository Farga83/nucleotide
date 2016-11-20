package main

import (
	"fmt"
)

// nucleotide is a synonyn for supporting the requested interface
type nucleotide rune

const (
	ε nucleotide = 'ε'
	// A Represents the A tide
	A nucleotide = 'A'
	// C Represents the A tide
	C nucleotide = 'C'
	// G Represents the A tide
	G nucleotide = 'G'
	// T Represents the A tide
	T nucleotide = 'T'
)

// New creates a new nucleotide with one of the defined values or an error otherwise
func New(char rune) (nucleotide, error) {
	switch char {
	case 'A':
		return A, nil
	case 'C':
		return C, nil
	case 'G':
		return G, nil
	case 'T':
		return T, nil
	default:
		return 0, fmt.Errorf("Invalid nucleotide value: %U", char)
	}
}
