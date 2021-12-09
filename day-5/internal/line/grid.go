package line

import (
	"math"

	"github.com/PrimozLavric/advent-of-code-2021/day-5/internal/util"
)

// Grid that is used to find intersections.
type Grid [][]int

// MakeGrid creates new grid with the provided number of rows and columns filled with zeros.
func MakeGrid(rows, columns int) Grid {
	grid := make([][]int, rows)

	for i := 0; i < rows; i++ {
		grid[i] = make([]int, columns)
	}

	return grid
}

// ApplyLine increments grid cells' value if they intersect with the provided line.
func (grid Grid) ApplyLine(line *Line) {
	xDiff := util.AbsDiff(line.A.X, line.B.X)
	yDiff := util.AbsDiff(line.A.Y, line.B.Y)

	maxDiff := math.Max(float64(xDiff), float64(yDiff))

	for i := 0; i <= int(maxDiff); i++ {
		t := float64(i) / maxDiff
		x := math.Round(util.Lerp(float64(line.A.X), float64(line.B.X), t))
		y := math.Round(util.Lerp(float64(line.A.Y), float64(line.B.Y), t))

		grid[int(x)][int(y)] += 1
	}
}

// CountIntersections counts number of cells at which the applied lines intersected.
func (grid Grid) CountIntersections() int {
	isects := 0

	for _, row := range grid {
		for _, value := range row {
			if value > 1 {
				isects++
			}
		}
	}

	return isects
}
