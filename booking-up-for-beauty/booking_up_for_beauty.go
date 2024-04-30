package booking

import (
	"fmt"
	"os"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// "July 25, 2019 13:45:00"
	const layout = "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// "Thursday, July 25, 2019 13:45:00"
	const layout = "Monday, January 2, 2006 15:05:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// >= 12:00 and < 18:00

	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// "You have an appointment on Thursday, July 25, 2019, at 13:45."
	// 6/6/2005 10:30:00
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "1/2/2006 15:04:05"
	const format = "Monday, January 2, 2006, at 15:04"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	f := t.Format(format)
	return fmt.Sprintf("You have an appointment on %s.", f)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
