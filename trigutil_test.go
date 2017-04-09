package main

import (
	"testing"
	"fmt"
)

// TestIsPointInTriangle - Test the point in triangle area verificaion
func TestIsPointInTriangle(t *testing.T) {
	point0 := NewPoint(50, 25)
	point1 := NewPoint(0, 0)
	point2 := NewPoint(50, 50)
	point3 := NewPoint(100, 0)

	if !IsPointInTriangle(*point0, *point1, *point2, *point3) {
		t.Error("Expected true got false")
	}

	point0 = NewPoint(95, 1)
	point1 = NewPoint(0, 0)
	point2 = NewPoint(50, 50)
	point3 = NewPoint(100, 0)

	if !IsPointInTriangle(*point0, *point1, *point2, *point3) {
		t.Error("Expected true got false")
	}

	point0 = NewPoint(5, 1)
	point1 = NewPoint(0, 0)
	point2 = NewPoint(50, 50)
	point3 = NewPoint(100, 0)

	if !IsPointInTriangle(*point0, *point1, *point2, *point3) {
		t.Error("Expected true got false")
	}

	point0 = NewPoint(50, 49)
	point1 = NewPoint(0, 0)
	point2 = NewPoint(50, 50)
	point3 = NewPoint(100, 0)

	if !IsPointInTriangle(*point0, *point1, *point2, *point3) {
		t.Error("Expected true got false")
	}

	point0 = NewPoint(101, 1)
	point1 = NewPoint(0, 0)
	point2 = NewPoint(50, 50)
	point3 = NewPoint(100, 0)

	if IsPointInTriangle(*point0, *point1, *point2, *point3) {
		t.Error("Expected false got true")
	}

	point0 = NewPoint(-1, 1)
	point1 = NewPoint(0, 0)
	point2 = NewPoint(50, 50)
	point3 = NewPoint(100, 0)

	if IsPointInTriangle(*point0, *point1, *point2, *point3) {
		t.Error("Expected false got true")
	}

	point0 = NewPoint(50, 51)
	point1 = NewPoint(0, 0)
	point2 = NewPoint(50, 50)
	point3 = NewPoint(100, 0)

	if IsPointInTriangle(*point0, *point1, *point2, *point3) {
		t.Error("Expected false got true")
	}
}

// TestGetTriangleAreaByPoints - Test the triangle area calculation
func TestGetTriangleAreaByPoints(t *testing.T) {
	point1 := NewPoint(0, 0)
	point2 := NewPoint(50, 50)
	point3 := NewPoint(100, 0)

	area := GetTriangleAreaByPoints(*point1, *point2, *point3)
	if !AreApproximatelyEqual(area, 2500) {
		t.Error(fmt.Sprintf("Expected 2500 got %f", area))
	}

	point1 = NewPoint(-100, 0)
	point2 = NewPoint(50, 50)
	point3 = NewPoint(0, 0)

	area = GetTriangleAreaByPoints(*point1, *point2, *point3)
	if !AreApproximatelyEqual(area, 2500) {
		t.Error(fmt.Sprintf("Expected 2500 got %f", area))
	}

	point1 = NewPoint(0, 0)
	point2 = NewPoint(0, 0)
	point3 = NewPoint(0, 0)

	area = GetTriangleAreaByPoints(*point1, *point2, *point3)
	if !AreApproximatelyEqual(area, 0) {
		t.Error(fmt.Sprintf("Expected 0 got %f", area))
	}

	point1 = NewPoint(453, 354)
	point2 = NewPoint(2345, 253)
	point3 = NewPoint(1235, 553)

	area = GetTriangleAreaByPoints(*point1, *point2, *point3)
	if !AreApproximatelyEqual(area, 227745) {
		t.Error(fmt.Sprintf("Expected 227745 got %f", area))
	}

	point1 = NewPoint(-5, -10)
	point2 = NewPoint(-4, -10)
	point3 = NewPoint(-3, -8)

	area = GetTriangleAreaByPoints(*point1, *point2, *point3)
	if !AreApproximatelyEqual(area, 1) {
		t.Error(fmt.Sprintf("Expected 1 got %f", area))
	}
}
