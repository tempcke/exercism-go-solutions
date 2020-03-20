package pythagorean

import "math"

// Triplet is a set of 3 ints which make a right triangle such as 3,4,5
type Triplet [3]int

// Range finds all Triplet's within the provided min and max range
func Range(min, max int) []Triplet {
	ts := make([]Triplet, 0, max-min)
	for a := min; a < max-2; a++ {
		if t, ok := findTripletInRange(a, max); ok {
			ts = append(ts, t)
		}
	}
	return ts
}

// Sum finds all Triplet's whose sumed value matches the provided number
func Sum(p int) []Triplet {
	ts := make([]Triplet, 0)

	upperLimit := p / 3

	for a := 1; a < upperLimit; a++ {
		if t, ok := findTripletSum(a, p); ok {
			ts = append(ts, t)
		}
	}
	return ts
}

func findTripletInRange(a, max int) (t Triplet, ok bool) {
	sqMax := max * max
	for b := a + 1; a*a+b*b <= sqMax; b++ {
		sqC := a*a + b*b
		c := int(math.Sqrt(float64(sqC)))
		if c*c == sqC {
			return Triplet{a, b, c}, true
		}
	}
	return
}

func findTripletSum(a, sum int) (t Triplet, ok bool) {
	for b, c := a+1, a+2; a+b+c < sum; b++ {
		sqC := a*a + b*b
		c = int(math.Sqrt(float64(sqC)))
		if a+b+c == sum && c*c == sqC {
			return Triplet{a, b, c}, true
		}
	}
	return
}
