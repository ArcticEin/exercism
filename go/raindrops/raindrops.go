package raindrops

import (
	"fmt"
)

const testVersion = 3

// Convert converts a number into a string based on the factors present in it
func Convert(num int) string {
	var ppp string
	factors := []int{3, 5, 7}
	raindrop := []string{"Pling", "Plang", "Plong"}
	for i := 0; i < len(factors); i++ {
		if num%factors[i] == 0 {
			ppp += raindrop[i]
		}
	}
	if len(ppp) == 0 {
		return fmt.Sprintf("%d", num)
	}
	return ppp
}
