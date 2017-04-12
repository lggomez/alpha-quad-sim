package point

import (
	"math"
	"strconv"
	"fmt"
)

// PolarPoint - Polar coordinate point
type PolarPoint struct {
	Ratio float32
	Angle float32
}

// NewPolarPoint - Create a new polar point given x and y coordinates
func NewPolarPoint(ratio float32, angle float32) *PolarPoint {
	return &PolarPoint{
		Ratio: ratio,
		Angle: angle,
	}
}

// ToCartesianPoint - Returns a new cartesian point from the current polar point
func (p *PolarPoint) ToCartesianPoint() *Point {
	x := p.Ratio * float32(math.Cos(float64(p.Angle*0.0174533)))
	y := p.Ratio * float32(math.Sin(float64(p.Angle*0.0174533)))
	return &Point{
		X: x,
		Y: y,
	}
}

// ToString - Creates a string representation of the point
func (p *PolarPoint) ToString() string {
	r := strconv.FormatFloat(float64(p.Ratio), 'f', 4, 32)
	a := strconv.FormatFloat(float64(p.Angle), 'f', 4, 32)

	return fmt.Sprintf("(%s, %s°)", r, a)
}
