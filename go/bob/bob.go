package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Hey return Bob's response to Alice's question
func Hey(in string) string {
	in = strings.TrimSpace(in)
	if isSilence(in) {
		return "Fine. Be that way!"
	}
	if hasLetters(in) && isYelling(in) {
		return "Whoa, chill out!"
	}
	if isQuestion(in) {
		return "Sure."
	}
	return "Whatever."
}

func hasLetters(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func isYelling(s string) bool {
	for _, c := range s {
		if unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func isQuestion(s string) bool {
	if s[len(s)-1:] == "?" {
		return true
	}
	return false
}

func isSilence(s string) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}
