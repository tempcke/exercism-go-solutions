package summultiples

// SumMultiples sums all unique multiples of particular numbers up to but not including a provided limit
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d > 0 && i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
