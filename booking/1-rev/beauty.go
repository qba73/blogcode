package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	min, hour, day, month, year, weekday := t.Minute(), t.Hour(), t.Day(), t.Month(), t.Year(), t.Weekday()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", weekday, month, day, year, hour, min)

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	tString := fmt.Sprintf("%d-09-15 00:00:00 +0000 UTC", time.Now().Year())
	t, err := time.Parse("2006-01-02 15:04:05 -0700 UTC", tString)
	if err != nil {
		panic(err)
	}
	return t
}
