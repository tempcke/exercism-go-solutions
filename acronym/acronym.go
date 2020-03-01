// Package acronym is used to create acronyms
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate creates an acronym from provided string
func Abbreviate(s string) (acronym string) {
	words := strings.Fields(stripSpecialChars(s))

	for _, w := range words {
		acronym += w[0:1]
	}

	return strings.ToUpper(acronym)
}

func stripSpecialChars(s string) string {
	r := regexp.MustCompile(`[-_]`)
	return r.ReplaceAllString(s, " ")
}
