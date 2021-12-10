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
