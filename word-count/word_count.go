package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is a word count map
type Frequency map[string]int

func (f Frequency) incWord(word string) {
	f[strings.Trim(word, "'")]++
}

// WordCount counts the words in the provided input
func WordCount(input string) Frequency {
	f := make(Frequency)

	for _, word := range getWords(input) {
		f.incWord(word)
	}

	return f
}

func getWords(input string) []string {
	return strings.FieldsFunc(strings.ToLower(input), isNotWordChar)
}

func isNotWordChar(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
}
