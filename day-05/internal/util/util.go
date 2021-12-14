package util

import "strconv"

// SliceAtoi converts slice of strings to slice of integers.
func SliceAtoi(strSlice []string) ([]int, error) {
	intSlice := make([]int, 0, len(strSlice))

	for _, a := range strSlice {
		i, err := strconv.Atoi(a)
		if err != nil {
			return nil, err
		}

		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}

// AbsDiff computes absolute difference between values a and b.
func AbsDiff(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

// Lerp performs linear interpolation between values a and b using interpolant t
func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}
