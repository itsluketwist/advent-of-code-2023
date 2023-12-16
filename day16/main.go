package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 16

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
	var tiles [][]string
	for _, line := range data {
		tiles = append(tiles, strings.Split(line, ""))
	}
	return tiles
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

var tileMoves = map[string]map[string][]string{
	".":  {"u": {"u"}, "d": {"d"}, "l": {"l"}, "r": {"r"}},
	"/":  {"u": {"r"}, "d": {"l"}, "l": {"d"}, "r": {"u"}},
	"\\": {"u": {"l"}, "d": {"r"}, "l": {"u"}, "r": {"d"}},
	"-":  {"u": {"l", "r"}, "d": {"l", "r"}, "l": {"l"}, "r": {"r"}},
	"|":  {"u": {"u"}, "d": {"d"}, "l": {"u", "d"}, "r": {"u", "d"}},
}

type Location struct {
	row int
	col int
}

type TileMove struct {
	idx Location
	dir string
}

func processTile(
	move TileMove,
	grid [][]string,
	cache map[TileMove]bool,
	energized map[Location]bool,
) {
	if !cache[move] {
		cache[move] = true
		energized[move.idx] = true

		tileType := grid[move.idx.row][move.idx.col]
		nextDirections := tileMoves[tileType][move.dir]

		for _, next := range nextDirections {
			nextRow := move.idx.row + directions[next].rowChange
			nextCol := move.idx.col + directions[next].colChange
			if 0 <= nextRow && nextRow < len(grid) && 0 <= nextCol && nextCol < len(grid[0]) {
				processTile(TileMove{Location{nextRow, nextCol}, next}, grid, cache, energized)
			}
		}
	}
}

func Solve(row int, col int, dir string, grid [][]string) int {
	cache := make(map[TileMove]bool)
	energized := make(map[Location]bool)
	processTile(TileMove{Location{row, col}, dir}, grid, cache, energized)
	return len(energized)
}

func SolvePartOne(data []string) int {
	grid := parseData(data)
	return Solve(0, 0, "r", grid)
}

func SolvePartTwo(data []string) int {
	max := 0
	var next int
	grid := parseData(data)

	for n := 0; n < len(grid); n++ {
		// leftmost column
		next = Solve(n, 0, "r", grid)
		if next > max {
			max = next
		}

		// rightmost column
		next = Solve(n, len(grid[0])-1, "l", grid)
		if next > max {
			max = next
		}
	}

	for n := 0; n < len(grid[0]); n++ {
		// top row
		next = Solve(0, n, "d", grid)
		if next > max {
			max = next
		}

		// bottom row
		next = Solve(len(grid)-1, n, "u", grid)
		if next > max {
			max = next
		}
	}

	return max
}
