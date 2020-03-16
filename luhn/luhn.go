package luhn

// Valid checks if the input string is valid luhn
func Valid(input string) bool {
	runes := []rune(input)
	length := len(runes)
	var sum int32
	var count int32

	for i := 0; i < length; i++ {
		r := runes[length-i-1]
		v := r - '0'
		switch {
		default:
			return false
		case r == ' ':
			continue
		case isNum(v):
			sum += modVal(v, count)
			count++
		}
	}
	return count > 1 && sum%10 == 0
}

func isNum(v int32) bool {
	return 0 <= v && v <= 9
}

func modVal(v, count int32) int32 {
	if count%2 != 0 {
		v *= 2
		if v > 9 {
			v -= 9
		}
	}
	return v
}
