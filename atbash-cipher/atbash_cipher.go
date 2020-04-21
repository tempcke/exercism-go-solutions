package atbash

import "strings"

var encode = make(map[rune]rune, 62)

func init() {
	// encode lower
	for r := 'a'; r <= 'z'; r++ {
		encode[r] = 'z' - (r - 'a')
	}

	// encode upper to lower
	for r := 'A'; r <= 'Z'; r++ {
		encode[r] = 'z' - (r - 'A')
	}

	// include digits unencoded
	for r := '0'; r <= '9'; r++ {
		encode[r] = r
	}
}

// Atbash encode string
func Atbash(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if c, ok := encode[r]; ok {
			// every 6th char should be a space
			if (sb.Len()+1)%6 == 0 {
				sb.WriteRune(' ')
			}
			// append encoded rune
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
