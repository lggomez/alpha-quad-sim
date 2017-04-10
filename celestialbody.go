package main

import (
	"fmt"
)

// CelestialBody - Celestial body representation
// Assumes a circular orbit with constant speed
type CelestialBody struct {
	PolarPosition     *PolarPoint
	CartesianPosition *Point
	name              string
	anglesPerDay      int8
	clockWiseOrbit    bool
}

// NewCelestialBody - Create a new point given x and y coordinates
func NewCelestialBody(initialCoord *PolarPoint, anglesPerDay int8, clockWiseOrbit bool, name string) *CelestialBody {
	newBody := &CelestialBody{
		PolarPosition:  initialCoord,
		name:           name,
		clockWiseOrbit: clockWiseOrbit,
		anglesPerDay:   anglesPerDay,
	}
	newBody.CartesianPosition = newBody.PolarPosition.ToCartesianPoint()

	return newBody
}

// AdvancePosition - Calculate the position of the body
// after x days and update its coordinates. Assumes a
// constant orbit so the ratio isn't updated
func (c *CelestialBody) AdvancePosition(days int) {
	angle := c.anglesPerDay

	if c.clockWiseOrbit {
		angle = -angle
	}

	c.PolarPosition.Angle += float32(angle)
	c.CartesianPosition = c.PolarPosition.ToCartesianPoint()
}

// ToString - Creates a string representation of the body
func (c *CelestialBody) ToString() string {
	return fmt.Sprintf("%s: %s/%s", c.name, c.PolarPosition.ToString(), c.CartesianPosition.ToString())
}
