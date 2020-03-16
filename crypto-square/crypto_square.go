package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	s := normalize(pt)
	if len(s) <= 1 { return s }
	rowCount := rowCount(s)
	rows := make([]string, rowCount)
	for i, r := range s {
		row := i
		for row >= rowCount {
			row -= rowCount
		}
		rows[row] += string(r)
	}

	remainder := len(s) % rowCount
	if remainder > 0 {
		pad := rowCount - remainder
		for i := 0; i < pad; i++ {
			rows[rowCount-1-i] += " "
		}
	}
	return strings.Join(rows, " ")
}

func normalize(input string) (s string) {
	for _, r := range input {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}
		s += string(unicode.ToLower(r))
	}
	return s
}

func rowCount(s string) int {
	return int(math.Ceil(math.Sqrt(float64(len(s)))))
}