package beer

import (
	"fmt"
	"strings"
)

// verse templates
const (
	verseZero     = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
	verseTemplate = "%[1]s of beer on the wall, %[1]s of beer.\nTake %[2]s down and pass it around, %[3]s of beer on the wall.\n"
)

// Song returns the entire song
func Song() string {
	s, _ := Verses(99, 0)
	return s
}

// Verses returns all verses within the set
func Verses(a, b int) (string, error) {
	if a > 99 || b < 0 || a < b {
		return "", fmt.Errorf("verses out of bounds: %v - %v", a, b)
	}

	var sb strings.Builder
	for i := a; i >= b; i-- {
		sb.WriteString(verseText(i) + "\n")
	}
	return sb.String(), nil
}

// Verse returns a single verse in the song
func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", fmt.Errorf("invalid verse number: %v", n)
	}
	return verseText(n), nil
}

func verseText(n int) string {
	switch n {
	case 0:
		return verseZero
	case 1:
		return fmt.Sprintf(verseTemplate, "1 bottle", "it", "no more bottles")
	case 2:
		return fmt.Sprintf(verseTemplate, "2 bottles", "one", "1 bottle")
	default:
		return fmt.Sprintf(verseTemplate, fmt.Sprintf("%v bottles", n), "one", fmt.Sprintf("%v bottles", n-1))
	}
}
