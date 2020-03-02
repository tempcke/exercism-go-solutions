package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides.
type DNA []byte

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range d {
		if _, ok := h[nucleotide]; !ok {
			return nil, fmt.Errorf("invalid nucleotide %v", nucleotide)
		}
		h[nucleotide]++
	}
	return h, nil
}
