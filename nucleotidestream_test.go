package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProcessStream_InvalidStream(t *testing.T) {
	invalidStream := s{strings.NewReader("Invalid")}
	a := assert.New(t)
	processor := NewProcessor([]nucleotide{A, C, G, T})
	a.Panics(func() { processor.ProcessStream(&invalidStream, 0, 0) })
}

func ExampleProcessStream() {
	validStream := s{strings.NewReader("AGCCCCAGC")}
	processor := NewProcessor([]nucleotide{A, G, C})
	processor.ProcessStream(&validStream, 2, 0)
	// Output:
	// AGC
	// CCAGC
}

func ExampleProcessStream_ExampleSequence_ExamplePattern() {
	file, err := os.Open("sequence.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	processor := NewProcessor([]nucleotide{A, G, T, A})
	processor.ProcessStream(&s{bufio.NewReader(file)}, 5, 7)
	// Output:
	// AAGTACGTGCAG
	// CAGTGAGTAGTAGACC
	// TGAGTAGTAGACCTGA
	// ATATAAGTAGCTA
}

func ExampleProcessStream_LongSequence_LongPattern() {
	sequence := "AAGTAGTCACGTGCAGTGAGTAGTAGACCTGACGTAGACCGATATAAGTAGTCAGCTA"
	validStream := s{strings.NewReader(sequence)}
	processor := NewProcessor([]nucleotide{A, G, T, A, G, T, C, A})
	processor.ProcessStream(&validStream, 0, 2)
	// Output:
	// AGTAGTCACG
	// AGTAGTCAGC
}

func ExampleProcessStream_SuffixRollOver() {
	sequence := "TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTAGTAAGATTTT"
	validStream := s{strings.NewReader(sequence)}
	processor := NewProcessor([]nucleotide{A, G, T, A})
	processor.ProcessStream(&validStream, 2, 3)
	// Output:
	// TTAGTAAGA
}

func ExampleProcessStream_MatchRollOver() {
	sequence := "TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTAGTTT"
	validStream := s{strings.NewReader(sequence)}
	processor := NewProcessor([]nucleotide{A, G})
	processor.ProcessStream(&validStream, 2, 2)
	// Output:
	// TTAGTT
}

func ExampleProcessStream_MatchAtEndOfBuffer_RollOver() {
	sequence := "TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT" +
		"TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTAGTTT"
	validStream := s{strings.NewReader(sequence)}
	processor := NewProcessor([]nucleotide{A, G})
	processor.ProcessStream(&validStream, 2, 2)
	// Output:
	// TTAGTT
}

func BenchmarkProcessStream(b *testing.B) {
	for n := 0; n < b.N; n++ {
		file, err := os.Open("largesequence.txt")
		if err != nil {
			panic(err)
		}
		processor := NewProcessor([]nucleotide{A, G, T, A})
		processor.ProcessStream(&s{bufio.NewReader(file)}, 5, 7)
		file.Close()
	}
}
