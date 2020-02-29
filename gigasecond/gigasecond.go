// Package gigasecond allows for gigasecond time calculations
package gigasecond

import "time"

// GigaSecond is 10^9 (1,000,000,000) seconds
const GigaSecond = 1000000000

// AddGigasecond adds a gigasecond to the provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * GigaSecond)
}
