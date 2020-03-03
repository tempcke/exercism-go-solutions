package collatzconjecture

import "fmt"

// CollatzConjecture takes any positive integer n.
// If n is even, divide n by 2 to get n / 2.
// If n is odd, multiply n by 3 and add 1 to get 3n + 1
// repeat and return how long it takes to reach 1
func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, fmt.Errorf("positive int required %d given", input)
	}
	return collatz(input, 0), nil
}

func collatz(n, i int) int {
	if n == 1 {
		return i
	}
	return collatz(changeN(n), i+1)
}

func changeN(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}
