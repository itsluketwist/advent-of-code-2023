package main

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 17

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

func parseData(data []string) [][]int {
	var grid [][]int
	for _, line := range data {
		var numLine []int
		for _, numStr := range strings.Split(line, "") {
			num, _ := strconv.Atoi(string(numStr))
			numLine = append(numLine, num)
		}
		grid = append(grid, numLine)
	}
	return grid
}

type Node struct {
	row int
	col int
	dir string
}

func Solve(data []string, min int, max int) int {
	grid := parseData(data)
	heap := utils.BucketHeap[Node]{}
	heap.Add(Node{row: 0, col: 0, dir: "s"}, 0)
	var visited []Node

	for {

		next, value := heap.Pop()
		if slices.Contains(visited, next) {
			continue
		}
		visited = append(visited, next)

		if next.row == (len(grid)-1) && next.col == (len(grid[0])-1) {
			return value
		}

		last_dir := string(next.dir[0])

		if last_dir != "r" && last_dir != "l" {
			var lefStr, rigStr string
			var lefNum, rigNum int
			for i := 1; i <= max; i++ {
				if next.col+i < len(grid[0]) {
					// can move r(ight)
					rigStr += "r"
					rigNum += grid[next.row][next.col+i]
					if min <= i {
						heap.Add(Node{row: next.row, col: next.col + i, dir: rigStr}, value+rigNum)
					}
				}

				if 0 <= next.col-i {
					// can move l(eft)
					lefStr += "l"
					lefNum += grid[next.row][next.col-i]
					if min <= i {
						heap.Add(Node{row: next.row, col: next.col - i, dir: lefStr}, value+lefNum)
					}
				}
			}
		}

		if last_dir != "u" && last_dir != "d" {
			var upStr, dowStr string
			var upNum, dowNum int
			for i := 1; i <= max; i++ {
				if next.row+i < len(grid) {
					// can move d(own)
					dowStr += "d"
					dowNum += grid[next.row+i][next.col]
					if min <= i {
						heap.Add(Node{row: next.row + i, col: next.col, dir: dowStr}, value+dowNum)
					}
				}

				if 0 <= next.row-i {
					// can move u(p)
					upStr += "u"
					upNum += grid[next.row-i][next.col]
					if min <= i {
						heap.Add(Node{row: next.row - i, col: next.col, dir: upStr}, value+upNum)
					}
				}
			}
		}

	}
}

func SolvePartOne(data []string) int {
	return Solve(data, 1, 3)
}

func SolvePartTwo(data []string) int {
	return Solve(data, 4, 10)
}
