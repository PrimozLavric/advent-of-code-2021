package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/PrimozLavric/advent-of-code-2021/day-9/internal/util"
)

// An application contains application wide data such as Logger.
type application struct {
	log *log.Logger
}

func (app *application) parseHeightmapFile(filePath string) (*util.Heightmap, error) {
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

	var heights [][]int

	for i := 1; fileScanner.Scan(); i++ {
		row := strings.Split(fileScanner.Text(), "")

		rowHeights, err := util.SliceAtoi(row)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("bad signal patterns file format, failed to parse row %d (%s)", i, err.Error()))
		}

		heights = append(heights, rowHeights)
	}

	return util.NewHeightmap(heights)
}

func (app *application) findAndPrintSumOfRiskLevels(hm *util.Heightmap) {
	localMinimums := hm.FindLocalMinimums()

	riskLevelSum := 0

	for _, locMin := range localMinimums {
		riskLevelSum += hm.Value(locMin) + 1
	}

	fmt.Printf("Sum of risk levels: %d \n", riskLevelSum)
}

func (app *application) findAndPrintThreeLargesBasinSizesMultiplied(hm *util.Heightmap) {
	localMinimums := hm.FindLocalMinimums()

	var basinSizes []int

	for _, locMin := range localMinimums {
		basinSizes = append(basinSizes, hm.FindBasinSize(locMin))
	}

	sort.Ints(basinSizes)

	if len(basinSizes) < 3 {
		app.log.Fatal("Found less than 3 basins.")
	}

	largestBasinMultiplied := 1

	for i := len(basinSizes) - 3; i < len(basinSizes); i++ {
		largestBasinMultiplied *= basinSizes[i]
	}

	fmt.Printf("Largest three basin sizes multiplied: %d \n", largestBasinMultiplied)
}

func main() {
	var heightmapFile = flag.String("file", "input.txt", "Heightmap file.")
	flag.Parse()

	app := application{log: log.Default()}

	hm, err := app.parseHeightmapFile(*heightmapFile)

	if err != nil {
		app.log.Fatalf("Encountered error during signal patterns file parsing (%s).", err.Error())
	}

	app.findAndPrintSumOfRiskLevels(hm)
	app.findAndPrintThreeLargesBasinSizesMultiplied(hm)
}
