// Package grains calculate the number of grains of wheat on a chessboard given that the number on each square doubles.
package grains

import (
	"fmt"
)

// Square calculates how many grains were on each square.
func Square(square int) (uint64, error) {
	if 1 > square || square > 64 {
		return 0, fmt.Errorf("square must be between 1 and 64")
	}

	return 1 << uint64(square-1), nil
}

// Total calculates the total number of grains on a chessboard.
func Total() uint64 {
	return ^uint64(0)
}
