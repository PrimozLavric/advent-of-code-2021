package test

import (
	"testing"

	"day-12/internal/cave"
)

func TestCaveGraphFindAllPathsToEnd(t *testing.T) {
	testData := map[string][]string{
		"start": {"A", "b"},
		"A":     {"c", "b", "end"},
		"b":     {"d", "end"},
	}

	caveGraph, err := cave.NewGraph(testData)

	if err != nil {
		t.Errorf("Unexpected error occured (%s).", err.Error())
	}

	// Visit small nodes only once.
	pathCount := caveGraph.FindNumberOfPathsToEnd(false)

	if pathCount != 10 {
		t.Errorf("Expected %d paths, but found %d paths", 10, pathCount)
	}

	// Visit one small node twice.
	pathCount = caveGraph.FindNumberOfPathsToEnd(true)

	if pathCount != 36 {
		t.Errorf("Expected %d paths, but found %d paths", 36, pathCount)
	}
}
