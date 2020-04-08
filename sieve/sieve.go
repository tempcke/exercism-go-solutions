package sieve

// Sieve of Eratosthenes to find all the primes from 2 up to a given number
func Sieve(limit int) []int {
	marked := make([]bool, limit+1)
	primes := make([]int, 0, limit/2)

	for i := 2; i <= limit; i++ {
		if marked[i] {
			continue
		}
		primes = append(primes, i)
		for mult := i * 2; mult <= limit; mult += i {
			marked[mult] = true
		}
	}
	return primes
}
