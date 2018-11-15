/*
Package luhn is a simple checksum formula used to validate a variety of identification numbers,
such as credit card numbers and Canadian Social Insurance Numbers.
*/
package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if a given string is valid.
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	if len(input) <= 1 || !IsNumeric(input) {
		return false
	}

	return Doubling(input)%10 == 0
}

// Doubling doubles every second digit, starting from the right and returns the sum of the string.
func Doubling(s string) int {
	var sum int
	double := len(s)%2 == 0

	for _, e := range s {
		e := int(e - '0')
		if double {
			e *= 2
			if e >= 9 {
				e -= 9
			}
		}
		sum += e
		double = !double
	}

	return sum
}

// IsNumeric checks if a string contains non-digit characters.
func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
