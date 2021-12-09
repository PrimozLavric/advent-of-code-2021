package test

import (
	"testing"

	"github.com/PrimozLavric/advent-of-code-2021/day-6/internal/simulator"
)

func TestReproductionSimulator(t *testing.T) {

	testInput := [simulator.MaxInternalTimer]int{0, 1, 1, 2, 1, 0, 0, 0, 0}

	sim := simulator.NewReproductionSimulator(testInput)

	if sim.CountFish() != 5 {
		t.Errorf("expected 5 fishes, actual %d", sim.CountFish())
	}

	sim.Simulate(18)

	if sim.CountFish() != 26 {
		t.Errorf("expected 26 fishes, actual %d", sim.CountFish())
	}

	sim.Simulate(62)

	if sim.CountFish() != 5934 {
		t.Errorf("expected 5934 fishes, actual %d", sim.CountFish())
	}
}
