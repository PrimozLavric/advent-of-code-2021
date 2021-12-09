package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PrimozLavric/advent-of-code-2021/day-6/internal/simulator"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// parseFishFile reads fish internal timers from the provided file and generates reproduction histogram.
func (app *application) parseFishFile(filePath string) ([simulator.MaxInternalTimer]int, error) {
	file, err := os.Open(filePath)

	var timerHistogram [simulator.MaxInternalTimer]int

	if err != nil {
		return timerHistogram, err
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
		return timerHistogram, errors.New("bad file format, expected single line with comma separated values")
	}

	for i, record := range records[0] {
		timerValue, err := strconv.Atoi(record)

		if err != nil {
			return timerHistogram, errors.New(fmt.Sprintf("could not convert record %d to integer (%s)", i+1, err.Error()))
		}

		timerHistogram[timerValue]++
	}

	return timerHistogram, nil
}

func main() {
	var fishFile = flag.String("file", "input.txt", "Lanternfish internal timers file.")
	flag.Parse()

	app := application{log: log.Default()}

	timerHistogram, err := app.parseFishFile(*fishFile)

	if err != nil {
		app.log.Fatalf("Encountered error during lanternfish internal timers file parsing (%s).", err.Error())
	}

	sim := simulator.NewReproductionSimulator(timerHistogram)

	// Simulate 80 days.
	sim.Simulate(80)
	fmt.Printf("There is %d fish after 80 days.\n", sim.CountFish())

	// Simulate 176 more days to get to 256 days.
	sim.Simulate(176)
	fmt.Printf("There is %d fish after 256 days.\n", sim.CountFish())
}
