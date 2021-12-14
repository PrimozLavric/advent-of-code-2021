package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PrimozLavric/advent-of-code-2021/day-13/internal/origami"
	mapset "github.com/deckarep/golang-set"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// readTransparentOrigamiInstructionsFile transparent origami instructions file and returns a set of dot positions and slice of fold instructions.
func (app *application) readTransparentOrigamiInstructionsFile(filePath string) (mapset.Set, []origami.FoldInstruction, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, nil, err
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

	rowNum := 1

	// Read dot positions
	dotsPositions := mapset.NewSet()

	for ; fileScanner.Scan(); rowNum++ {
		// Stop reading dot positions when empty line is read
		if len(strings.TrimSpace(fileScanner.Text())) == 0 {
			rowNum++
			break
		}

		strPos := strings.Split(fileScanner.Text(), ",")

		if len(strPos) != 2 {
			return nil, nil, errors.New(fmt.Sprintf("failed to parse row %d in transparent origami instructions file", rowNum))
		}

		var position origami.Coordinate

		position.X, err = strconv.Atoi(strPos[0])

		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("failed to parse row %d in transparent origami instructions file (%s)", rowNum, err.Error()))
		}

		position.Y, err = strconv.Atoi(strPos[1])

		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("failed to parse row %d in transparent origami instructions file (%s)", rowNum, err.Error()))
		}

		dotsPositions.Add(position)
	}

	// Read fold instructions.
	var folds []origami.FoldInstruction
	foldRegex := regexp.MustCompile("^fold along ([x|y])=([0-9]+)$")

	for ; fileScanner.Scan(); rowNum++ {
		foldStr := foldRegex.FindStringSubmatch(fileScanner.Text())

		if len(foldStr) != 3 {
			return nil, nil, errors.New(fmt.Sprintf("failed to parse row %d in transparent origami instructions file", rowNum))
		}

		var ax origami.Axis
		switch foldStr[1] {
		case "x":
			ax = origami.AxisX
		case "y":
			ax = origami.AxisY
		default:
			return nil, nil, errors.New(fmt.Sprintf("failed to parse row %d in transparent origami instructions file", rowNum))
		}

		position, err := strconv.Atoi(foldStr[2])

		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("failed to parse row %d in transparent origami instructions file (%s)", rowNum, err.Error()))
		}

		folds = append(folds, origami.FoldInstruction{Ax: ax, Position: position})
	}

	return dotsPositions, folds, nil
}

func main() {
	var instructionsFile = flag.String("file", "input.txt", "Transparent origami instructions file.")
	flag.Parse()

	app := application{log: log.Default()}

	dotPositions, folds, err := app.readTransparentOrigamiInstructionsFile(*instructionsFile)

	if err != nil {
		app.log.Fatalf("Encountered error during transparent origami instructions file parsing (%s).", err.Error())
	}

	// Execute first fold (part one of the exercise)
	dotPositions = origami.Fold(dotPositions, folds[0])
	fmt.Printf("Number of dots after 1 fold: %d\n\n", dotPositions.Cardinality())

	// Execute remaining folds.
	for _, f := range folds[1:] {
		dotPositions = origami.Fold(dotPositions, f)
	}

	origami.PrintDots(dotPositions)
}
