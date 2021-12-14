package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PrimozLavric/advent-of-code-2021/day-11/internal/octopus"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// readEnergyLevelsFromFile reads octopus energy levels from the input file.
func (app *application) readEnergyLevelsFromFile(filePath string) ([][]uint, error) {
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

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var energyLevels [][]uint

	for rowNum := 1; fileScanner.Scan(); rowNum++ {
		rowEnergies := make([]uint, 0, len(fileScanner.Text()))

		for _, char := range fileScanner.Text() {
			val, err := strconv.Atoi(string(char))

			if err != nil {
				return nil, errors.New(fmt.Sprintf("failed to parse row %d in energy levels file (%s)", rowNum, err.Error()))
			}

			rowEnergies = append(rowEnergies, uint(val))
		}

		energyLevels = append(energyLevels, rowEnergies)
	}

	return energyLevels, nil
}

func main() {
	var energiesFile = flag.String("file", "input.txt", "Octopus energy file.")
	flag.Parse()

	app := application{log: log.Default()}

	energyLevels, err := app.readEnergyLevelsFromFile(*energiesFile)

	if err != nil {
		app.log.Fatalf("Encountered error during octopus energy file parsing (%s).", err.Error())
	}

	sim, err := octopus.NewSimulator(energyLevels)

	if err != nil {
		app.log.Fatalf("Encountered error during simulator creation (%s).", err.Error())
	}

	// Part one:
	sim.SimulateNSteps(100)
	fmt.Printf("Simulating 100 steps. Octopuses flashed %d times.\n", sim.FlashCount())

	// Part two:
	sim.Reset()
	sim.SimulateUntilAllFlash()
	fmt.Printf("It took %d steps before all octopuses flashed in a single step.\n", sim.StepCount())
}
