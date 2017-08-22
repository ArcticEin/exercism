package hamming

import (
	"fmt"
)

const testVersion = 6

// Distance calculates the Hamming difference between two DNA strands
// Returns an error if they are not of equal length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Strands not of equal length: [%d != %d]", len(a), len(b))
	}
	var diff int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
