package util

import "errors"

// Heightmap contains grid of height values.
type Heightmap struct {
	heights [][]int
}

// NewHeightmap creates new map from the provided two-dimensional slice
func NewHeightmap(heights [][]int) (*Heightmap, error) {
	if len(heights) > 0 {
		rowLen := len(heights[0])
		for _, row := range heights {
			if len(row) != rowLen {
				return nil, errors.New("not all rows are of same length")
			}
		}
	}

	hm := Heightmap{heights: heights}

	return &hm, nil
}

// Width retrieves width of height map.
func (hm *Heightmap) Width() int {
	return len(hm.heights)
}

// Height retrieves height of height map.
func (hm *Heightmap) Height() int {
	if len(hm.heights) == 0 {
		return 0
	}

	return len(hm.heights[0])
}

// Value retrieves value of the heightmap at the provided position.
func (hm *Heightmap) Value(position [2]int) int {
	return hm.heights[position[0]][position[1]]
}

// IsLocalMinimum checks if the provided position is local minimum.
func (hm *Heightmap) IsLocalMinimum(pos [2]int) bool {
	if pos[0]-1 >= 0 && hm.heights[pos[0]][pos[1]] >= hm.heights[pos[0]-1][pos[1]] {
		return false
	}

	if pos[0]+1 < hm.Width() && hm.heights[pos[0]][pos[1]] >= hm.heights[pos[0]+1][pos[1]] {
		return false
	}

	if pos[1]-1 >= 0 && hm.heights[pos[0]][pos[1]] >= hm.heights[pos[0]][pos[1]-1] {
		return false
	}

	if pos[1]+1 < hm.Height() && hm.heights[pos[0]][pos[1]] >= hm.heights[pos[0]][pos[1]+1] {
		return false
	}

	return true
}

// FindLocalMinimums finds all local minimums and returns their coordinates.
func (hm *Heightmap) FindLocalMinimums() [][2]int {
	var localMinimums [][2]int

	for i := 0; i < hm.Width(); i++ {
		for j := 0; j < hm.Height(); j++ {
			if !hm.IsLocalMinimum([2]int{i, j}) {
				continue
			}

			localMinimums = append(localMinimums, [2]int{i, j})
		}
	}

	return localMinimums
}

// FindBasinSize finds size of the basin originating from the given local minimum.
func (hm *Heightmap) FindBasinSize(locMin [2]int) int {
	// Visited cells set has (defined as x * width + y)
	visitedCells := make(map[int]bool)

	// Closure used to cache already visited cells.
	var findBasinSize func([2]int) int
	findBasinSize = func(pos [2]int) int {
		// Stop at cells larger than 9.
		if hm.Value(pos) >= 9 {
			return 0
		}

		// Check if cell was already visited
		if visitedCells[pos[0]*hm.Width()+pos[1]] {
			return 0
		}

		// Mark cell as visited.
		visitedCells[pos[0]*hm.Width()+pos[1]] = true

		basinSize := 1

		// Continue searching the neighbours.
		if pos[0]-1 >= 0 && hm.heights[pos[0]][pos[1]] <= hm.heights[pos[0]-1][pos[1]] {
			basinSize += findBasinSize([2]int{pos[0] - 1, pos[1]})
		}

		if pos[0]+1 < hm.Width() && hm.heights[pos[0]][pos[1]] <= hm.heights[pos[0]+1][pos[1]] {
			basinSize += findBasinSize([2]int{pos[0] + 1, pos[1]})
		}

		if pos[1]-1 >= 0 && hm.heights[pos[0]][pos[1]] <= hm.heights[pos[0]][pos[1]-1] {
			basinSize += findBasinSize([2]int{pos[0], pos[1] - 1})
		}

		if pos[1]+1 < hm.Height() && hm.heights[pos[0]][pos[1]] <= hm.heights[pos[0]][pos[1]+1] {
			basinSize += findBasinSize([2]int{pos[0], pos[1] + 1})
		}

		return basinSize
	}

	return findBasinSize(locMin)
}
