package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 18

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

type Dig struct {
	dir string
	num int
	hex string
}

func parseDataOne(data []string) []Dig {
	var digs []Dig
	for _, line := range data {
		lineSplit := strings.Split(line, " ")
		num, _ := strconv.Atoi(lineSplit[1])
		hex := strings.Trim(lineSplit[2], "()")
		digs = append(digs, Dig{lineSplit[0], num, hex})
	}
	return digs
}

func parseDataTwo(data []string) []Dig {
	var directionCodes = map[byte]string{
		'0': "R", '1': "D", '2': "L", '3': "U",
	}

	var digs []Dig
	for _, line := range data {
		hex := strings.Split(line, " ")[2]
		num, _ := strconv.ParseInt(hex[2:7], 16, 64)

		digs = append(digs, Dig{directionCodes[hex[7]], int(num), hex})
	}
	return digs
}

type Move struct {
	xDiff int
	yDiff int
}

var moves = map[string]Move{
	"U": {0, +1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {+1, 0},
}

func Solve(plan []Dig) int {
	x, y, outside := 0, 0, 0
	var points [][]int
	for _, dig := range plan {
		points = append(points, []int{x, y})
		x += dig.num * moves[dig.dir].xDiff
		y += dig.num * moves[dig.dir].yDiff
		outside += dig.num
	}

	// use Shoelace formula and Pick's theorem
	inside := utils.PointsInside(points, true)
	return outside + inside
}

func SolvePartOne(data []string) int {
	return Solve(parseDataOne(data))
}

func SolvePartTwo(data []string) int {
	return Solve(parseDataTwo(data))
}
