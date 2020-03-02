package scale

import (
	"fmt"
	"strings"
)

var (
	chromaticScale      = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	chromaticScaleFlats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	diatonicScaleFlats  = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
	intervals           = map[rune]int{
		'm': 1,
		'M': 2,
		'A': 3,
	}
)

// Scale generates the musical scale starting with the tonic and following the specified interval pattern
func Scale(tonic string, interval string) []string {
	notes := wrapScale(determineScale(tonic), tonic)

	if len(interval) > 0 {
		return cherryPickNotes(notes, interval)
	}
	return notes
}

func determineScale(tonic string) []string {
	if contains(diatonicScaleFlats, tonic) {
		return chromaticScaleFlats
	}
	return chromaticScale
}

func wrapScale(notes []string, startNote string) []string {
	pos, _ := notePos(notes, strings.Title(startNote))
	return append(notes[pos:], notes[0:pos]...)
}

func notePos(notes []string, startNote string) (int, error) {
	for pos, note := range notes {
		if note == startNote {
			return pos, nil
		}
	}
	return 0, fmt.Errorf("no such note %s", startNote)
}

func contains(set []string, item string) bool {
	for _, a := range set {
		if a == item {
			return true
		}
	}
	return false
}

func intervalToIndexes(ms string) []int {
	i := 0
	result := make([]int, len(ms))
	for j, m := range ms {
		result[j] = i
		i += intervals[m]
	}
	return result
}

func cherryPickNotes(notes []string, interval string) []string {
	indexes := intervalToIndexes(interval)
	result := make([]string, len(indexes))
	for i, index := range indexes {
		result[i] = notes[index]
	}
	return result
}
