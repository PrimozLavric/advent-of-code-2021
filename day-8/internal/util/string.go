package util

import (
	"sort"
	"strings"
)

// sortRunes is a helper type for string sorting
type sortRunes []rune

// Less checks if rune on i-th place compares less than rune on j-th index.
func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap swaps runes on i-th and j-th index.
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Len counts number or runes.
func (s sortRunes) Len() int {
	return len(s)
}

// StringSort returns a character sorted copy of the provided string.
func StringSort(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// StringContainsAll checks if provided string str contains all characters of the chars string
func StringContainsAll(str string, chars string) bool {
	for _, char := range chars {
		if !strings.ContainsRune(str, char) {
			return false
		}
	}

	return true
}

// StringCountContainedChars counts number of characters in chars string that are contained in str
func StringCountContainedChars(str string, chars string) int {
	counter := 0
	for _, char := range chars {
		if strings.ContainsRune(str, char) {
			counter++
		}
	}

	return counter
}
