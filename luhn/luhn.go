package luhn

import "strings"

// Valid checks if the input string is valid luhn
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}

	var sum int
	var change bool = len(input)%2 == 0

	for _, r := range input {

		v := int(r - '0')

		if v < 0 || v > 9 {
			return false
		}

		if change {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}

		sum += v
		change = !change
	}

	return sum%10 == 0
}
