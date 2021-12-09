package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// readDepthReport reads depth report data from the file located on the provided path. It expects file to have a single
// unsigned depth entry per line.
func (app *application) readDepthReport(path string) ([]uint, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	// Defer close the file.
	defer func() {
		err = file.Close()

		if err != nil {
			app.log.Printf("Failed to close file: %s\n", path)
		}
	}()

	var depthReportData []uint

	// Read depth data. Each line should contain one entry.
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		depth, err := strconv.Atoi(fileScanner.Text())

		if err != nil {
			return nil, errors.New(fmt.Sprintf("bad input file format. Could not convert line '%s' to int", fileScanner.Text()))
		}

		if depth < 0 {
			return nil, errors.New(fmt.Sprintf("bad input file. Read negative depth '%d'", depth))
		}

		depthReportData = append(depthReportData, uint(depth))
	}

	return depthReportData, nil
}

// countSequentialIncrements counts number of neighbour entries in a slice where sequence[i - 1] < sequence[i].
func countSequentialIncrements(sequence []uint) uint {
	numIncrements := uint(0)

	for i := 1; i < len(sequence); i++ {
		if sequence[i-1] < sequence[i] {
			numIncrements++
		}
	}

	return numIncrements
}

// countSequentialWindowSumIncrements counts number if entries where sum of sequence[i - windowSize:i] < sequence[(i+1) - windowSize:i+1]
func countSequentialWindowSumIncrements(sequence []uint, windowSize uint) (uint, error) {
	if windowSize == 0 {
		return 0, errors.New(fmt.Sprintf("invalid window size %d", windowSize))
	}

	// Special case where window size is larger or equal to sequence length.
	if windowSize >= uint(len(sequence)) {
		return 0, nil
	}

	// Compute initial window sum.
	currentSum := uint(0)

	for i := uint(0); i < windowSize; i++ {
		currentSum += sequence[i]
	}

	// Compute number of running window sum increments.
	numIncrements := uint(0)

	for i := windowSize; i < uint(len(sequence)); i++ {
		lastSum := currentSum
		currentSum -= sequence[i-windowSize]
		currentSum += sequence[i]

		if currentSum > lastSum {
			numIncrements++
		}
	}

	return numIncrements, nil
}

func main() {
	var depthFilePath = flag.String("file", "input.txt", "Path to the depth report file.")
	var windowSize = flag.Uint("window-size", 1, "Size of the sum window.")
	flag.Parse()

	app := application{log.Default()}

	// Read depth report from file provided by the user.
	depthData, err := app.readDepthReport(*depthFilePath)

	if err != nil {
		app.log.Fatalf("Failed to read depth report (%s)\n", err.Error())
	}

	// Use exercise part 1 solution when window size is 1.
	if *windowSize == 1 {
		depthIncrementCount := countSequentialIncrements(depthData)

		fmt.Printf("There are %d depth increments.\n", depthIncrementCount)
	} else {
		depthIncrementCount, err := countSequentialWindowSumIncrements(depthData, *windowSize)

		if err != nil {
			app.log.Fatalf("Encountered an error while counting the window sum increments (%s).", err.Error())
		}

		fmt.Printf("With window size %d, there are %d depth sum increments.\n", *windowSize, depthIncrementCount)
	}
}
