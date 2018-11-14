// Package hamming provides a method for calculating Hamming difference.
package hamming

import (
	"fmt"
)

// Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	var counter int

	if len(a) == len(b) {
		return counter, fmt.Errorf("Strands are not of equal length")
	}

	for i := range a {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}
