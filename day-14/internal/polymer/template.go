package polymer

import (
	"math"
)

// Template contains polymer template info.
type Template struct {
	characterCounts map[rune]uint64
	pairsCounts     map[string]uint64
}

// NewTemplate creates new Template from the provided polymer template string.
func NewTemplate(strTemplate string) *Template {
	characterCounts := make(map[rune]uint64)
	pairsCounts := make(map[string]uint64)

	if len(strTemplate) == 0 {
		return &Template{characterCounts: characterCounts, pairsCounts: pairsCounts}
	}

	characterCounts[rune(strTemplate[0])]++

	for i := 1; i < len(strTemplate); i++ {
		characterCounts[rune(strTemplate[i])]++
		pairsCounts[strTemplate[i-1:i+1]]++
	}

	return &Template{characterCounts: characterCounts, pairsCounts: pairsCounts}
}

// PerformPairInsertion inserts new characters based on pair insertion rules.
func (pt *Template) PerformPairInsertion(pairInsertionRules map[string]rune) {
	newPairsCounts := make(map[string]uint64)

	for pair, count := range pt.pairsCounts {
		if char, ok := pairInsertionRules[pair]; ok {
			pt.characterCounts[char] += count
			newPairsCounts[string(pair[0])+string(char)] += count
			newPairsCounts[string(char)+string(pair[1])] += count
		} else {
			newPairsCounts[pair] += count
		}
	}

	pt.pairsCounts = newPairsCounts
}

// Score computes Template score defined as score = max char count - min char count.
func (pt *Template) Score() uint64 {
	minCount := uint64(math.MaxUint64)
	maxCount := uint64(0)

	for _, count := range pt.characterCounts {
		if count < minCount {
			minCount = count
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount - minCount
}
