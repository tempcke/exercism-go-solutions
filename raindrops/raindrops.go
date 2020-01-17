// Package raindrops is used to convert integers into rain drop sounds
package raindrops

import (
	"sort"
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

	for _, n := range orderedKeys() {
		if num%n == 0 {
			r += sounds[n]
		}
	}

	if r == "" {
		r = strconv.Itoa(num)
	}

	return r
}

func orderedKeys() []int {
	var keys []int

	for i := range sounds {
		keys = append(keys, i)
	}

	sort.Ints(keys)
	return keys
}
