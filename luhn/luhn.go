package luhn

import "strings"

// Valid checks if the input string is valid luhn
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}

	var sum int32
	var change bool = len(input)%2 == 0

	for _, r := range input {
		if r < '0' || r > '9' {
			return false
		}

		// v := r - '0'

		// if change {
		// 	v *= 2
		// 	if v > 9 {
		// 		v -= 9
		// 	}
		// }

		sum += getValue(r-'0', change)
		change = !change
	}

	return sum%10 == 0
}

func getValue(v int32, change bool) int32 {
	if change {
		if v > 4 {
			return 2*v - 9
		}
		return 2 * v
	}
	return v
}
