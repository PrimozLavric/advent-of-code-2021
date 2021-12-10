package syntax

import (
	"errors"
	"fmt"
)

// Analyzer is used to compute syntax error scores and autocomplete scores.
type Analyzer struct {
	symbolPairs       map[rune]rune
	symbolPenalties   map[rune]int
	autocompleteScore map[rune]int
}

// NewAnalyzer creates Analyzer and initializes matching symbol pairs, penalties and autocomplete scores.
func NewAnalyzer() *Analyzer {
	v := Analyzer{
		map[rune]rune{
			'(': ')',
			'[': ']',
			'{': '}',
			'<': '>',
		},
		map[rune]int{
			')': 3,
			']': 57,
			'}': 1197,
			'>': 25137,
		},
		map[rune]int{
			')': 1,
			']': 2,
			'}': 3,
			'>': 4,
		},
	}

	return &v
}

// ComputeSyntaxErrorScore computes syntax error score using the incorrect symbol penalty value.
func (v *Analyzer) ComputeSyntaxErrorScore(line string) (int, error) {
	var openSymbolStack []rune

	for _, ch := range line {
		// If it is an opening symbol add it to stack.
		if _, ok := v.symbolPairs[ch]; ok {
			openSymbolStack = append(openSymbolStack, ch)
			continue
		}

		if len(openSymbolStack) == 0 {
			return 0, errors.New(fmt.Sprintf("expected an opening sybmol, but got '%c'", ch))
		}

		// Pop opening symbol from the top of the stack.
		openSymbol := openSymbolStack[len(openSymbolStack)-1]
		openSymbolStack = openSymbolStack[:len(openSymbolStack)-1]

		match, _ := v.symbolPairs[openSymbol]

		if ch != match {
			penalty, ok := v.symbolPenalties[ch]

			if !ok {
				return 0, errors.New(fmt.Sprintf("encountered invalid character '%c'", ch))
			}

			return penalty, nil
		}
	}

	return 0, nil
}

// ComputeAutocompleteScore computes autocomplete score based on the missing symbols and symbol scores.
func (v *Analyzer) ComputeAutocompleteScore(line string) (int, error) {
	var openSymbolStack []rune

	for _, ch := range line {
		// If it is an opening symbol add it to stack.
		if _, ok := v.symbolPairs[ch]; ok {
			openSymbolStack = append(openSymbolStack, ch)
			continue
		}

		if len(openSymbolStack) == 0 {
			return 0, errors.New(fmt.Sprintf("expected an opening sybmol, but got '%c'", ch))
		}

		// Pop opening symbol from the top of the stack.
		openSymbol := openSymbolStack[len(openSymbolStack)-1]
		openSymbolStack = openSymbolStack[:len(openSymbolStack)-1]

		match, _ := v.symbolPairs[openSymbol]

		// Detected syntax error. This line cannot be autocompleted.
		if ch != match {
			return 0, nil
		}
	}

	// Compute autocomplete score.
	score := 0

	for i := len(openSymbolStack) - 1; i >= 0; i-- {
		score *= 5
		score += v.autocompleteScore[v.symbolPairs[openSymbolStack[i]]]
	}

	return score, nil
}
