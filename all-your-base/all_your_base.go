package allyourbase

import "errors"

// ConvertToBase takes a series of digits with an input base
// and returns a seres of digits in an output base
func ConvertToBase(inBase int, inDigits []int, base int) (digits []int, err error) {
	if inBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if base < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	// find input value v
	v := 0
	for _, digit := range inDigits {
		if digit < 0 || digit >= inBase {
			return []int{0}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		v = v*inBase + digit
	}

	// return slice of digits of input value
	if v == 0 {
		return []int{0}, nil
	}

	for ; v > 0; v = v / base {
		digits = append([]int{v % base}, digits...)
	}

	return digits, nil
}
