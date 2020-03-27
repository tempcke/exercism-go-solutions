package grains

import "fmt"

// Square returns the number of grains on the nth chess square
// where we start with 1 and then double it each time
// 2^n == 1 << n
// 2^0 == 1 << 0 ==    1 == 1
// 2^1 == 1 << 1 ==   10 == 2
// 2^2 == 1 << 2 ==  100 == 4
// 2^3 == 1 << 3 == 1000 == 8
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("n must be a positive integer")
	}
	return 1 << (n - 1), nil
}

// Total returns the sum of Square(1):Square(64)
// 2^0 == 1 << 0  ==    1 ==  1
// 2^1 == 1 << 1  ==   10 ==  2
// 2^2 == 1 << 2  ==  100 ==  4
// 2^3 == 1 << 3  == 1000 ==  8
//                   ==========
// 2^4 -1 == 16-1 == 1111 == 15
func Total() uint64 {
	return (1 << 64) - 1
}
