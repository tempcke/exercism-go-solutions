// Package raindrops is used to convert integers into rain drop sounds
package raindrops

import (
	"strconv"
)

var sounds = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert number into raindrop sounds
func Convert(num int) string {
	var r string
	for n, s := range sounds {
		if num%n == 0 {
			r += s
		}
	}
	if r == "" {
		r = strconv.Itoa(num)
	}
	return r
}
