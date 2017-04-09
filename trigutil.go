package main

import (
	"math"
)

// IsPointInTriangle - Checks if a point is in a triangle formed by a given set of points
func IsPointInTriangle(p Point, p1 Point, p2 Point, p3 Point) bool {
	s := p1.Y*p3.X - p1.X*p3.Y + (p3.Y-p1.Y)*p.X + (p1.X-p3.X)*p.Y
	t := p1.X*p2.Y - p1.Y*p2.X + (p1.Y-p2.Y)*p.X + (p2.X-p1.X)*p.Y

	if (s < 0) != (t < 0) {
		return false
	}

	a := -p2.Y*p3.X + p1.Y*(p3.X-p2.X) + p1.X*(p2.Y-p3.Y) + p2.X*p3.Y

	if a < 0.0 {
		s = -s
		t = -t
		a = -a
	}

	return s > 0 && t > 0 && (s+t) <= a
}

// GetTriangleAreaByPoints - Gets the area of a triangle formed by a given set of points
func GetTriangleAreaByPoints(p1 Point, p2 Point, p3 Point) float32 {
	p := p1.X * (p2.Y - p3.Y)
	p += p2.X * (p3.Y - p1.Y)
	p += p3.X * (p1.Y - p2.Y)
	return float32(math.Abs(float64(p / 2)))
}