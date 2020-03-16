package wordcount

import "unicode"

// Frequency is a word count map
type Frequency map[string]int

// WordCount counts the words in the provided input
func WordCount(input string) Frequency {
	f := make(Frequency)

	var word string

	for i, char := range input {
		if !isValidChar(input, i) {
			if len(word) > 0 {
				f[word]++
				word = ""
			}
			continue
		}
		word += string(unicode.ToLower(char))
	}

	if len(word) > 0 {
		f[word]++
	}

	return f
}

func isValidChar(input string, index int) bool {
	return isLetterOrNumber(rune(input[index])) || isApostrophe(input, index)
}

func isLetterOrNumber(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isApostrophe(input string, index int) bool {
	if input[index] != '\'' {
		return false
	}

	// ensure ' is not the first or last char in string
	if index == 0 || index == len(input)-1 {
		return false
	}

	// must have a letter before and after it
	if !isLetter(input[index-1]) || !isLetter(input[index+1]) {
		return false
	}

	return true
}

func isLetter(char byte) bool {
	return unicode.IsLetter(rune(char))
}
