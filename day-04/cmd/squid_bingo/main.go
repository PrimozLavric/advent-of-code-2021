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

	"github.com/PrimozLavric/advent-of-code-2021/day-4/internal/bingo"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

func (app *application) parseBingoFile(filePath string) ([]int, []*bingo.Board, error) {
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

	// Read bingo sequence.
	var bingoSequence []int

	if !fileScanner.Scan() {
		return nil, nil, errors.New("bad file format, file contains no data")
	}

	for i, strEntry := range strings.Split(fileScanner.Text(), ",") {
		entry, err := strconv.Atoi(strEntry)

		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("could not parse %d-th element of bingo sequence (%s)", i, err.Error()))
		}

		bingoSequence = append(bingoSequence, entry)
	}

	var bingoBoards []*bingo.Board
	var boardValues [bingo.BoardSize][bingo.BoardSize]int

	rowIdx := 0
	for fileScanner.Scan() {
		rowValues := strings.Fields(fileScanner.Text())

		if len(rowValues) == 0 {
			if rowIdx == bingo.BoardSize {
				bingoBoards = append(bingoBoards, bingo.NewBoard(boardValues))
				rowIdx = 0
			}

			if rowIdx != bingo.BoardSize && rowIdx != 0 {
				return nil, nil, errors.New("bad file format, bad bingo board format")
			}

			rowIdx = 0
			continue
		}

		if len(rowValues) != bingo.BoardSize {
			return nil, nil, errors.New(fmt.Sprintf("bad file format, invalid column count %d, expected %d", len(rowValues), bingo.BoardSize))
		}

		// Populate row
		for i := 0; i < bingo.BoardSize; i++ {
			value, err := strconv.Atoi(rowValues[i])

			if err != nil {
				return nil, nil, errors.New(fmt.Sprintf("bad file format, bad bingo board format (%s)", err.Error()))
			}

			boardValues[rowIdx][i] = value
		}

		rowIdx++
	}

	return bingoSequence, bingoBoards, nil
}

func (app *application) findAndPrintWinningBoard(bingoSequence []int, bingoBoards []*bingo.Board) {
	for _, value := range bingoSequence {
		for i, board := range bingoBoards {
			won, err := board.MarkValue(value)

			if err != nil {
				app.log.Panicf("Encountered error while playing bingo (%s).", err.Error())
				return
			}

			if won {
				fmt.Printf("Board %d won with score %d\n", i+1, board.Score())
				return
			}
		}
	}

	fmt.Println("No board won.")
}

func (app *application) findAndPrintBoardThatWinsLast(bingoSequence []int, bingoBoards []*bingo.Board) {
	boardsWonCount := 0

	for _, value := range bingoSequence {
		for i, board := range bingoBoards {
			if board.Won() {
				continue
			}

			won, err := board.MarkValue(value)

			if err != nil {
				app.log.Panicf("Encountered error while playing bingo (%s).", err.Error())
				return
			}

			if won {
				boardsWonCount++
				if boardsWonCount == len(bingoBoards) {
					fmt.Printf("Board %d wins last with score %d", i+1, board.Score())
					return
				}
			}
		}
	}

	fmt.Println("Not every board won.")
}

func main() {
	var bingoFile = flag.String("file", "input.txt", "File containing bingo data.")
	flag.Parse()

	app := application{log: log.Default()}

	bingoSequence, bingoBoards, err := app.parseBingoFile(*bingoFile)

	if err != nil {
		app.log.Fatalf("Encountered error during bingo file parsing (%s).", err.Error())
	}

	app.findAndPrintWinningBoard(bingoSequence, bingoBoards)

	// Reset boards.
	for _, board := range bingoBoards {
		board.Reset()
	}

	app.findAndPrintBoardThatWinsLast(bingoSequence, bingoBoards)
}
