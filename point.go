package main

import (
	"strconv"
	"fmt"
)

// Point - Cartesian coordinate point
type Point struct {
	X float32
	Y float32
}

// NewPoint - Create a new point given x and y coordinates
func NewPoint(x float32, y float32) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// ToString - Creates a string representation of the point
func (p *Point) ToString() string {
	x := strconv.FormatFloat(float64(p.X), 'f', 4, 32)
	y := strconv.FormatFloat(float64(p.Y), 'f', 4, 32)

	return fmt.Sprintf("(%s, %s)", x, y)
}