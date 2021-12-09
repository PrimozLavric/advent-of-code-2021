package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/PrimozLavric/advent-of-code-2021/day-5/internal/line"
	"github.com/PrimozLavric/advent-of-code-2021/day-5/internal/util"
)

type application struct {
	log *log.Logger
}

func (app *application) parseLinesFile(filePath string) ([]*line.Line, error) {
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

	lineRegex := regexp.MustCompile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")

	var lines []*line.Line

	rowIdx := 0
	for fileScanner.Scan() {
		matches := lineRegex.FindStringSubmatch(fileScanner.Text())

		if matches == nil || len(matches) != 5 {
			return nil, errors.New(fmt.Sprintf("failed to parse line %d (%s)", rowIdx, fileScanner.Text()))
		}

		intEntries, err := util.SliceAtoi(matches[1:])

		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to parse line %d (%s)", rowIdx, fileScanner.Text()))
		}

		lines = append(lines, line.NewLine(intEntries[0], intEntries[1], intEntries[2], intEntries[3]))

		rowIdx++
	}

	return lines, nil
}

func main() {
	var linesFile = flag.String("file", "input.txt", "Hydrothermal vents lines file.")
	flag.Parse()

	app := application{log: log.Default()}

	lines, err := app.parseLinesFile(*linesFile)

	if err != nil {
		app.log.Fatalf("Encountered error during hydrothermal vent lines file parsing (%s).", err.Error())
	}

	maxX, maxY := line.FindMaxXY(lines)

	gridHV := line.MakeGrid(maxX+1, maxY+1)

	// Compute number of horizontal and vertical lines intersections.
	for _, l := range lines {
		if l.IsHorizontal() || l.IsVertical() {
			gridHV.ApplyLine(l)
		}
	}

	fmt.Printf("Number of horizontal and vertical lines intersections: %d\n", gridHV.CountIntersections())

	gridALL := line.MakeGrid(maxX+1, maxY+1)

	// Compute number of all intersections.
	for _, l := range lines {
		gridALL.ApplyLine(l)
	}

	fmt.Printf("Number of all lines intersections: %d\n", gridALL.CountIntersections())
}
