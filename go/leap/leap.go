package leap

import "math"

const testVersion = 3

// IsLeapYear description
func IsLeapYear(year int) bool {
	if math.Mod(float64(year), 4) == 0 {
		if math.Mod(float64(year), 100) == 0 {
			if math.Mod(float64(year), 400) == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
