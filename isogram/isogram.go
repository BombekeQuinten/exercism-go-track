// Package isogram checks for isograms.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is a isogram.
func IsIsogram(input string) bool {
	present := make(map[rune]bool)

	for _, letter := range strings.ToLower(input) {
		if unicode.IsLetter(letter) {
			if present[letter] {
				return false
			}
			present[letter] = true
		}
	}
	return true
}
