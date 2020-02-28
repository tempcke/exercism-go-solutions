// Package scrabble deals with the scrabble game
package scrabble

import "unicode"

var runeValMap = map[rune]int{}

// Score returns the word score or sum of letter values in the word
func Score(word string) int {
	s := 0
	for _, letter := range word {
		s += runeValMap[letter]
	}
	return s
}

var letterValues = map[string]int{
	"aeioulnrst": 1,
	"dg":         2,
	"bcmp":       3,
	"fhvwy":      4,
	"k":          5,
	"jx":         8,
	"qz":         10,
}

//init builds the runeValMap with both lower and uppercase letters.
func init() {
	for pool, val := range letterValues {
		for _, char := range pool {
			runeValMap[char] = val
			runeValMap[unicode.ToUpper(char)] = val
		}
	}
}
