package diffsquares

// SquareOfSum sums the first n natural numbers then squares it
func SquareOfSum(n int) int {
	var sum int
	for ; n > 0; n-- {
		sum += n
	}
	return sum * sum
}

// SumOfSquares sums the square of the first n natural numbers
func SumOfSquares(n int) int {
	var sum int
	for ; n > 0; n-- {
		sum += n * n
	}
	return sum
}

// Difference subtracts the SquareOfSum from the SumOfSquares to find the difference
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
