package pangram

import (
	"unicode"
)

const testVersion = 2

// IsPangramMap returns a bool indicating if the input sentence contains every letter [a-z] (case insensitive)
// Benchmark: 31519 ns/op
func IsPangramMap(s string) bool {
	if len(s) < 26 {
		return false
	}
	present := make(map[rune]bool)
	for _, c := range s {
		present[unicode.ToLower(c)] = true
	}
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		if pres, _ := present[c]; !pres {
			return false
		}
	}
	return true
}

// Above solution felt slow, so I wanted to try using bitmasks for comparison
// Ended up being >20 times faster

// IsPangram returns a bool indicating if input string contains all characters [a-z]
// Benchmark 1506 ns/op
func IsPangram(s string) bool {
	if len(s) < 26 {
		return false
	}
	var bitmask int
	for _, c := range s {
		if i := uint(unicode.ToLower(c)); i >= 0 {
			bitmask = setBit(bitmask, i-97)
		}
	}
	// Compare bitmask to expected mask: b00000011 11111111 11111111 11111111 = d67108863
	if bitmask == 67108863 {
		return true
	}
	return false
}

// Code from: https://stackoverflow.com/a/23192263
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}
