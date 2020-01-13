// Package hamming is used to calculate the Hamming Distance between two DNA strands.
package hamming

//InputError strings must be the same length
type InputError string

func (e InputError) Error() string {
	return string(e)
}

//ErrLengthMismatch strings must be the same length
const ErrLengthMismatch = InputError("strings must be the same length")

//Distance Find the number of differences between 2 equal length strings
func Distance(a, b string) (int, error) {
	x, y, d := []rune(a), []rune(b), 0

	if len(x) != len(y) {
		return 0, ErrLengthMismatch
	}

	for i := range x {
		if x[i] != y[i] {
			d++
		}
	}

	return d, nil
}
