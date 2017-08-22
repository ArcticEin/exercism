package clock

import (
	"fmt"
)

const testVersion = 4

// Clock struct
// Attempted to make the compare test pass
type Clock int

// New function
func New(hour, minute int) Clock {
	return Clock(hour*60 + minute)
}

// String method
func (c Clock) String() string {
	// Values outside 24hours are not important
	minute := int(c) % (24 * 60)
	// Adjust one day back if minutes are negative
	if minute < 0 {
		minute += 60 * 24
	}
	hour := (minute) / 60 % 24
	minute -= hour * 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

// Add method
func (c Clock) Add(minutes int) Clock {
	return Clock(int(c) + minutes)
}
