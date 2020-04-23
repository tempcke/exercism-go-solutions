package armstrong

// IsNumber is used to to determine whether a number is an Armstrong number.
// An Armstrong number is a number that is the sum of its own digits each raised to the power of the number of digits.
func IsNumber(n int) bool {
	var sum int
	ds := digits(n)
	l := len(ds)
	for _, d := range ds {
		v := d
		for i := 0; i < l-1; i++ {
			v *= d
		}
		sum += v
	}
	return sum == n
}

func digits(n int) []int {
	ints := make([]int, 0, 10)
	for n > 0 {
		d := n % 10
		ints = append(ints, d)
		n /= 10
	}
	return ints
}
