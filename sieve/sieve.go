package sieve

func Sieve(limit int) []int {
	s := make([]bool, limit+1)
	p := make([]int, 0, limit)

	s[0], s[1] = true, true

	for i := 2; i <= limit; i++ {
		isPrime := !s[i]
		if isPrime {
			p = append(p, i)
			for j := i; j <= int(limit/i); j++ {
				notPrime := i * j
				s[notPrime] = true
			}
		}
	}
	return p
}
