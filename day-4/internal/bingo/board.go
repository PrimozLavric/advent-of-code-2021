package bingo

import "errors"

const BoardSize = 5

type field struct {
	value  int
	marked bool
}

type Board struct {
	grid  [BoardSize][BoardSize]field
	won   bool
	score int
}

func NewBoard(values [BoardSize][BoardSize]int) *Board {
	b := Board{}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			b.grid[i][j].value = values[i][j]
		}
	}

	return &b
}

func (b *Board) MarkValue(value int) (bool, error) {
	if b.won {
		return true, errors.New("cannot mark value, because board already won")
	}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			field := &b.grid[i][j]
			if field.value == value {
				if field.marked {
					// Value already marked. Nothing to do.
					return false, nil
				}

				field.marked = true

				if b.checkWinningCondition(i, j) {
					b.won = true
					b.score = b.computeScore(value)
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func (b *Board) Reset() {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			b.grid[i][j].marked = false
		}
	}
	b.won = false
	b.score = 0
}

func (b *Board) Won() bool {
	return b.won
}

func (b *Board) Score() int {
	if !b.won {
		return 0
	}

	return b.score
}

func (b *Board) computeScore(lastValue int) int {
	score := 0

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			field := &b.grid[i][j]
			if !field.marked {
				score += field.value
			}
		}
	}

	return lastValue * score
}

func (b *Board) checkWinningCondition(row int, column int) bool {
	// Check rows.
	allMarked := true

	for i := 0; i < BoardSize; i++ {
		if !b.grid[row][i].marked {
			allMarked = false
			break
		}
	}

	// If all entries in a row are marked, winning condition in met.
	if allMarked {
		return true
	}

	// Check columns.
	allMarked = true

	for i := 0; i < BoardSize; i++ {
		if !b.grid[i][column].marked {
			allMarked = false
			break
		}
	}

	return allMarked
}
