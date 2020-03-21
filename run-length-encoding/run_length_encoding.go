package encode

import (
	"strconv"
	"unicode"
)

// RunLengthEncode Run-length encoding (RLE) is a simple form of data compression, where runs
// (consecutive data elements) are replaced by just one data value and count.
func RunLengthEncode(input string) string {
	var encoded string
	runes := []rune(input)
	for i := 0; i < len(runes); {
		r, n := runeRunLength(runes, i)
		encoded += encodeRune(r, n)
		i += n
	}
	return string(encoded)
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
	var decoded, run string
	for i := 0; i < len(runes); {
		run, i = decodedRuneRun(runes, i)
		decoded += run
	}
	return decoded
}

func decodedRuneRun(runes []rune, i int) (string, int) {
	if unicode.IsDigit(runes[i]) {
		digits := make([]rune, 0, 3)
		for j := i; j < len(runes); j++ {
			r := runes[j]
			if !unicode.IsDigit(r) {
				count, _ := strconv.Atoi(string(digits))
				return runeRun(r, count), j + 1
			}
			digits = append(digits, r)
		}
	}
	return string(runes[i : i+1]), i + 1
}

func runeRun(r rune, count int) string {
	runes := make([]rune, count)
	for i := 0; i < count; i++ {
		runes[i] = r
	}
	return string(runes)
}
