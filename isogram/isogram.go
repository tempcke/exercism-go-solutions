package isogram

import "unicode"

// IsIsogram checks if a word or phrase has no repeating letters
func IsIsogram(input string) bool {
	var hasLetter = map[rune]bool{}
	for _, char := range input {
		if char == ' ' || char == '-' {
			continue
		}
		letter := unicode.ToLower(char)
		if hasLetter[letter] {
			return false
		}
		hasLetter[letter] = true
	}
	return true
}
