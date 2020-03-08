package diffsquares

// SquareOfSum sums the first n natural numbers then squares it
func SquareOfSum(n int) int {
	sum := ((n + 1) * n) / 2
	return sum * sum
}

// SumOfSquares sums the square of the first n natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference subtracts the SquareOfSum from the SumOfSquares to find the difference
func Difference(n int) int {
	return (n * (3*n + 2) * (n - 1) * (n + 1)) / 12
}
