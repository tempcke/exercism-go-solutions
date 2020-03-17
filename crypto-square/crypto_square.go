package cryptosquare

import (
	"math"
	"unicode"
)

// Encode converts a given string into square code
func Encode(s string) string {
	return new(cryptoSquare).encode(s)
}

type cryptoSquare struct {
	runes     []rune
	runeCount int
	rowCount  int
	codeLen   int
}

func (c *cryptoSquare) encode(s string) string {
	c.runes = c.normalize(s)
	c.runeCount = len(c.runes)

	if c.runeCount <= 1 {
		return string(c.runes)
	}

	c.solveDementions()

	encoded := make([]rune, c.codeLen)
	for i := 0; i < c.codeLen; i++ {
		encoded[i] = c.runeForPos(i)
	}

	return string(encoded)
}

func (c *cryptoSquare) normalize(input string) []rune {
	runes := make([]rune, len(input))
	var i int
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			runes[i] = unicode.ToLower(r)
			i++
		}
	}
	return runes[0:i]
}

func (c *cryptoSquare) solveDementions() {
	l := float64(c.runeCount)
	height := math.Ceil(math.Sqrt(l))
	width := math.Ceil(l / height)
	c.rowCount = int(height)
	c.codeLen = int(((width + 1) * height) - 1)
}

func (c *cryptoSquare) runeForPos(i int) rune {
	pos := i * c.rowCount
	for pos > c.codeLen {
		pos -= c.codeLen
	}
	if pos >= c.runeCount {
		return ' '
	}
	return c.runes[pos]
}
