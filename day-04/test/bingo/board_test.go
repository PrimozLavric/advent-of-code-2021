package test

import (
	"testing"

	"github.com/PrimozLavric/advent-of-code-2021/day-4/internal/bingo"
)

func TestBingoBoard(t *testing.T) {
	var boardValues [bingo.BoardSize][bingo.BoardSize]int

	for i := 0; i < bingo.BoardSize; i++ {
		for j := 0; j < bingo.BoardSize; j++ {
			boardValues[i][j] = i*bingo.BoardSize + j
		}
	}

	board := bingo.NewBoard(boardValues)

	assertNotWon := func() {
		if board.Won() {
			t.Fatalf("board marked as won, but winning condition was not yet met")
		}

		if board.Score() != 0 {
			t.Fatalf("board has not won yet, but score is not 0, actual score %d", board.Score())
		}
	}

	markExpectNotWon := func(value int) {
		won, err := board.MarkValue(value)

		// No error should occur.
		if err != nil {
			t.Fatalf("encountered error (%s) when marking the board", err.Error())
		}

		// Winning condition not expected.
		if won {
			t.Fatalf("board marked as won, but winning condition was not yet met")
		}

		// Validate getters.
		assertNotWon()
	}

	markExpectWon := func(value int) {
		won, err := board.MarkValue(value)

		// No error should occur.
		if err != nil {
			t.Fatalf("encountered error (%s) when marking the board", err.Error())
		}

		// Winning condition not expected.
		if !won || !board.Won() {
			t.Fatalf("winning condition should be met, but board is not marked as won")
		}
	}

	assertNotWon()

	// Mark diagonal
	markExpectNotWon(0)
	markExpectNotWon(6)
	markExpectNotWon(12)
	markExpectNotWon(18)
	markExpectNotWon(24)

	// Mark third column
	markExpectNotWon(2)
	markExpectNotWon(7)
	markExpectNotWon(17)
	markExpectWon(22)

	expectedScore := 192 * 22
	if board.Score() != expectedScore {
		t.Fatalf("expected score %d, actual score %d", expectedScore, board.Score())
	}

	// Try to mark another value. Expecting error.
	_, err := board.MarkValue(3)
	if err == nil {
		t.Fatalf("expected error, but none occured")
	}
}
