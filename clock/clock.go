package clock

import "fmt"

// Clock contains minute in a day
type Clock struct {
	min int
}

// New Clock
func New(hour, minute int) Clock {
	m := minute + hour*60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}

// String of a Clock hh:mm
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}

// Add minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.min+minutes)
}

// Subtract minutes from a Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.min-minutes)
}
