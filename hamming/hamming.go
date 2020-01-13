// Package hamming is used to calculate the Hamming Distance between two DNA strands.
package hamming

import "errors"

const errLengthMismatch = "strings must be the same length"

//Distance Find the number of differences between 2 equal length strings
func Distance(a, b string) (int, error) {
	x, y := []rune(a), []rune(b)

	if len(x) != len(y) {
		return 0, errors.New(errLengthMismatch)
	}

	var d int

	for i := range x {
		if x[i] != y[i] {
			d++
		}
	}

	return d, nil
}
