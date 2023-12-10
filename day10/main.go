package main

import (
	"flag"
	"fmt"
	"math"
	"slices"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 10

func main() {
	part := flag.Int("part", 0, "Which parts to try?")
	try := flag.Int("try", 0, "Whether to try the real input?")
	flag.Parse()

	fmt.Println("Running day", Day, "( part:", *part, ", try:", *try, ")")

	exampleOne, _ := utils.ReadFileToArray(Day, "example1", false)
	exampleTwo, _ := utils.ReadFileToArray(Day, "example2", false)
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

type Move struct {
	rowChange int
	colChange int
}

var directions = map[string]Move{
	"u": {-1, 0},
	"d": {+1, 0},
	"l": {0, -1},
	"r": {0, +1},
}

var pipeMoves = map[string]map[string]string{
	"|": {"u": "u", "d": "d"},
	"-": {"l": "l", "r": "r"},
	"F": {"l": "d", "u": "r"},
	"7": {"r": "d", "u": "l"},
	"L": {"l": "u", "d": "r"},
	"J": {"r": "u", "d": "l"},
}

func parseData(data []string) ([][]string, []int) {
	var pipes [][]string
	var start []int
	for i, line := range data {
		var row []string
		for j, cell := range line {
			if cell == 'S' {
				start = []int{i, j}
			}
			row = append(row, string(cell))
		}
		pipes = append(pipes, row)
	}
	return pipes, start
}

func SolvePartOne(data []string) int {
	pipes, start := parseData(data)
	row, col := start[0], start[1]

	// work out starting direction
	var in string
	if slices.Contains([]string{"|", "F", "7"}, pipes[row-1][col]) {
		row -= 1
		in = "u"
	} else if slices.Contains([]string{"|", "L", "J"}, pipes[row+1][col]) {
		row += 1
		in = "d"
	} else if slices.Contains([]string{"-", "F", "L"}, pipes[row][col-1]) {
		col -= 1
		in = "l"
	} else if slices.Contains([]string{"-", "J", "7"}, pipes[row][col+1]) {
		col += 1
		in = "r"
	} else {
		return 0
	}

	// count steps through pipes
	steps := 1
	for {
		nextPipe := pipes[row][col]
		if nextPipe == "S" {
			break
		}

		out := pipeMoves[nextPipe][in]
		row += directions[out].rowChange
		col += directions[out].colChange
		in = out
		steps++
	}

	return steps / 2
}

func SolvePartTwo(data []string) int {
	pipes, start := parseData(data)
	row, col := start[0], start[1]

	// work out starting direction
	var in string
	if slices.Contains([]string{"-", "F", "L"}, pipes[row][col-1]) {
		col -= 1
		in = "l"
	} else if slices.Contains([]string{"|", "F", "7"}, pipes[row-1][col]) {
		row -= 1
		in = "u"
	} else if slices.Contains([]string{"|", "L", "J"}, pipes[row+1][col]) {
		row += 1
		in = "d"
	} else if slices.Contains([]string{"-", "J", "7"}, pipes[row][col+1]) {
		col += 1
		in = "r"
	} else {
		return 0
	}

	// build path through pipes
	var path = [][]int{start}
	for {
		nextPipe := pipes[row][col]
		if nextPipe == "S" {
			break
		}

		path = append(path, []int{row, col})

		out := pipeMoves[nextPipe][in]
		pipes[row][col] = "S"
		row += directions[out].rowChange
		col += directions[out].colChange
		in = out
	}

	// Shoelace formula for polygonal area
	// https://en.wikipedia.org/wiki/Shoelace_formula
	total := 0
	for x := 0; x < len(path); x++ {
		tileA := path[x]
		tileB := path[(x+1)%len(path)]

		total += ((tileA[0] * tileB[1]) - (tileA[1] * tileB[0]))
	}
	A := int(math.Abs(float64(total))) / 2 // area

	// Pick's theorem for points in polygon: A = i + b/2 - 1
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	b := len(path) // boundary points

	return A - (b / 2) + 1 // = i (points inside)
}
