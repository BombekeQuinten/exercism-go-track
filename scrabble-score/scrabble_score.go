// Package scrabble calculates letter- and wordscore.
package scrabble

import (
	"unicode"
)

// Score returns the score of a word.
func Score(word string) int {
	var score int

	for _, e := range word {
		score += LetterScore(e)
	}

	return score
}

// LetterScore returns the score of a letter.
func LetterScore(letter rune) int {
	switch unicode.ToUpper(letter) {
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 1
	}
}
