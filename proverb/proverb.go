// Package proverb is used to generate proverbs
package proverb

import "fmt"

const nthLine = "For want of a %s the %s was lost."
const lastLine = "And all for the want of a %s."

// Proverb generates a relevant proverb given a list of inputs
func Proverb(rhyme []string) []string {
	var result []string
	for i := range rhyme {
		result = append(result, line(rhyme, i))
	}
	return result
}

func line(rhyme []string, i int) string {
	if i < len(rhyme)-1 {
		return fmt.Sprintf(nthLine, rhyme[i], rhyme[i+1])
	}
	return fmt.Sprintf(lastLine, rhyme[0])
}
