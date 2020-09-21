package rotationalcipher

import "unicode"

// RotationalCipher encodes a given input by a set shift distance
func RotationalCipher(input string, shift int) string {
	encoded := make([]rune, len(input))
	for i, r := range input {
		encoded[i] = runeShift(r, shift)
	}
	return string(encoded)
}

func runeShift(r rune, distance int) rune {
	if !unicode.IsLetter(r) {
		return r
	}
	a := int(aOrA(r)) // 97 is the ascii value of the letter a
	return rune(a + (26+int(r)-a+distance)%26)
}

func aOrA(r rune) rune {
	if r <= 'Z' {
		return 'A'
	}
	return 'a'
}
