package prime

const sieveLimit = 1000

// Factors computes the prime factors of a given natural number
func Factors(num int64) []int64 {
	limit := num
	if num > sieveLimit {
		limit = sieveLimit
	}
	marked := make([]bool, limit+1)
	factors := make([]int64, 0, 10)

	for i := int64(2); i <= num; i++ {
		// if i is not prime then skip it
		if i < limit && marked[i] {
			continue
		}

		for num%i == 0 {
			factors = append(factors, i)
			num /= i
		}

		// mark multiples as not prime
		for mult := i * 2; mult <= num && mult <= limit; mult += i {
			marked[mult] = true
		}
	}

	return factors
}
