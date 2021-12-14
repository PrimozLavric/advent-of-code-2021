package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PrimozLavric/advent-of-code-2021/day-7/internal/util"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// parseFishFile reads fish internal timers from the provided file and generates reproduction histogram.
func (app *application) parseCrabPositionsFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	// Defer close the file.
	defer func() {
		err = file.Close()

		if err != nil {
			app.log.Printf("Failed to close file: %s\n", filePath)
		}
	}()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if len(records) != 1 {
		return nil, errors.New("bad file format, expected single line with comma separated values")
	}

	var crabPositions []int

	for i, record := range records[0] {
		position, err := strconv.Atoi(record)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not convert record %d to integer (%s)", i+1, err.Error()))
		}

		crabPositions = append(crabPositions, position)
	}

	return crabPositions, nil
}

// computeFuelConsumptionPartOne computes minimal amount of fuel required to align crabs on the same position. Each
// move costs 1 fuel unit.
func computeFuelConsumptionPartOne(crabPositions []int) int {
	// Optimal alignment position is median of crab positions.
	medianPosition := util.ComputeMedian(crabPositions)

	usedFuel := 0

	for _, position := range crabPositions {
		if position > medianPosition {
			usedFuel += position - medianPosition
		} else {
			usedFuel += medianPosition - position
		}
	}

	return usedFuel
}

// computeFuelConsumptionPartOne computes minimal amount of fuel required to align crabs on the same position. Each
// subsequent move is more costly. For example 4 moves of the same crab cost (1 + 2 + 3 + 4).
func computeFuelConsumptionPartTwo(crabPositions []int) int {
	// Optimal alignment position is mean of crab positions.
	mean := util.ComputeMean(crabPositions)

	usedFuel := 0

	for _, position := range crabPositions {
		positionDiff := mean - position
		if position > mean {
			positionDiff = position - mean
		}

		usedFuel += positionDiff * (positionDiff + 1) / 2
	}

	return usedFuel
}

func main() {
	var fishFile = flag.String("file", "input.txt", "Crab positions file.")
	flag.Parse()

	app := application{log: log.Default()}

	crabPositions, err := app.parseCrabPositionsFile(*fishFile)

	if err != nil {
		app.log.Fatalf("Encountered error during crab positions file parsing (%s).", err.Error())
	}

	fmt.Printf("Fuel used part one: %d\n", computeFuelConsumptionPartOne(crabPositions))
	fmt.Printf("Fuel used part two: %d\n", computeFuelConsumptionPartTwo(crabPositions))
}
