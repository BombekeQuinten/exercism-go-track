// Package diffsquares calculates the difference between the square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

import "math"

// SquareOfSum returns the square of the sum of the first N natural numbers.
func SquareOfSum(number int) int {
	var sum int

	for i := 1; i <= number; i++ {
		sum += i
	}

	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares returns the sum of the squares of the first N natural numbers.
func SumOfSquares(number int) int {
	var sum int

	for i := 1; i <= number; i++ {
		sum += int(math.Pow(float64(i), 2))
	}

	return sum
}

// Difference returns the difference between the SquareOfSum and SumOfSquares.
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
