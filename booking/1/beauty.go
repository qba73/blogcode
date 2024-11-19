package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, err := time.Parse(layout, date)
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
	} else if time.Now().After(t) {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := t.Hour()
	if err != nil {
		panic(err)
	} else if hour >= 12 && hour < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	min, hour, day, month, year, weekday := t.Minute(), t.Hour(), t.Day(), t.Month(), t.Year(), t.Weekday()
	if err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", weekday, month, day, year, hour, min)
	}
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	thisYear := time.Now().Year()
	tString := fmt.Sprintf("%d-09-15 00:00:00 +0000 UTC", thisYear)
	formatted := "2006-01-02 15:04:05 -0700 UTC"
	t, err := time.Parse(formatted, tString)
	if err != nil {
		panic(err)
	} else {
		return t
	}
}
