package prime

// Factors computes the prime factors of a given natural number
func Factors(num int64) []int64 {
	factors := make([]int64, 0, 5)

	for i := int64(2); i <= num; i++ {
		for num%i == 0 {
			factors = append(factors, i)
			num /= i
		}
	}

	return factors
}
