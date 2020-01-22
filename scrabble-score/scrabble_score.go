// Package scrabble deals with the scrabble game
package scrabble

import "strings"

var letterValues = map[string]int{
	"aeioulnrst": 1,
	"dg":         2,
	"bcmp":       3,
	"fhvwy":      4,
	"k":          5,
	"jx":         8,
	"qz":         10,
}

// Score returns the word score or sum of letter values in the word
func Score(word string) int {
	s := 0
	for _, letter := range strings.ToLower(word) {
		s += letterValue(string(letter))
	}
	return s
}

func letterValue(letter string) int {
	for pool, val := range letterValues {
		if strings.Contains(pool, letter) {
			return val
		}
	}
	return 0
}
