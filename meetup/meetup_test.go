package meetup

import (
	"testing"
	"time"
)

/* API
package meetup
type Weekschedule
WeekSchedule First
WeekSchedule Second
WeekSchedule Third
WeekSchedule Fourth
WeekSchedule Last
WeekSchedule Teenth
func Day(WeekSchedule, time.Weekday, time.Month, int) int
*/

var weekName = map[WeekSchedule]string{
	First:  "first",
	Second: "second",
	Third:  "third",
	Fourth: "fourth",
	Teenth: "teenth",
	Last:   "last",
}

var cases = []struct {
	nth      WeekSchedule
	weekday  time.Weekday
	month    time.Month
	year     int
	expected int
}{
	{First, time.Tuesday, time.September, 2020, 1},
	{First, time.Wednesday, time.September, 2020, 2},
	{First, time.Thursday, time.September, 2020, 3},
	{First, time.Friday, time.September, 2020, 4},
	{First, time.Saturday, time.September, 2020, 5},
	{First, time.Sunday, time.September, 2020, 6},
	{First, time.Monday, time.September, 2020, 7},
	{Second, time.Tuesday, time.September, 2020, 8},
	{Third, time.Wednesday, time.September, 2020, 16},
	{Fourth, time.Thursday, time.September, 2020, 24},
	{Last, time.Friday, time.September, 2020, 25},
	{Teenth, time.Sunday, time.September, 2020, 13},
	{Teenth, time.Saturday, time.September, 2020, 19},
}

func TestSept2020(t *testing.T) {
	for _, tc := range cases {
		actual := Day(tc.nth, tc.weekday, tc.month, tc.year)
		if actual != tc.expected {
			t.Fatalf("expected %d, got %d", tc.expected, actual)
		}
	}
}

func TestDay(t *testing.T) {
	for _, test := range testCases {
		res := Day(test.week, test.weekday, test.month, test.year)
		if res != test.expDay {
			t.Fatalf("For %s %s of %s 2013 got date of %d, want %d",
				weekName[test.week], test.weekday, test.month, res, test.expDay)
		}
	}
}
