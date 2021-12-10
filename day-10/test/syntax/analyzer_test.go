package test

import (
	"testing"

	"day-10/internal/syntax"
)

func TestComputeAutocompleteScore(t *testing.T) {
	testCases := map[string]int{
		"[({(<(())[]>[[{[]{<()<>>": 288957,
		"[(()[<>])]({[<{<<[]>>(":   5566,
		"(((({<>}<{<{<>}{[]{[]{}":  1480781,
		"{<[[]]>}<{[{[{[]{()[[[]":  995444,
		"<{([{{}}[<[[[<>{}]]]>[]]": 294,
		"<{([])}>":                 0,
	}

	a := syntax.NewAnalyzer()

	for line, expectedScore := range testCases {
		score, err := a.ComputeAutocompleteScore(line)

		if err != nil {
			t.Errorf("encountered error while computing autocomplete score (%s)", err.Error())
		}

		if score != expectedScore {
			t.Errorf("computed autocomplete score %d for string '%s', actual %d", score, line, expectedScore)
		}
	}
}
