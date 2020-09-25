package allyourbase

import "errors"

func ConvertToBase(inBase int, inDigits []int, outBase int) (digits []int, err error) {
	if inBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	inNum, err := toInt(inBase, inDigits)
	if err != nil {
		return digits, err
	}
	if inNum == 0 {
		return []int{0}, nil
	}
	return baseDigits(inNum, outBase), nil
}

func toInt(inBase int, inDigits []int) (int, error) {
	n := 0
	l := len(inDigits)
	for i, digit := range inDigits {
		if digit < 0 || digit >= inBase {
			return n, errors.New("all digits must satisfy 0 <= d < input base")
		}
		n += digit * pow(inBase, l-i-1)
	}
	return n, nil
}

func baseDigits(n, b int) []int {
	var bits []int
	q, r := n, 0
	for q > 0 {
		q, r = q/b, q%b
		bits = append(bits, r)
	}
	return reverseSlice(bits)
}

func reverseSlice(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func pow(x, y int) int {
	switch y {
	case 0:
		return 1
	case 1:
		return x
	default:
		r := x
		for i := 0; i < y-1; i++ {
			r *= x
		}
		return r
	}

}
