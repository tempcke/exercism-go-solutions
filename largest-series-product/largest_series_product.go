package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct will calculate the largest product for a contiguous substring of digits
func LargestSeriesProduct(strDigits string, span int) (int64, error) {
	digits := []rune(strDigits)

	if err := validate(digits, span); err != nil {
		return 0, err
	}

	var largestProduct int64

	for i := 0; i+span <= len(digits); i++ {
		p, err := product(digits[i : i+span])
		if err != nil {
			return 0, err
		}
		if p > largestProduct {
			largestProduct = p
		}
	}

	return largestProduct, nil
}

func validate(digits []rune, span int) error {
	if len(digits) < span {
		return errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return errors.New("span must be greater than zero")
	}
	return nil
}

func product(digits []rune) (int64, error) {
	var p int64 = 1
	for _, d := range digits {
		if !unicode.IsDigit(d) {
			return 0, errors.New("digits input must only contain digits")
		}
		p *= int64(d - '0')
	}
	return p, nil
}
