package test

import (
	"strconv"
	"testing"

	"github.com/PrimozLavric/advent-of-code-2021/day-3/internal/report_parser"
)

func parseExampleData(exampleData []string) []uint32 {
	var parsedData []uint32

	for _, entry := range exampleData {
		binaryValue, _ := strconv.ParseUint(entry, 2, 5)
		parsedData = append(parsedData, uint32(binaryValue))
	}

	return parsedData
}

var exampleData = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

var parsedExampleData = parseExampleData(exampleData)

func TestExampleComputeGammaAndEpsilon(t *testing.T) {
	gamma, epsilon := report_parser.FindGammaAndEpsilonRate(parsedExampleData, uint8(len(exampleData[0])))

	if gamma != 22 {
		t.Error("gamma should be 22, but got", gamma)
	}

	if epsilon != 9 {
		t.Error("epsilon should be 9, but got", epsilon)
	}
}

func TestExampleFindOxygenGeneratorRating(t *testing.T) {
	oxygenGenRating := report_parser.FindOxygenGeneratorRating(parsedExampleData, uint8(len(exampleData[0])))

	if oxygenGenRating != 23 {
		t.Error("result should be 23, but got", oxygenGenRating)
	}
}

func TestExampleFindCO2ScrubberRating(t *testing.T) {
	oxygenGenRating := report_parser.FindCO2ScrubberRating(parsedExampleData, uint8(len(exampleData[0])))

	if oxygenGenRating != 10 {
		t.Error("result should be 10, but got", oxygenGenRating)
	}
}
