package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	submarine2 "github.com/PrimozLavric/advent-of-code-2021/day-2/internal/submarine"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

type moveInstruction struct {
	dir      submarine2.Direction
	distance uint
}

func (app *application) readMoveInstructions(filePath string) ([]moveInstruction, error) {
	const rowFieldCount = 2

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

	var instructions []moveInstruction

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	rowIdx := 0
	for fileScanner.Scan() {
		// Tokenize row.
		entryFields := strings.Fields(fileScanner.Text())

		if len(entryFields) != rowFieldCount {
			return nil, errors.New(fmt.Sprintf("bad move instructions file (row %d has %d entries, but expected %d entries)", rowIdx, len(entryFields), rowFieldCount))
		}

		// Parse direction.
		dir, err := submarine2.MakeDirection(entryFields[0])

		if err != nil {
			return nil, errors.New(fmt.Sprintf("bad move instructions file (failed to parse row %d direction entry '%s')", rowIdx, entryFields[0]))
		}

		// Parse distance.
		distance, err := strconv.ParseUint(entryFields[1], 10, 32)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("bad move instructions file (failed to parse row %d distance entry '%s')", rowIdx, entryFields[1]))
		}

		instructions = append(instructions, moveInstruction{
			dir, uint(distance),
		})

		rowIdx++
	}

	return instructions, nil
}

func main() {
	var instructionsFile = flag.String("file", "input.txt", "File from which the translation data will be read.")
	flag.Parse()

	app := application{log: log.Default()}

	instructions, err := app.readMoveInstructions(*instructionsFile)

	if err != nil {
		app.log.Fatalf("Failed to parse move instructions file (%s)", err.Error())
		return
	}

	subPartOne := submarine2.Submarine{}
	subPartTwo := submarine2.Submarine{}

	for _, instuct := range instructions {
		err := subPartOne.MovePartOne(instuct.dir, instuct.distance)

		if err != nil {
			app.log.Fatalf("Failed to move exercise part one submarine (%s)", err.Error())
			return
		}

		err = subPartTwo.MovePartTwo(instuct.dir, instuct.distance)

		if err != nil {
			app.log.Fatalf("Failed to move exercise part two submarine (%s)", err.Error())
			return
		}
	}

	fmt.Printf("Exercise Part One:\n\tHorizontal position: %d\n\tVertical position: %d\n\tMultiplied positions %d\n", subPartOne.PositionX(), subPartOne.PositionY(), subPartOne.MultipliedPositions())
	fmt.Printf("Exercise Part Two:\n\tHorizontal position: %d\n\tVertical position: %d\n\tMultiplied positions %d\n", subPartTwo.PositionX(), subPartTwo.PositionY(), subPartTwo.MultipliedPositions())
}
