package rotationalcipher

import (
	"strings"
)

// RotationalCipher encodes a given input by a set shift distance
func RotationalCipher(input string, distance int) string {
	return strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return shiftRune(r, distance, 'A')
		}
		if r >= 'a' && r <= 'z' {
			return shiftRune(r, distance, 'a')
		}
		return r
	}, input)
}

func shiftRune(r rune, distance int, a rune) rune {
	return a + (26+r-a+int32(distance))%26
}
