package gigasecond

import "time"

const testVersion = 4

// A gigasecond repsented by nanoseconds
const gigasecond = 1000000000 * time.Second

// AddGigasecond adds a gigasecond to the input time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigasecond))
}
