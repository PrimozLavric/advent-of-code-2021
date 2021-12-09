package util

import "fmt"

// IsIthBitSet checks if the i-th bit of the given value is set.
func IsIthBitSet(value uint32, i uint8) bool {
	const bitCount = 32
	if i >= bitCount {
		panic(fmt.Sprintf("bit index %d out of range, must be less than %d", i, bitCount))
	}

	return value&(1<<i) != 0
}

// CountSetIthBits counts numbers in the provided slice in which the i-th bit is set.
func CountSetIthBits(data []uint32, i uint8) uint32 {
	const bitCount = 32
	if i >= bitCount {
		panic(fmt.Sprintf("bit index %d out of range, must be less than %d", i, bitCount))
	}

	bitCounter := uint32(0)

	for _, entry := range data {
		if IsIthBitSet(entry, i) {
			bitCounter++
		}
	}

	return bitCounter
}

// FindMostCommonIthBit finds the most common i-th bit in the numbers in the provided slice. If zero is more common it
// returns false otherwise true.
func FindMostCommonIthBit(data []uint32, i uint8) bool {
	count := CountSetIthBits(data, i)

	return count*2 >= uint32(len(data))
}

// FilterDataBasedOnMostCommonBitValue produces filtered copy of the given slice in which only values that have same
// i-th bit as the most common i-th bit in all values are kept.
func FilterDataBasedOnMostCommonBitValue(data []uint32, i uint8) []uint32 {
	mostCommonBit := FindMostCommonIthBit(data, i)

	var filteredData []uint32

	for _, entry := range data {
		if IsIthBitSet(entry, i) == mostCommonBit {
			filteredData = append(filteredData, entry)
		}
	}

	return filteredData
}

// FilterDataBasedOnLeastCommonBitValue produces filtered copy of the given slice in which only values that have same
// i-th bit as the least common i-th bit in all values are kept.
func FilterDataBasedOnLeastCommonBitValue(data []uint32, i uint8) []uint32 {
	leastCommonBit := !FindMostCommonIthBit(data, i)

	var filteredData []uint32

	for _, entry := range data {
		if IsIthBitSet(entry, i) == leastCommonBit {
			filteredData = append(filteredData, entry)
		}
	}

	return filteredData
}
