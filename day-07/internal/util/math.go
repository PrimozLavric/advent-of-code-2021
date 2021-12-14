package util

import "sort"

// ComputeMedian computes median of the values in the provided slice.
func ComputeMedian(values []int) int {
	valuesCpy := make([]int, len(values))
	copy(valuesCpy, values)

	sort.Ints(valuesCpy)
	medianIdx := len(valuesCpy) / 2

	if medianIdx%2 != 0 {
		return valuesCpy[medianIdx]
	}

	return (valuesCpy[medianIdx-1] + valuesCpy[medianIdx]) / 2
}

// ComputeMean computes mean of the values in the provided slice.
func ComputeMean(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum / len(values)
}
