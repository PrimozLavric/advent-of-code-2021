package main

import (
	"testing"

	"day-11/internal/octopus"
)

// TestOctopusSimulatorPartOne tests simulation for the first part of the exercise (number of flashes after N steps)
func TestOctopusSimulatorPartOne(t *testing.T) {
	testData := [][]uint{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}

	sim, err := octopus.NewSimulator(testData)

	if err != nil {
		t.Errorf("Encountered error while trying to initialize simulator (%s)", err.Error())
	}

	sim.SimulateNSteps(10)

	if sim.StepCount() != 10 {
		t.Errorf("Step counting is not working correctly.")
	}

	if sim.FlashCount() != 204 {
		t.Errorf("Expected 204 flashes after 10 steps, but got %d flashes", sim.FlashCount())
	}

	sim.SimulateNSteps(90)

	if sim.StepCount() != 100 {
		t.Errorf("Step counting is not working correctly.")
	}

	if sim.FlashCount() != 1656 {
		t.Errorf("Expected 1656 flashes after 10 steps, but got %d flashes", sim.FlashCount())
	}
}

// TestOctopusSimulatorPartTwo tests simulation for the second part of the exercise (number of steps until all flash)
func TestOctopusSimulatorPartTwo(t *testing.T) {
	testData := [][]uint{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}

	sim, err := octopus.NewSimulator(testData)

	if err != nil {
		t.Errorf("Encountered error while trying to initialize simulator (%s)", err.Error())
	}

	sim.SimulateUntilAllFlash()

	if sim.StepCount() != 195 {
		t.Errorf("Expected %d steps before simultaneous flash, actual %d", 195, sim.StepCount())
	}
}
