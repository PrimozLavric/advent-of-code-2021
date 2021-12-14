package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/PrimozLavric/advent-of-code-2021/day-10/internal/syntax"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// readLinesFromFile reads lines from the file on the provided file path and returns them in a string slice.
func (app *application) readLinesFromFile(filePath string) ([]string, error) {
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

	var lines []string

	for i := 1; fileScanner.Scan(); i++ {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}

// computeAndPrintSyntaxErrorScore computes and prints total syntax error score (sum of syntax error score of every line)
func (app *application) computeAndPrintSyntaxErrorScore(lines []string) {
	a := syntax.NewAnalyzer()

	penaltySum := 0

	for i, line := range lines {
		penalty, err := a.ComputeSyntaxErrorScore(line)

		if err != nil {
			app.log.Fatalf(fmt.Sprintf("Encountered error while computing syntax error score for line %d: %s\n", i+1, err.Error()))
			return
		}

		penaltySum += penalty
	}

	fmt.Printf("Total syntax error score: %d\n", penaltySum)
}

// computeAndPrintAutocompleteScore computes and prints middle autocomplete score of all incomplete lines.
func (app *application) computeAndPrintAutocompleteScore(lines []string) {
	a := syntax.NewAnalyzer()

	var scores []int

	for i, line := range lines {
		score, err := a.ComputeAutocompleteScore(line)

		if err != nil {
			app.log.Fatalf(fmt.Sprintf("Encountered error while computing autocomplete score for line %d: %s\n", i+1, err.Error()))
			return
		}

		if score > 0 {
			scores = append(scores, score)
		}
	}

	if len(scores) == 0 {
		fmt.Printf("No line could be autocompleted.")
	}

	sort.Ints(scores)

	fmt.Printf("Middle autocomplete score: %d\n", scores[len(scores)/2])
}

func main() {
	var navSubsystemFile = flag.String("file", "input.txt", "Navigation subsystem file.")
	flag.Parse()

	app := application{log: log.Default()}

	lines, err := app.readLinesFromFile(*navSubsystemFile)

	if err != nil {
		app.log.Fatalf("Encountered error during navigation subsystem file parsing (%s).", err.Error())
	}

	app.computeAndPrintSyntaxErrorScore(lines)
	app.computeAndPrintAutocompleteScore(lines)
}
