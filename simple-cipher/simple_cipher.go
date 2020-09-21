package cipher

import "strings"

// caesarCipher is a cipher style
type shiftCipher struct {
	distance int
}

// NewCaesar returns a caesar style cipher with a 3 char shift
func NewCaesar() Cipher {
	return shiftCipher{distance: 3}
}

// NewShift returns a cipher with a custom shift
func NewShift(distance int) Cipher {
	if distance >= 26 || distance <= -26 || distance == 0 {
		return nil
	}
	return shiftCipher{distance: distance}
}

func (c shiftCipher) Encode(s string) string {
	runes := make([]rune, len(s))
	i := 0 // used to track only a-z chars
	for _, r := range strings.ToLower(s) {
		if r >= 'a' && r <= 'z' {
			runes[i] = runeShift(r, c.distance)
			i++
		}
	}
	return string(runes[0:i])
}

func (c shiftCipher) Decode(s string) string {
	// a decode is just an ecode with a negative shift...
	decoder := NewShift(-1 * c.distance)
	return decoder.Encode(s)
}

// vigenereCipher is a cipher style
type vigenereCipher struct {
	distance []int
}

// NewVigenere defines a more complex cipher using a string as key value: a Vigenère cipher
func NewVigenere(key string) Cipher {
	if strings.Count(key, "a") == len(key) {
		return nil
	}
	distances := make([]int, len(key))
	for i, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		distances[i] = int(r - 'a')
	}
	return vigenereCipher{distance: distances}
}

func (c vigenereCipher) Encode(s string) string {
	runes := make([]rune, len(s))
	i, dCount := 0, len(c.distance)
	for _, r := range strings.ToLower(s) {
		if r >= 'a' && r <= 'z' {
			runes[i] = runeShift(r, c.distance[i%dCount])
			i++
		}
	}
	return string(runes[0:i])
}

func (c vigenereCipher) Decode(s string) string {
	distances := make([]int, len(c.distance))
	for i, v := range c.distance {
		distances[i] = -1 * v
	}
	decoder := vigenereCipher{distance: distances}
	return decoder.Encode(s)
}

func runeShift(r rune, shift int) rune {
	a := int('a') // 97 is the ascii value of the letter a
	l := int(r)   // ascii value of rune
	return rune(a + (26+l-a+shift)%26)
}
