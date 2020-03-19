package isbn

// IsValidISBN validates isbn10 numbers
func IsValidISBN(isbn string) bool {
	var t, m int32
	m = 10
	for _, r := range isbn {
		if v, ok := isbnCharToInt(r); ok {
			if m < 1 { // to long
				return false
			}
			if v == 10 && m != 1 { // X somewhere other than last position
				return false
			}
			t += m * v
			m--
		}
	}
	if m > 0 { // to short
		return false
	}
	return t%11 == 0
}

func isbnCharToInt(char rune) (v int32, ok bool) {
	if '0' <= char && char <= '9' {
		return char - '0', true
	}
	if char == 'X' {
		return 10, true
	}
	return 0, false
}
