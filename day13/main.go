package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 13

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

func parseData(data []string) [][][]string {
	var grids [][][]string
	var subGrid [][]string
	for _, line := range data {
		if line == "" {
			grids = append(grids, subGrid)
			subGrid = [][]string{}
		} else {
			subGrid = append(subGrid, []string(strings.Split(line, "")))
		}
	}
	return grids
}

func Solve(data []string, errorsAllowed int) int {
	grids := parseData(data)
	total := 0

	for _, grid := range grids {
		gridRows := len(grid)
		gridCols := len(grid[0])

		// horizontal symmetry
		for i := 0; i < gridRows-1; i++ {
			errors := 0
			lower := i
			upper := i + 1
			for {
				for k := 0; k < gridCols; k++ {
					if grid[lower][k] != grid[upper][k] {
						errors++
					}
				}

				lower--
				upper++

				if lower < 0 || upper >= gridRows {
					break
				}
			}
			if errors == errorsAllowed {
				total += (i + 1) * 100
				break
			}
		}

		// vertical symmetry
		for j := 0; j < gridCols-1; j++ {
			errors := 0
			left := j
			right := j + 1
			for {
				for _, row := range grid {
					if row[left] != row[right] {
						errors++
					}
				}

				left--
				right++

				if left < 0 || right >= gridCols {
					break
				}
			}
			if errors == errorsAllowed {
				total += (j + 1)
				break
			}
		}

	}

	return total
}

func SolvePartOne(data []string) int {
	return Solve(data, 0)
}

func SolvePartTwo(data []string) int {
	return Solve(data, 1)
}
