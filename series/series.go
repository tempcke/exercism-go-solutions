package series

// All returns a list of all substrings of s with length n
func All(n int, s string) []string {
	substrs := make([]string, 0, len(s))
	for i := 0; i <= len(s)-n; i++ {
		substrs = append(substrs, s[i:i+n])
	}
	return substrs
}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First returns the first substring of s with length n and ok to indicate there wasn't a problem
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
