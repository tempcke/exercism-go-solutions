package anagram

import (
	"sort"
	"strings"
)

// Detect which candidates are an anagram of the provided subject
func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)

	match := make([]string, 0, len(candidates))
	sortedSubject := sortedString(subject)

	for _, c := range candidates {
		if subject == strings.ToLower(c) {
			continue
		}
		if sortedString(c) == sortedSubject {
			match = append(match, c)
		}
	}
	return match
}

func sortedString(input string) string {
	s := sortRunes([]rune(strings.ToLower(input)))
	sort.Sort(s)
	return string(s)
}

type sortRunes []rune

func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
