package twelve

import (
	"fmt"
	"strings"
)

const verseFormat = "On the %s day of Christmas my true love gave to me: %s."

var gifts = []string{
	"twelve Drummers Drumming, ",
	"eleven Pipers Piping, ",
	"ten Lords-a-Leaping, ",
	"nine Ladies Dancing, ",
	"eight Maids-a-Milking, ",
	"seven Swans-a-Swimming, ",
	"six Geese-a-Laying, ",
	"five Gold Rings, ",
	"four Calling Birds, ",
	"three French Hens, ",
	"two Turtle Doves, and ",
	"a Partridge in a Pear Tree",
}

var days = []string{
	"first", "second", "third", "fourth",
	"fifth", "sixth", "seventh", "eighth",
	"ninth", "tenth", "eleventh", "twelfth",
}

var verse []string

// init builds each verse as each is constant
func init() {
	verse = make([]string, len(days))
	for i := range days {
		verse[i] = fmt.Sprintf(
			verseFormat,
			days[i],
			strings.Join(gifts[11-i:], ""),
		)
	}
}

// Song returns the lyrics to 'The Twelve Days of Christmas'.
func Song() string {
	return strings.Join(verse, "\n")
}

// Verse returns the nth verse to 'The Twelve Days of Christmas'.
func Verse(n int) string {
	return verse[n-1]
}
