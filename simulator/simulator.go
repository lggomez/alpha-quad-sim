package simulator

import (
	"fmt"
	"alpha-quad-sim/point"
	"alpha-quad-sim/database"
	"alpha-quad-sim/trigutil"
)

// AlphaQuadSimulator - Alpha quadrant planet simulation
type AlphaQuadSimulator struct {
	Sun            *CelestialBody
	Ferengi        *CelestialBody
	Betasoide      *CelestialBody
	Vulcano        *CelestialBody
	ClimateMap     map[string]int
	day            int
	currentClimate string
}

// SimulatorConfig - Simulation configuration flags
type SimulatorConfig struct {
	ReportToConsole bool
	PersistClimates bool
}

func NewSimluatorConfig(reportToConsole bool, persistClimates bool) *SimulatorConfig {
	return &SimulatorConfig{
		ReportToConsole: reportToConsole,
		PersistClimates: persistClimates,
	}
}

// NewSimulation - Create a new simulation starting from default coordinates
func NewSimulation() *AlphaQuadSimulator {
	referencePoint := point.NewPolarPoint(0, 0)
	ferengiStartingPoint := point.NewPolarPoint(500, 90)
	betasoideStartingPoint := point.NewPolarPoint(2000, 90)
	vulcanoStartingPoint := point.NewPolarPoint(1000, 90)

	climateMap := make(map[string]int)

	climateMap["regular"] = 0
	climateMap["dry"] = 0
	climateMap["optimal"] = 0
	climateMap["rain"] = 0

	return &AlphaQuadSimulator{
		Sun:            NewCelestialBody(referencePoint, 0, true, "Sun"),
		Ferengi:        NewCelestialBody(ferengiStartingPoint, 1, true, "Ferengi"),
		Betasoide:      NewCelestialBody(betasoideStartingPoint, 3, true, "Betasoide"),
		Vulcano:        NewCelestialBody(vulcanoStartingPoint, 5, false, "Vulcano"),
		day:            0,
		ClimateMap:     climateMap,
		currentClimate: "",
	}
}

// Advance - Advance a day in the simulation and update the positions of all planets
func (a *AlphaQuadSimulator) Advance(days int) {
	a.day += days
	a.Vulcano.AdvancePosition(days)
	a.Betasoide.AdvancePosition(days)
	a.Ferengi.AdvancePosition(days)
}

// ChangeClimate - Update climate
func (a *AlphaQuadSimulator) ChangeClimate(climate string) {
	if (climate != a.currentClimate) && (climate != "") {
		a.currentClimate = climate
		a.ClimateMap[a.currentClimate] = a.ClimateMap[a.currentClimate] + 1
	}
}

// PrintAsString - Print the system planet's current positions
func (a *AlphaQuadSimulator) PrintAsString() {
	fmt.Printf("Day: %v\n", a.day)
	fmt.Println(a.Vulcano.ToString())
	fmt.Println(a.Betasoide.ToString())
	fmt.Println(a.Ferengi.ToString())
}

// Simulate - Simulate the system's planet current positions and prints the results, persisting them into the database
func (sim *AlphaQuadSimulator) Simulate(days int, cfg *SimulatorConfig) (string, error) {
	minAreaDay := 0
	var minArea float32
	lastDayClimate := "regular"
	error := error(nil)

	// Planets are aligned with the sun on day 0
	sim.ChangeClimate("dry")

	for i := 1; i <= days; i++ {
		sim.Advance(i)
		area := trigutil.GetTriangleAreaByPoints(*sim.Vulcano.CartesianPosition, *sim.Ferengi.CartesianPosition, *sim.Betasoide.CartesianPosition)

		// Determine min area for max rain intensity day
		if (area < minArea && area != 0) || i == 1 {
			minArea = area
		}

		if area == 0 {
			// Area tends to 0, planets are aligned
			if trigutil.GetTriangleAreaByPoints(*sim.Sun.CartesianPosition, *sim.Vulcano.CartesianPosition, *sim.Betasoide.CartesianPosition) == 0 {
				// Planets are aligned with the sun
				sim.ChangeClimate("dry")
			} else {
				sim.ChangeClimate("optimal")
			}
		} else {
			// Planets are in a triangle
			if trigutil.IsPointInTriangle(*sim.Sun.CartesianPosition, *sim.Vulcano.CartesianPosition, *sim.Ferengi.CartesianPosition, *sim.Betasoide.CartesianPosition) {
				sim.ChangeClimate("rain")
				minAreaDay = i
			} else {
				sim.ChangeClimate("regular")
			}
		}

		if i == days {
			lastDayClimate = sim.currentClimate
		}

		if cfg.PersistClimates {
			error = database.SaveClimate(i, sim.currentClimate)
		}
	}

	if cfg.ReportToConsole {
		fmt.Println("Max rain intensity reported on day:", minAreaDay)
		fmt.Println("Dry periods:", sim.ClimateMap["dry"])
		fmt.Println("Optimal climate periods:", sim.ClimateMap["optimal"])
	}

	return lastDayClimate, error
}
