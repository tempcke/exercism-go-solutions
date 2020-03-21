package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode Run-length encoding (RLE) is a simple form of data compression, where runs
// (consecutive data elements) are replaced by just one data value and count.
func RunLengthEncode(input string) string {
	var encoded strings.Builder
	runes := []rune(input)
	for i := 0; i < len(runes); {
		r, n := runeRunLength(runes, i)
		encoded.WriteString(encodeRune(r, n))
		i += n
	}
	return encoded.String()
}

func runeRunLength(runes []rune, i int) (rune, int) {
	r := runes[i]
	n := 1
	for j := i + 1; j < len(runes); j++ {
		if runes[j] != r {
			return r, n
		}
		n++
	}
	return r, n
}

func encodeRune(r rune, count int) string {
	if count > 1 {
		return strconv.Itoa(count) + string(r)
	}
	return string(r)
}

// RunLengthDecode decodes a run-length encoded string
func RunLengthDecode(input string) string {
	runes := []rune(input)
	var run string
	var decoded strings.Builder
	for i := 0; i < len(runes); {
		// get the string run and the new pos for i
		run, i = decodedRuneRun(runes, i)
		decoded.WriteString(run)
	}
	return decoded.String()
}

// decodeRuneRune finds the next set of repeated runes
// if the first rune is not a digit it is returned with i++
// else it collects the rune digits and repeats the first non digit run
func decodedRuneRun(runes []rune, i int) (string, int) {
	numLen := numLength(runes[i:])
	j := i + numLen
	r := runes[j]
	if numLen == 0 { // no digits, so just return the next rune
		return string(r), i + 1
	}
	count := runeDigitsToInt(runes[i:j])
	return runeRepeat(r, count), j + 1
}

// numLength returns the number of consecutive runes which are digits
func numLength(runes []rune) int {
	for i, r := range runes {
		if !unicode.IsDigit(r) {
			// i is the index of the first non-digit rune
			// therefore it is also the number of consecutive digits
			return i
		}
	}
	// for this exercise this should never happen
	// this means that all the runes where digits...
	return len(runes)
}

func runeDigitsToInt(num []rune) int {
	strnum := string(num)
	n, _ := strconv.Atoi(strnum)
	return n
}

func runeRepeat(r rune, count int) string {
	return strings.Repeat(string(r), count)
}
