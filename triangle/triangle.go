// Package triangle three sided polygon
package triangle

import "math"

// Kind - type of triangle such as equilateral, isosceles, or scalene
type Kind string

// Types of triangles
const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case isNat(a, b, c):
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}

func isNat(a, b, c float64) bool {
	if !hasValidSides(a, b, c) {
		return true
	}
	if !sumOfAnyTwoSidesIsGreaterThanThird(a, b, c) {
		return true
	}
	return false
}

func hasValidSides(sides ...float64) bool {
	for _, length := range sides {
		if !isPositiveRealNumber(length) {
			return false
		}
	}
	return true
}

func isPositiveRealNumber(s float64) bool {
	return !math.IsInf(s, 0) && !math.IsNaN(s) && s > 0
}

func sumOfAnyTwoSidesIsGreaterThanThird(a, b, c float64) bool {
	return a+b >= c && b+c >= a && a+c >= b
}
