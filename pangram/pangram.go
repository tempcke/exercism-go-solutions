package pangram

import "strings"

// IsPangram determines if a sentence is a pangram which uses every letter at lest once.
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for c := 'a'; c <= 'z'; c++ {
		if !strings.ContainsRune(s, c) {
			return false
		}
	}
	return true
}
