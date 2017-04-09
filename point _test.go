package main

import (
	"testing"
)

func TestToCartesianPoint(t *testing.T) {
	point := NewPolarPoint(0, 0).ToCartesianPoint()
	if !AreApproximatelyEqual(point.Y, 0) || !AreApproximatelyEqual(point.X, 0) && point.X != 0 {
		t.Error("Expected (0,0), got ", point.ToString())
	}

	point = NewPolarPoint(90, 90).ToCartesianPoint()
	if !AreApproximatelyEqual(point.Y, 90) || !AreApproximatelyEqual(point.X, 0) {
		t.Error("Expected (0,90), got ", point.ToString())
	}

	point = NewPolarPoint(90, -90).ToCartesianPoint()
	if !AreApproximatelyEqual(point.Y, -90) || !AreApproximatelyEqual(point.X, 0) {
		t.Error("Expected (0,-90), got ", point.ToString())
	}

	point = NewPolarPoint(548, 956).ToCartesianPoint()
	if !AreApproximatelyEqual(point.X, -306.4347) || !AreApproximatelyEqual(point.Y, -454.3146) {
		t.Error("Expected (-306.4347,-454.3146), got ", point.ToString())
	}
}
