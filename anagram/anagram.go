package anagram

import (
	"strings"
	"unicode"
)

// Detect which candidates are an anagram of the provided subject
func Detect(s string, candidates []string) []string {
	s = strings.ToLower(s)
	sMap := letterMap(s)
	matching := make([]string, len(candidates))
	var i int

	for _, candidate := range candidates {
		c := strings.ToLower(candidate)
		if !isPossibleMatch(s, c) {
			continue
		}
		if !isExactMatch(letterMap(c), sMap) {
			continue
		}
		matching[i] = candidate
		i++
	}

	return matching[0:i]
}

func isPossibleMatch(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == b {
		return false
	}
	return true
}

func letterMap(input string) []int {
	chars := make([]int, 26)
	for _, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}
		c = unicode.ToLower(c)
		chars[c-'a']++
	}
	return chars
}

func isExactMatch(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
