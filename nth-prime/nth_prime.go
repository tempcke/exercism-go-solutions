package prime

// Nth prime number
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	var prime int
	for i, j := 2, 0; j < n; i++ {
		if isPrime(i) {
			prime = i
			j++
		}
	}
	return prime, true
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n%2 == 0 {
		return n == 2
	}
	if n%3 == 0 {
		return n == 3
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
