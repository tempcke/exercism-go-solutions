package cipher

import "strings"

// caesarCipher is a cipher style
type caesarCipher struct {
	offset int
}

// NewCaesar returns a caesar style cipher with a 4 char offset
func NewCaesar() Cipher {
	return caesarCipher{offset: 3}
}

// NewShift returns a cipher with a custom offset
func NewShift(distance int) Cipher {
	if distance >= 26 || distance <= -26 || distance == 0 {
		return nil
	}
	return caesarCipher{offset: distance}
}

func (c caesarCipher) Encode(s string) string {
	runes := make([]rune, len(s))
	i := 0
	for _, r := range strings.ToLower(s) {
		if r >= 'a' && r <= 'z' {
			runes[i] = runeShift(r, c.offset)
			i++
		}
	}
	return string(runes[0:i])
}

func (c caesarCipher) Decode(s string) string {
	// a decode is just an ecode with a negative offset...
	decoder := NewShift(-1 * c.offset)
	return decoder.Encode(s)
}

// vigenereCipher is a cipher style
type vigenereCipher struct {
	key     string
	offsets []int
}

// NewVigenere defines a more complex cipher using a string as key value: a Vigenère cipher
func NewVigenere(key string) Cipher {
	if strings.Count(key, "a") == len(key) {
		return nil
	}
	offsets := make([]int, len(key))
	for i, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		offsets[i] = int(r - 'a')
	}
	return vigenereCipher{key: key, offsets: offsets}
}

func (c vigenereCipher) Encode(s string) string {
	runes := make([]rune, len(s))
	i, n := 0, 0
	for _, r := range strings.ToLower(s) {
		if r >= 'a' && r <= 'z' {
			if n == len(c.offsets) {
				n -= len(c.offsets)
			}
			offset := c.offsets[n]
			runes[i] = runeShift(r, int(offset))
			i++
			n++
		}
	}
	return string(runes[0:i])
}

func (c vigenereCipher) Decode(s string) string {
	decoder := newVigenereDecoder(c.key)
	return decoder.Encode(s)
}

func newVigenereDecoder(encoderKey string) Cipher {
	offsets := make([]int, len(encoderKey))
	for i, r := range encoderKey {
		offsets[i] = -1 * int(r-'a')
	}
	return vigenereCipher{key: encoderKey, offsets: offsets}
}

func runeShift(r rune, shift int) rune {
	a := int('a') // 97 is the ascii value of the letter a
	l := int(r)   // ascii value of rune
	return rune(a + (26+l-a+shift)%26)
}
