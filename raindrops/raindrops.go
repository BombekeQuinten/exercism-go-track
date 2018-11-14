// Package raindrops converts a number to "raindrop-speak".
package raindrops

import (
	"strconv"
)

// Convert converts a number to a string, the contents of which depend on the number's factors.
func Convert(num int) string {
	key := [3]int{3, 5, 7}
	val := [3]string{"Pling", "Plang", "Plong"}
	var out string

	for i, e := range key {
		if num%e == 0 {
			out += val[i]
		}
	}

	if out == "" {
		out += strconv.Itoa(num)
	}

	return out
}
