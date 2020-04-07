package sieve

// Sieve of Eratosthenes to find all the primes from 2 up to a given number
func Sieve(limit int) []int {
	s := make([]bool, limit+1)
	p := make([]int, 0, limit)

	s[0], s[1] = true, true

	for i := 2; i <= limit; i++ {
		if !s[i] {
			p = append(p, i)
			for j := i; j <= int(limit/i); j++ {
				s[i*j] = true
			}
		}
	}
	return p
}
