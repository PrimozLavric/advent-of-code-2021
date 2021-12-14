package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PrimozLavric/advent-of-code-2021/day-8/internal/decoder"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

// entry used to store encoded digits and encoded output from row of the encoded digits file.
type entry struct {
	encodedDigits []string
	encodedOutput []string
}

// parseEncodedDigitsFile reads entries from encoded digits file.
func (app *application) parseEncodedDigitsFile(filePath string) ([]*entry, error) {
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

	var entries []*entry

	for i := 1; fileScanner.Scan(); i++ {
		entryParts := strings.Split(fileScanner.Text(), "|")

		if len(entryParts) != 2 {
			return nil, errors.New(fmt.Sprintf("bad signal patterns file format, failed to parse row %d", i))
		}

		encodedDigits := strings.Fields(entryParts[0])
		encodedOutput := strings.Fields(entryParts[1])

		if err != nil {
			return nil, errors.New(fmt.Sprintf("bad signal patterns file format, failed to parse row %d (%s)", i, err.Error()))
		}

		entries = append(entries, &entry{encodedDigits, encodedOutput})
	}

	return entries, nil
}

func (app *application) decodeAndPrintPartOne(entries []*entry) {
	// Counters for 2, 3, 4, 7
	counter := 0

	// Only search for 2, 3, 4, 7
	for _, entry := range entries {
		for _, encodedDigit := range entry.encodedOutput {
			switch len(encodedDigit) {
			case 2:
				counter++
			case 3:
				counter++
			case 4:
				counter++
			case 7:
				counter++
			}
		}
	}

	fmt.Printf("Found %d (2, 3, 4, 7) digits combined\n", counter)
}

func (app *application) decodeAndPrintPartTwo(entries []*entry) {
	// Sum of all outputs (4-digit numbers)
	sum := 0

	// Only search for 2, 3, 4, 7
	for _, e := range entries {
		dd, err := decoder.NewDigitDecoder(e.encodedDigits)

		if err != nil {
			app.log.Fatalf("Decoding failed (%s)", err.Error())
		}

		decodedNumber := 0
		for _, encodedDigit := range e.encodedOutput {
			// Shift left (decimal).
			decodedNumber *= 10

			// Decode next digit
			digit, err := dd.Decode(encodedDigit)

			if err != nil {
				app.log.Fatalf("Decoding failed (%s)", err.Error())
			}

			decodedNumber += digit
		}

		sum += decodedNumber
	}

	fmt.Printf("Sum of all decoded outputs is: %d\n", sum)
}

func main() {
	var encodedDigitsFile = flag.String("file", "input.txt", "Encoded digits file.")
	flag.Parse()

	app := application{log: log.Default()}

	entries, err := app.parseEncodedDigitsFile(*encodedDigitsFile)

	if err != nil {
		app.log.Fatalf("Encountered error during signal patterns file parsing (%s).", err.Error())
	}

	app.decodeAndPrintPartOne(entries)
	app.decodeAndPrintPartTwo(entries)
}
