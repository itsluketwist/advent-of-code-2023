package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 14

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
	var grid [][]string
	for _, line := range data {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func calculateLoad(grid [][]string) int {
	total := 0
	for col := 0; col < len(grid[0]); col++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][col] == "O" {
				total += (len(grid) - j)
			}
		}
	}
	return total
}

func doFullCycle(grid [][]string) {
	// tilt north
	for col := 0; col < len(grid[0]); col++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][col] == "." {
				for k := j + 1; k < len(grid); k++ {
					if grid[k][col] == "O" {
						// stone rolls
						grid[j][col] = "O"
						grid[k][col] = "."
						break
					} else if grid[k][col] == "#" {
						break
					}
				}
			}
		}
	}

	// tilt west
	for row := 0; row < len(grid); row++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[row][j] == "." {
				for k := j + 1; k < len(grid[0]); k++ {
					if grid[row][k] == "O" {
						// stone rolls
						grid[row][j] = "O"
						grid[row][k] = "."
						break
					} else if grid[row][k] == "#" {
						break
					}
				}
			}
		}
	}

	// tilt south
	for col := 0; col < len(grid[0]); col++ {
		for j := len(grid) - 1; j >= 0; j-- {
			if grid[j][col] == "." {
				for k := j - 1; k >= 0; k-- {
					if grid[k][col] == "O" {
						// stone rolls
						grid[j][col] = "O"
						grid[k][col] = "."
						break
					} else if grid[k][col] == "#" {
						break
					}
				}
			}
		}
	}

	// tilt east
	for row := 0; row < len(grid); row++ {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if grid[row][j] == "." {
				for k := j - 1; k >= 0; k-- {
					if grid[row][k] == "O" {
						// stone rolls
						grid[row][j] = "O"
						grid[row][k] = "."
						break
					} else if grid[row][k] == "#" {
						break
					}
				}
			}
		}
	}
}

func SolvePartOne(data []string) int {
	grid := parseData(data)

	// tilt north
	for col := 0; col < len(grid[0]); col++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][col] == "." {
				for k := j + 1; k < len(grid); k++ {
					if grid[k][col] == "O" {
						// stone rolls
						grid[j][col] = "O"
						grid[k][col] = "."
						break
					} else if grid[k][col] == "#" {
						break
					}
				}
			}
		}
	}

	return calculateLoad(grid)
}

func SolvePartTwo(data []string) int {
	grid := parseData(data)

	totalLoops := 1000000000
	initialLoops := 153

	// loop enough for pattern to repeat
	for n := 0; n < initialLoops; n++ {
		doFullCycle(grid)
	}

	// calculate the length of the pattern
	patternLen := 0
	loadAfterInitial := calculateLoad(grid)
	for n := 0; n < initialLoops; n++ {
		doFullCycle(grid)
		load := calculateLoad(grid)
		if load == loadAfterInitial {
			patternLen = n + 1
			break
		}
	}

	// loop to the right point of the pattern
	extraLoop := (totalLoops - initialLoops - patternLen) % patternLen
	for n := 0; n < extraLoop; n++ {
		doFullCycle(grid)
	}

	return calculateLoad(grid)
}
