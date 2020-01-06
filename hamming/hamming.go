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
	if len(a) != len(b) {
		return -1, ErrLengthMismatch
	}

	d := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}
