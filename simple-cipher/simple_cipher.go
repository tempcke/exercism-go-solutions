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
			runes[i] = r + c.offsetMod(r)
			i++
		}
	}
	return string(runes[0:i])
}

func (c caesarCipher) Decode(s string) string {
	c.offset *= -1
	decoded := c.Encode(s)
	c.offset *= -1
	return decoded
}

func (c caesarCipher) offsetMod(r rune) int32 {
	o := int32(c.offset)
	if r+o > 'z' { // encode alphabet wrap
		return -1*('z'-'a'-o) - 1
	}
	if r+o < 'a' { // decode alphabet wrap
		return 'z' - 'a' + o + 1
	}
	return o
}

// vigenereCipher is a cipher style
type vigenereCipher struct {
	key string
}

// NewVigenere defines a more complex cipher using a string as key value: a Vigenère cipher
func NewVigenere(key string) Cipher {
	return vigenereCipher{key: key}
}

func (c vigenereCipher) Encode(s string) string {
	return ""
}
func (c vigenereCipher) Decode(s string) string {
	return ""
}
