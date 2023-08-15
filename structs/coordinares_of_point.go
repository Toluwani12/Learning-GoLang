package structs

import "math"

type Point struct{
	X, Y float64
}

func (P Point) Abs() float64{
	return math.Sqrt(P.X*P.X +P.Y*P.Y)
}

// Scale multiplies the coordinates of the Point by the scale factor
func (p Point) Scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}
