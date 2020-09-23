package twelve

import (
	"fmt"
	"strings"
)

const verse = "On the %s day of Christmas my true love gave to me: %s."

var gifts = []string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"a Partridge in a Pear Tree",
}

var days = []string{
	"first", "second", "third", "fourth",
	"fifth", "sixth", "seventh", "eighth",
	"ninth", "tenth", "eleventh", "twelfth",
}

// Song returns the lyrics to 'The Twelve Days of Christmas'.
func Song() string {
	verses := make([]string, len(days))
	for i := range days {
		verses[i] = Verse(i + 1)
	}
	return strings.Join(verses, "\n")
}

// Verse returns the nth verse to 'The Twelve Days of Christmas'.
func Verse(n int) string {
	return fmt.Sprintf(verse, days[n-1], listGifts(n))
}

func listGifts(n int) string {
	if n == 1 {
		return gifts[11]
	}
	return fmt.Sprintf(
		"%s, and %s",
		strings.Join(gifts[12-n:11], ", "),
		gifts[11],
	)
}
