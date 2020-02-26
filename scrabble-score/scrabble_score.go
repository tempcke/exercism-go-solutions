// Package scrabble deals with the scrabble game
package scrabble

import (
	"unicode"
)

var runeValMap = map[rune]int{
	'a': 1,
	'b': 3,
	'c': 3,
	'd': 2,
	'e': 1,
	'f': 4,
	'g': 2,
	'h': 4,
	'i': 1,
	'j': 8,
	'k': 5,
	'l': 1,
	'm': 3,
	'n': 1,
	'o': 1,
	'p': 3,
	'q': 10,
	'r': 1,
	's': 1,
	't': 1,
	'u': 1,
	'v': 4,
	'w': 4,
	'x': 8,
	'y': 4,
	'z': 10,
}

// Score returns the word score or sum of letter values in the word
func Score(word string) int {
	s := 0
	for _, letter := range word {
		s += runeValMap[unicode.ToLower(letter)]
	}
	return s
}
