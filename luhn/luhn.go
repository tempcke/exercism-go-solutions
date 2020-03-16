package luhn

// Valid checks if the input string is valid luhn
func Valid(input string) bool {
	sum, count := check([]rune(input))

	return count > 1 && sum%10 == 0
}

func check(runes []rune) (sum, count int32) {
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if r == ' ' {
			continue
		}
		if !isNum(r) {
			return 0, 0
		}

		sum += modVal(r-'0', count)
		count++
	}
	return sum, count
}

func isNum(r rune) bool {
	return '0' <= r && r <= '9'
}

func modVal(value, count int32) int32 {
	if count%2 != 0 {
		value *= 2
		if value > 9 {
			value -= 9
		}
	}
	return value
}
