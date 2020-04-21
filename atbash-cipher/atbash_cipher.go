package atbash

var cipher = make(map[rune]rune, 62)

func init() {
	// encode lower
	for r := 'a'; r <= 'z'; r++ {
		cipher[r] = 'z' - (r - 'a')
	}
	// encode upper to lower
	for r := 'A'; r <= 'Z'; r++ {
		cipher[r] = 'z' - (r - 'A')
	}
	// include digits unencoded
	for r := '0'; r <= '9'; r++ {
		cipher[r] = r
	}
}

// Atbash encode string
func Atbash(s string) string {
	result := make([]rune, 0, len(s))
	for _, r := range s {
		if c, ok := cipher[r]; ok {
			if (len(result)+1)%6 == 0 {
				result = append(result, ' ')
			}
			result = append(result, c)
		}
	}
	return string(result)
}
