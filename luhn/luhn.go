package luhn

// Valid checks if the input string is valid luhn
func Valid(input string) bool {
	runes := []rune(input)
	var sum, count int32

	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if r == ' ' {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}

		value := r - '0'

		if count%2 != 0 {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}

		sum += value
		count++
	}

	return count > 1 && sum%10 == 0
}
