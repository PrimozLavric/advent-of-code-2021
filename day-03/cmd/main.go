package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PrimozLavric/advent-of-code-2021/day-3/internal/report_parser"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// readDiagnosticReport reads diagnostic report and number of bits per entry from the given file.
func (app *application) readDiagnosticReport(filePath string) (uint8, []uint32, error) {
	const entryBase = 2

	file, err := os.Open(filePath)

	if err != nil {
		return 0, nil, err
	}

	// Defer close the file.
	defer func() {
		err = file.Close()

		if err != nil {
			app.log.Printf("Failed to close file: %s\n", filePath)
		}
	}()

	var reportEntries []uint32

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	entryBitSize := 0
	rowIdx := 0
	for fileScanner.Scan() {
		if entryBitSize == 0 {
			entryBitSize = len(fileScanner.Text())
		}

		entry, err := strconv.ParseUint(fileScanner.Text(), entryBase, entryBitSize)

		if err != nil {
			return 0, nil, errors.New(fmt.Sprintf("bad diagnostic report file (failed to parse row %d, %s)", rowIdx, err.Error()))
		}

		reportEntries = append(reportEntries, uint32(entry))

		rowIdx++
	}

	return uint8(entryBitSize), reportEntries, nil
}

func main() {
	var reportFile = flag.String("file", "input.txt", "File from which the diagnostic report will be read.")
	flag.Parse()

	app := application{log: log.Default()}

	// Read the report.
	entryBitSize, report, err := app.readDiagnosticReport(*reportFile)

	if err != nil {
		app.log.Fatalf("Failed to parse diagnostic report (%s)", err.Error())
		return
	}

	// Compute gamma nad epsilon.
	gamma, epsilon := report_parser.FindGammaAndEpsilonRate(report, entryBitSize)

	// Compute ratings
	oxygenGeneratorRating := report_parser.FindOxygenGeneratorRating(report, entryBitSize)
	co2ScrubberRating := report_parser.FindCO2ScrubberRating(report, entryBitSize)

	// Print out the results.
	fmt.Printf("Gamma: %d\nEpsilon: %d\nPower consumption: %d\n", gamma, epsilon, gamma*epsilon)
	fmt.Printf("Oxygen Generator Rating: %d\nCO2 Scrubber Rating: %d\nLife Support Rating: %d\n", oxygenGeneratorRating, co2ScrubberRating, oxygenGeneratorRating*co2ScrubberRating)
}
