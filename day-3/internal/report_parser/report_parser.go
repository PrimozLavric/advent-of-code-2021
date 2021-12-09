package report_parser

import "github.com/PrimozLavric/advent-of-code-2021/day-3/internal/util"

// FindGammaAndEpsilonRate computes gamma and epsilon value from the given report.
func FindGammaAndEpsilonRate(diagnosticReport []uint32, entryBitSize uint8) (uint32, uint32) {
	// Count number of ones for each bit.
	bitOnesCounts := make([]uint32, entryBitSize)

	for i := uint8(0); i < entryBitSize; i++ {
		bitOnesCounts[i] = util.CountSetIthBits(diagnosticReport, i)
	}

	// Compute gamma.
	var gamma = uint32(0)
	for i := uint8(0); i < entryBitSize; i++ {
		if int(bitOnesCounts[i]) > len(diagnosticReport)/2 {
			gamma |= 1 << i
		}
	}

	// Compute epsilon directly from gamma.
	epsilon := ^gamma & (1<<(entryBitSize) - 1)

	return gamma, epsilon
}

// FindOxygenGeneratorRating computes oxygen generator rating from the given report.
func FindOxygenGeneratorRating(diagnosticReport []uint32, entryBitSize uint8) uint32 {
	filteredReport := diagnosticReport
	// Filter until only one value is left, or we run out of bits.
	for i := int(entryBitSize) - 1; i >= 0; i-- {
		filteredReport = util.FilterDataBasedOnMostCommonBitValue(filteredReport, uint8(i))

		if len(filteredReport) <= 1 {
			break
		}
	}

	if len(filteredReport) == 0 {
		panic("Logic error. Filtering diagnosticReport resulted into an empty report. This should never happen.")
	}

	// More than 1 entry may be left. We always pick the first one
	return filteredReport[0]
}

// FindCO2ScrubberRating computes CO2 scrubber rating from the given report.
func FindCO2ScrubberRating(diagnosticReport []uint32, entryBitSize uint8) uint32 {
	filteredReport := diagnosticReport
	// Filter until only one value is left, or we run out of bits.
	for i := int(entryBitSize) - 1; i >= 0; i-- {
		filteredReport = util.FilterDataBasedOnLeastCommonBitValue(filteredReport, uint8(i))

		if len(filteredReport) <= 1 {
			break
		}
	}

	if len(filteredReport) == 0 {
		panic("Logic error. Filtering diagnosticReport resulted into an empty report. This should never happen.")
	}

	// More than 1 entry may be left. We always pick the first one
	return filteredReport[0]
}
