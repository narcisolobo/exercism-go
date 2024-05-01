package gigasecond

import "time"

/*
AddGigasecond determines the date and time one gigasecond
after a certain date.
A gigasecond is one thousand million seconds.
*/
func AddGigasecond(t time.Time) time.Time {
	gs := time.Second * 1000000000
	return t.Add(gs)
}
