package etl

import "strings"

// Transform old scrabble scores system to new system
func Transform(s map[int][]string) map[string]int {
	r := make(map[string]int, 26)
	for v, letters := range s {
		for _, l := range letters {
			r[strings.ToLower(l)] = v
		}
	}
	return r
}
