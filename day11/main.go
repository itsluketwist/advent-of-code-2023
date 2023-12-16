package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 11

func main() {
	part := flag.Int("part", 0, "Which parts to try?")
	try := flag.Int("try", 0, "Whether to try the real input?")
	flag.Parse()

	fmt.Println("Running day", Day, "( part:", *part, ", try:", *try, ")")

	exampleOne, _ := utils.ReadFileToArray(Day, "example1", false)
	exampleTwo, _ := utils.ReadFileToArray(Day, "example1", false)
	input, _ := utils.ReadFileToArray(Day, "input", false)

	if *part == 0 || *part == 1 {
		solutionOneExample := SolvePartOne(exampleOne)
		fmt.Println("Solution to part 1 (example):", solutionOneExample)

		if *try == 1 {
			SolutionOneInput := SolvePartOne(input)
			fmt.Println("Solution to part 1 (input):", SolutionOneInput)
		}
	}

	if *part == 0 || *part == 2 {
		SolutionTwoExample := SolvePartTwo(exampleTwo)
		fmt.Println("Solution to part 2 (example):", SolutionTwoExample)

		if *try == 1 {
			SolutionTwoInput := SolvePartTwo(input)
			fmt.Println("Solution to part 2 (input):", SolutionTwoInput)
		}
	}
}

func parseData(data []string) [][]string {
	var galaxy [][]string
	for _, line := range data {
		galaxy = append(galaxy, strings.Split(line, ""))
	}
	return galaxy
}

func Solve(data []string, expansionValue int) int {
	galaxy := parseData(data)

	var rowsEmpty []bool
	for i := 0; i < len(galaxy); i++ {
		rowsEmpty = append(rowsEmpty, true)
	}
	var colsEmpty []bool
	for i := 0; i < len(galaxy[0]); i++ {
		colsEmpty = append(colsEmpty, true)
	}

	// find panets, track empty rows and cols
	var planets [][]int
	for i, row := range galaxy {
		for j, space := range row {
			if space == "#" {
				planets = append(planets, []int{i, j})
				rowsEmpty[i] = false
				colsEmpty[j] = false
			}
		}
	}

	// calculate distances
	distance := 0
	for a := 0; a < len(planets); a++ {
		for b := a + 1; b < len(planets); b++ {
			rowDist := int(math.Abs(float64(planets[a][0] - planets[b][0])))
			rowStart := int(math.Min(float64(planets[a][0]), float64(planets[b][0])))
			for i := 0; i < rowDist; i++ {
				if rowsEmpty[rowStart+i] {
					distance += expansionValue
				} else {
					distance += 1
				}
			}

			colDist := int(math.Abs(float64(planets[a][1] - planets[b][1])))
			colStart := int(math.Min(float64(planets[a][1]), float64(planets[b][1])))
			for j := 0; j < colDist; j++ {
				if colsEmpty[colStart+j] {
					distance += expansionValue
				} else {
					distance += 1
				}
			}
		}
	}

	return distance
}

func SolvePartOne(data []string) int {
	return Solve(data, 2)
}

func SolvePartTwo(data []string) int {
	return Solve(data, 1000000)
}
