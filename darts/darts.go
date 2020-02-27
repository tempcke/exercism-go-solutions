package darts

import "math"

type ring struct {
	radius float64
	points int
}

var rings = [3]ring{
	ring{1, 10},
	ring{5, 5},
	ring{10, 1},
}

// Score the dart throw 10 for inner ring, 5 for mid, 1 for outer
func Score(x float64, y float64) int {
	r := math.Sqrt(x*x + y*y)
	for _, ring := range rings {
		if r <= ring.radius {
			return ring.points
		}
	}
	return 0
}
