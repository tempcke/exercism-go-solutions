package isbn

// IsValidISBN validates isbn10 numbers
func IsValidISBN(isbn string) bool {
	var t, m int32
	m = 10
	for _, r := range isbn {
		switch {
		case m < 1: // to long
			return false
		case '0' <= r && r <= '9':
			t += (r - '0') * m
			m--
		case r == 'X' && m == 1:
			t += 10 * m
			m--
		case r != '-': // invalid chars
			return false
		}
	}
	return m == 0 && t%11 == 0
}
