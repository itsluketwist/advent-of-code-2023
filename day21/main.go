package main

import (
	"flag"
	"fmt"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 21

func main() {
	part := flag.Int("part", 0, "Which parts to try?")
	try := flag.Int("try", 0, "Whether to try the real input?")
	flag.Parse()

	fmt.Println("Running day", Day, "( part:", *part, ", try:", *try, ")")

	exampleOne, _ := utils.ReadFileToArray(Day, "example1", false)
	input, _ := utils.ReadFileToArray(Day, "input", false)

	if *part == 0 || *part == 1 {
		solutionOneExample := SolvePartOne(exampleOne, 6)
		fmt.Println("Solution to part 1 (example):", solutionOneExample)

		if *try == 1 {
			SolutionOneInput := SolvePartOne(input, 64)
			fmt.Println("Solution to part 1 (input):", SolutionOneInput)
		}
	}

	if *part == 0 || *part == 2 {
		if *try == 1 {
			SolutionTwoInput := SolvePartTwo(input)
			fmt.Println("Solution to part 2 (input):", SolutionTwoInput)
		}
	}
}

func parseData(data []string) [][]string {
	var plot [][]string
	for _, line := range data {
		var row []string
		for _, char := range line {
			if char == 'S' {
				row = append(row, ".")
			} else {
				row = append(row, string(char))
			}
		}
		plot = append(plot, row)
	}
	return plot
}

func copyPlot(plot [][]string) [][]string {
	var newPlot [][]string
	for _, row := range plot {
		var newRow []string
		for _, cell := range row {
			newRow = append(newRow, cell)
		}
		newPlot = append(newPlot, newRow)
	}
	return newPlot
}

func updateAround(plot [][]string, row int, col int) {
	for _, n := range []int{-1, 1} {
		if 0 <= row+n && row+n < len(plot) {
			if plot[row+n][col] == "." {
				plot[row+n][col] = "O"
			}
		}

		if 0 <= col+n && col+n < len(plot[0]) {
			if plot[row][col+n] == "." {
				plot[row][col+n] = "O"
			}
		}
	}
}

func Solve(data []string, loops int, extra int) int {
	plot := parseData(data)

	// expand the plot if required
	if extra > 0 {
		gridScale := 1 + 2*extra
		var newPlot [][]string
		for i := 0; i < gridScale; i++ {
			for _, row := range plot {
				var newRow []string
				for j := 0; j < gridScale; j++ {
					for _, cell := range row {
						newRow = append(newRow, cell)
					}
				}
				newPlot = append(newPlot, newRow)
			}
		}
		plot = newPlot
	}

	baseEmpty := copyPlot(plot)                     // get a base empty grid
	plot[(len(plot)-1)/2][(len(plot[0])-1)/2] = "O" // set start location

	for i := 0; i < loops; i++ {
		// make a new grid, fill it in from the old one
		next := copyPlot(baseEmpty)

		for i, row := range plot {
			for j, cell := range row {
				if cell == "O" {
					// update surrounding cells in next
					updateAround(next, i, j)
				}
			}
		}
		plot = next
	}

	total := 0
	for _, row := range plot {
		for _, cell := range row {
			if cell == "O" {
				total++
			}
		}
	}
	return total
}

func SolvePartOne(data []string, steps int) int {
	return Solve(data, steps, 0)
}

func SolvePartTwo(data []string) int {
	takeSteps := 26501365
	initLength := 65  // to get through first grid
	gridLength := 131 // to get through other grids

	// infinite garden expands as a quadratic for steps covered wrt total grids, need the first few values
	for i := 0; i < 3; i++ {
		fmt.Println("x =", i, "| y =", Solve(data, 65+i*131, i))
	}

	// use https://www.wolframalpha.com/input?i=quadratic+fit+calculator to solve
	// get: y = 14988*x^2 + 15067*x + 3778

	x := (takeSteps - initLength) / gridLength
	y := 14988*x*x + 15067*x + 3778
	return y
}
