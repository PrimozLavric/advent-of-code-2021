package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"day-12/internal/cave"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// readCaveConnections reads cave connection mapping from the provided file.
func (app *application) readCaveConnections(filePath string) (map[string][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	// Defer close the file.
	defer func() {
		err = file.Close()

		if err != nil {
			app.log.Printf("Failed to close file: %s\n", filePath)
		}
	}()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	connections := make(map[string][]string)

	for rowNum := 1; fileScanner.Scan(); rowNum++ {
		connectedCaves := strings.Split(fileScanner.Text(), "-")

		if len(connectedCaves) != 2 || len(connectedCaves[0]) == 0 || len(connectedCaves[1]) == 0 {
			return nil, errors.New(fmt.Sprintf("failed to parse row %d in energy levels file (%s)", rowNum, err.Error()))
		}

		connections[connectedCaves[0]] = append(connections[connectedCaves[0]], connectedCaves[1])
	}

	return connections, nil
}

func main() {
	var caveConnectionsFile = flag.String("file", "input.txt", "Read cave connections file.")
	flag.Parse()

	app := application{log: log.Default()}

	connections, err := app.readCaveConnections(*caveConnectionsFile)

	if err != nil {
		app.log.Fatalf("Encountered error during cave connections file parsing (%s).", err.Error())
	}

	caveGraph, err := cave.NewGraph(connections)

	if err != nil {
		app.log.Fatalf("Encountered error during cave graph creation (%s).", err.Error())
	}

	fmt.Printf("Found %d paths when visiting small caves only once.\n", caveGraph.FindNumberOfPathsToEnd(false))
	fmt.Printf("Found %d paths when visiting one small cave twice.\n", caveGraph.FindNumberOfPathsToEnd(true))
}
