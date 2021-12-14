package octopus

import (
	"errors"
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

// offset 2D coordinate offset.
type offset struct {
	dx int
	dy int
}

// neighborOffsets array of neighbour cell offsets (on a 2D grid)
var neighborOffsets = [8]offset{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

// Simulator that is used to simulate octopus flashing.
type Simulator struct {
	initEnergies [][]uint
	energies     [][]uint
	flashed      mapset.Set
	stepCount    int
	flashCount   int
}

// NewSimulator creates the Simulator and initializes it with the provided energies.
func NewSimulator(initEnergies [][]uint) (*Simulator, error) {
	if len(initEnergies) == 0 {
		return nil, errors.New("energies matrix empty, cannot create Simulator with no energies")
	}
	// Validate that all rows are of same length.
	rowLen := len(initEnergies[0])

	for _, energyRow := range initEnergies {
		if rowLen != len(energyRow) {
			return nil, errors.New("encountered different length energy rows, cannot create Simulator")
		}
	}

	if rowLen == 0 {
		return nil, errors.New("energies matrix empty, cannot create Simulator with no energies")
	}

	return &Simulator{initEnergies: initEnergies, energies: copyEnergies(initEnergies), flashed: mapset.NewSet(), stepCount: 0, flashCount: 0}, nil
}

// Reset resets simulator to initial state.
func (sim *Simulator) Reset() {
	sim.energies = copyEnergies(sim.initEnergies)
	sim.stepCount = 0
	sim.flashCount = 0
}

// SimulateStep simulates a single step of octopus flashing simulation.
func (sim *Simulator) SimulateStep() bool {
	for x := range sim.energies {
		for y := range sim.energies[0] {
			sim.incrementEnergy(x, y)
		}
	}

	// Reset energies greater than 9 to 0
	for x := range sim.energies {
		for y := range sim.energies[0] {
			if sim.energies[x][y] > 9 {
				sim.energies[x][y] = 0
			}
		}
	}

	numFlashed := sim.flashed.Cardinality()
	sim.flashCount += numFlashed
	sim.flashed.Clear()
	sim.stepCount += 1

	return numFlashed == len(sim.energies)*len(sim.energies[0])
}

// SimulateNSteps simulates N steps of octopus flashing simulation.
func (sim *Simulator) SimulateNSteps(n int) {
	for i := 0; i < n; i++ {
		sim.SimulateStep()
	}
}

// SimulateUntilAllFlash runs the simulation until all octopuses flash.
func (sim *Simulator) SimulateUntilAllFlash() {
	for !sim.SimulateStep() {
		continue
	}
}

// StepCount retrieves number of already executed simulation steps.
func (sim *Simulator) StepCount() int {
	return sim.stepCount
}

// FlashCount retrieves total number of octopus flashes.
func (sim *Simulator) FlashCount() int {
	return sim.flashCount
}

// isCoordinateValid checks if the provided coordinate is within the grid.
func (sim *Simulator) isCoordinateValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(sim.energies) && y < len(sim.energies[0])
}

// incrementEnergy increments energy of octopus on the given coordinate and simulates flash if its energy exceeds 9.
func (sim *Simulator) incrementEnergy(x, y int) {
	sim.energies[x][y]++

	// If energy exceeded 9 simulate flash.
	if sim.energies[x][y] > 9 {
		sim.simulateFlash(x, y)
	}
}

// simulateFlash simulates flash of the octopus if it hasn't flashed yet in the current step of simulation.
// 					 		 Flashing increases energy of all neighbour octopuses by 1.
func (sim *Simulator) simulateFlash(x, y int) {
	if sim.energies[x][y] <= 9 {
		panic(fmt.Sprintf("logic error, tried to simulate flash on field [%d][%d], but energy level is not greater than 9", x, y))
	}

	// Check if octopus already flashed this iteration.
	uniqueId := x*len(sim.energies) + y
	if sim.flashed.Contains(uniqueId) {
		return
	}

	sim.flashed.Add(uniqueId)

	// Increase energy of all neighbors by 1.
	for _, nOffset := range neighborOffsets {
		nX, nY := x+nOffset.dx, y+nOffset.dy

		if sim.isCoordinateValid(nX, nY) {
			sim.incrementEnergy(nX, nY)
		}
	}
}

// copyEnergies makes a copy of the provided energies slices.
func copyEnergies(energies [][]uint) [][]uint {
	energiesCopy := make([][]uint, len(energies))

	for i, row := range energies {
		energiesCopy[i] = make([]uint, len(energies[0]))
		copy(energiesCopy[i], row)
	}

	return energiesCopy
}
