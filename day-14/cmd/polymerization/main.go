package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/PrimozLavric/advent-of-code-2021/day-14/internal/polymer"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// readPolymerInstructionsFile reads polymer instructions file.
func (app *application) readPolymerInstructionsFile(filePath string) (*polymer.Template, map[string]rune, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, nil, err
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

	if !fileScanner.Scan() {
		return nil, nil, errors.New(fmt.Sprintf("provided file is empty"))
	}

	polymerTemplate := polymer.NewTemplate(fileScanner.Text())

	// Skip empty row.
	fileScanner.Scan()

	foldRegex := regexp.MustCompile("^([a-zA-Z]+) -> ([a-zA-Z])$")

	pairInsertionRules := make(map[string]rune)

	for rowNum := 3; fileScanner.Scan(); rowNum++ {
		foldStr := foldRegex.FindStringSubmatch(fileScanner.Text())

		if len(foldStr) != 3 {
			return nil, nil, errors.New(fmt.Sprintf("failed to parse row %d in polymer instructions file", rowNum))
		}

		pairInsertionRules[foldStr[1]] = rune(foldStr[2][0])
	}

	return polymerTemplate, pairInsertionRules, nil
}

func pairInsertionStep(polymerTemplate string, pairInsertionMap map[string]string) string {
	if len(polymerTemplate) == 0 {
		return ""
	}

	var buffer bytes.Buffer

	buffer.WriteByte(polymerTemplate[0])

	for i := 1; i < len(polymerTemplate); i++ {
		if char, ok := pairInsertionMap[polymerTemplate[i-1:i+1]]; ok {
			buffer.WriteString(char)
		}

		buffer.WriteByte(polymerTemplate[i])
	}

	return buffer.String()
}

func main() {
	var instructionsFile = flag.String("file", "input.txt", "Polymer instructions file.")
	flag.Parse()

	app := application{log: log.Default()}

	polymerTemplate, pairInsertionRules, err := app.readPolymerInstructionsFile(*instructionsFile)

	if err != nil {
		app.log.Fatalf("Encountered error during transparent origami instructions file parsing (%s).", err.Error())
	}

	for i := 0; i < 10; i++ {
		polymerTemplate.PerformPairInsertion(pairInsertionRules)
	}

	fmt.Printf("Polymer template score after 10 insertion step: %d\n", polymerTemplate.Score())

	for i := 0; i < 30; i++ {
		polymerTemplate.PerformPairInsertion(pairInsertionRules)
	}

	fmt.Printf("Polymer template after 40 insertion step: %d\n", polymerTemplate.Score())
}
