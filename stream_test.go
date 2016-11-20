package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Next_Valid_Sequence(t *testing.T) {
	stream := s{strings.NewReader("ACGT")}
	a := assert.New(t)
	a.Equal(A, next(&stream))
	a.Equal(C, next(&stream))
	a.Equal(G, next(&stream))
	a.Equal(T, next(&stream))
	a.Equal(ε, next(&stream))
}

func Test_Next_NoElements_SpecialValue(t *testing.T) {
	stream := s{strings.NewReader("")}
	a := assert.New(t)
	a.Equal(ε, next(&stream))
}

func Test_Next_NewLine(t *testing.T) {
	stream := s{strings.NewReader("A\n")}
	a := assert.New(t)
	a.Equal(A, next(&stream))
	a.Equal(ε, next(&stream))
}

func Test_Next_Invalid_Input_SpecialChar(t *testing.T) {
	stream := s{strings.NewReader("D")}
	a := assert.New(t)
	a.Panics(func() { next(&stream) })
}
