package simulator

const MaxInternalTimer = 9
const ReproductionCycle = 6

// ReproductionSimulator used to simulate fish reproduction.
type ReproductionSimulator struct {
	reproductionHistogram [MaxInternalTimer]int
	histogramOffset       uint8
}

// NewReproductionSimulator creates and initializes ReproductionSimulator with provided reproduction histogram
func NewReproductionSimulator(reproductionHistogram [MaxInternalTimer]int) *ReproductionSimulator {
	sim := ReproductionSimulator{reproductionHistogram: reproductionHistogram, histogramOffset: 0}

	return &sim
}

// SimulateDay simulates singled day of reproduction.
func (sim *ReproductionSimulator) SimulateDay() {
	reproducingFish := sim.reproductionHistogram[0]

	for i := 1; i < len(sim.reproductionHistogram); i++ {
		sim.reproductionHistogram[i-1] = sim.reproductionHistogram[i]
	}

	sim.reproductionHistogram[ReproductionCycle] += reproducingFish
	sim.reproductionHistogram[len(sim.reproductionHistogram)-1] = reproducingFish
}

// Simulate simulates provided number of days of reproduction.
func (sim *ReproductionSimulator) Simulate(numDays int) {
	for i := 0; i < numDays; i++ {
		sim.SimulateDay()
	}
}

// CountFish counts the current number of fish.
func (sim *ReproductionSimulator) CountFish() int {
	fishCount := 0

	for i := 0; i < len(sim.reproductionHistogram); i++ {
		fishCount += sim.reproductionHistogram[i]
	}

	return fishCount
}
