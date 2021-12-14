package origami

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

// Coordinate 2D coordinate with X and Y axis.
type Coordinate struct {
	X int
	Y int
}

// Axis enum
type Axis int

const (
	AxisX Axis = 1
	AxisY Axis = 2
)

// FoldInstruction specifies how to execute a fold.
type FoldInstruction struct {
	Ax       Axis
	Position int
}

// Fold execute fold instruction on the given positions.
func Fold(dotPositions mapset.Set, fi FoldInstruction) mapset.Set {
	newDotPositions := mapset.NewSet()

	switch fi.Ax {
	case AxisX:
		// Fold over axis X.
		for entry := range dotPositions.Iterator().C {
			pos := entry.(Coordinate)

			if pos.X <= fi.Position {
				newDotPositions.Add(pos)
			} else {
				pos.X = 2*fi.Position - pos.X
				newDotPositions.Add(pos)
			}
		}
	case AxisY:
		// Fold over axis Y.
		for entry := range dotPositions.Iterator().C {
			pos := entry.(Coordinate)

			if pos.Y <= fi.Position {
				newDotPositions.Add(pos)
			} else {
				pos.Y = 2*fi.Position - pos.Y
				newDotPositions.Add(pos)
			}
		}
	}

	return newDotPositions
}

// PrintDots prints dot positions to stdout.
func PrintDots(dotPositions mapset.Set) {
	maxX, maxY := 0, 0

	// Find printing boundaries.
	for entry := range dotPositions.Iterator().C {
		pos := entry.(Coordinate)

		if pos.X > maxX {
			maxX = pos.X
		}

		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	// Print dots.
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if dotPositions.Contains(Coordinate{X: x, Y: y}) {
				fmt.Printf("O")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}
