package main

import (
	"flag"
	"fmt"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 23

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

func parseData(data []string, part int) ([][]string, map[int]map[int]bool, int) {
	var grid [][]string
	visited := make(map[int]map[int]bool)
	var start int
	for y, line := range data {
		visited[y] = make(map[int]bool)
		var row []string
		for x, char := range line {
			visited[y][x] = false
			if char == '#' {
				row = append(row, "#")
			} else if part == 1 {
				row = append(row, string(char))
			} else {
				row = append(row, ".")
			}
			if y == 0 && char == '.' {
				start = x
			}
		}
		grid = append(grid, row)
	}
	return grid, visited, start
}

type Point struct{ row, col int }

var moves = map[string][]Point{
	".": {Point{+1, 0}, Point{-1, 0}, Point{0, +1}, Point{0, -1}},
	">": {Point{0, +1}},
	"v": {Point{+1, 0}},
}

type Target struct {
	targetNode Point
	distance   int
}

func getValidNext(grid [][]string, row int, col int) []Point {
	var validMoves []Point
	nextMoves := moves[grid[row][col]]
	for _, next := range nextMoves {
		nextRow := row + next.row
		nextCol := col + next.col
		if 0 <= nextRow && nextRow < len(grid) && 0 <= nextCol && nextCol < len(grid[0]) {
			if grid[nextRow][nextCol] != "#" {
				validMoves = append(validMoves, Point{nextRow, nextCol})
			}
		}
	}
	return validMoves
}

func findNextVertex(grid [][]string, row int, col int, prevRow int, prevCol int) (Point, int) {
	fmt.Println("find next", row, col, prevRow, prevCol)
	distance := 0
	for {
		distance++
		nextMoves := getValidNext(grid, row, col)

		if len(nextMoves) != 2 {
			// new junction
			return Point{row, col}, distance
		}

		for _, move := range nextMoves {
			if !(move.row == prevRow && move.col == prevCol) {
				prevRow, prevCol = row, col
				row, col = move.row, move.col
				break
			}
		}
	}
}

func buildGraph(grid [][]string, visited map[int]map[int]bool, start int) map[Point][]Target {
	vertices := utils.Queue[Point]{}  // need a queue of vertices to process
	vertices.Push(Point{0, start})    // init queue with start vertex
	graph := make(map[Point][]Target) // init the graph

	for {
		if vertices.Len() == 0 {
			break // graph complete
		}

		// process next in queue, if it hasn't been handled already
		vertex := vertices.Pop()
		if visited[vertex.row][vertex.col] {
			continue
		}
		visited[vertex.row][vertex.col] = true

		// find and store vertex neighbours
		var targets []Target
		for _, move := range getValidNext(grid, vertex.row, vertex.col) {
			nextVertex, distance := findNextVertex(grid, move.row, move.col, vertex.row, vertex.col)
			targets = append(targets, Target{nextVertex, distance})
			vertices.Push(nextVertex) // need to process
		}
		graph[vertex] = targets

	}

	return graph
}

func traverseGraph(graph map[Point][]Target, start Point, end int, visited map[Point]bool, steps int) int {
	if start.row == end {
		return steps
	}

	visited[start] = true

	maxDistance := 0

	for _, next := range graph[start] {
		if val, ok := visited[next.targetNode]; ok {
			if val {
				continue
			}
		}

		distance := traverseGraph(graph, next.targetNode, end, visited, steps+next.distance)

		if maxDistance < distance {
			maxDistance = distance
		}
	}

	visited[start] = false

	return maxDistance
}

func SolvePartOne(data []string) int {
	grid, visited, start := parseData(data, 1)
	graph := buildGraph(grid, visited, start)
	visitedTwo := make(map[Point]bool)
	return traverseGraph(graph, Point{0, start}, len(grid)-1, visitedTwo, 0)
}

func SolvePartTwo(data []string) int {
	grid, visited, start := parseData(data, 2)
	graph := buildGraph(grid, visited, start)
	visitedTwo := make(map[Point]bool)
	return traverseGraph(graph, Point{0, start}, len(grid)-1, visitedTwo, 0)
}
