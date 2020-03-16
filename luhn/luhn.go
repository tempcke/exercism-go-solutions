package luhn

// Valid checks if the input string is valid luhn
func Valid(input string) bool {
	runes := []rune(input)
	length := len(runes)
	var sum, count int32

	for i := 0; i < length; i++ {
		r := runes[length-i-1]
		switch {
		default:
			return false
		case r == ' ':
			continue
		case isNum(r):
			sum += modVal(r-'0', count)
			count++
		}
	}
	return count > 1 && sum%10 == 0
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
