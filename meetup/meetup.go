package meetup

import "time"

// WeekSchedule describes the week of the month that the event will happen
type WeekSchedule int

// nth WeekSchedule
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns the day of the month an event will happen
func Day(nth WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	t := first(weekday, month, year)

	switch {
	case nth == First:
		return t.Day()
	case nth == Second || (nth == Teenth && t.Day() > 5):
		return t.Day() + 7
	case nth == Third || nth == Teenth:
		return t.Day() + 14
	case nth == Fourth:
		return t.Day() + 21
	case nth == Last:
		t = t.AddDate(0, 0, 14)
		for {
			nextweek := t.AddDate(0, 0, 7)
			if nextweek.Month() != t.Month() {
				return t.Day()
			}
			t = nextweek
		}
	}

	return 0
}

func first(weekday time.Weekday, month time.Month, year int) time.Time {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	offset := int(weekday - t.Weekday())

	if weekday >= t.Weekday() {
		return t.AddDate(0, 0, offset)
	}
	return t.AddDate(0, 0, offset+7)
}
