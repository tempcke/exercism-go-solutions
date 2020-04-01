package clock

import "fmt"

const (
	hrsPerDay  = 24
	minsPerHr  = 60
	minsPerDay = 1440
)

// Clock is minutes 0 to 1440
type Clock int

// New Clock
func New(hour, minute int) Clock {
	for minute < 0 {
		hour--
		minute += minsPerHr
	}

	minute += fixHour(hour) * minsPerHr

	if minute > minsPerDay {
		minute = minute % minsPerDay
	}

	return Clock(minute)
}

func fixHour(h int) int {
	h = h % hrsPerDay
	if h < 0 {
		h += hrsPerDay
	}
	return h
}

// String of a Clock hh:mm
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/minsPerHr, c%minsPerHr)
}

// Add minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

// Subtract minutes from a Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}
