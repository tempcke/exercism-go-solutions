package armstrong

// IsNumber is used to to determine whether a number is an Armstrong number.
// An Armstrong number is a number that is the sum of its own digits each raised to the power of the number of digits.
func IsNumber(n int) bool {
	var sum int
	nums := digits(n)
	exp := len(nums)
	for _, digit := range nums {
		sum += intPow(digit, exp)
	}
	return sum == n
}

// digits makes a slice of ints representing each digit in the num
// ex: digits(123) = []int{1,2,3}
func digits(n int) []int {
	ints := make([]int, 0, 10)
	for n > 0 {
		digit := n % 10
		ints = append(ints, digit)
		n /= 10
	}
	return ints
}

// intPow does n^exp faster than math.Pow
func intPow(n, exp int) int {
	val := n
	for i := 0; i < exp-1; i++ {
		val *= n
	}
	return val
}
