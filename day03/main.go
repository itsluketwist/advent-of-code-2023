package main

import (
	"flag"
	"fmt"
	"github.com/itsluketwist/advent-of-code-2023/utils"
	"strconv"
)

const Day = 3

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

func parseData(data []string) [][]string {
	var grid [][]string
	for _, line := range data {
		arr := make([]string, len(line))
		for i, r := range line {
			arr[i] = string(r)
		}
		grid = append(grid, arr)
	}
	return grid
}

func SolvePartOne(data []string) int {
	total := 0
	grid := parseData(data)
	xLen := len(grid[0])
	yLen := len(grid)

	for y := 0; y < yLen; y++ {
		numStart, numEnd := -1, -1
		numStr := ""
		for x := 0; x < xLen; x++ {
			// check if cell is a number
			if _, err := strconv.Atoi(grid[y][x]); err == nil {
				if numStr == "" {
					numStart = x
				}
				numStr += grid[y][x]
			} else {
				if numStr != "" {
					numEnd = x - 1
				}
			}

			if numEnd != -1 || (numStr != "" && x == xLen-1) {
				// handle end of line
				if x == xLen-1 {
					numEnd = x
				}

				// have a number, check if adjacent to symbol
				if CheckForSymbol(grid, numStart, numEnd, y) == true {
					num, _ := strconv.Atoi(numStr)
					total += num
				}

				// reset
				numStart, numEnd = -1, -1
				numStr = ""
			}
		}
	}

	return total
}

func CheckForSymbol(grid [][]string, start int, end int, y int) bool {
	// edges
	if IsSymbol(grid, start-1, y) == true || IsSymbol(grid, end+1, y) == true {
		return true
	}

	// above and below
	for x := start - 1; x <= end+1; x++ {
		if IsSymbol(grid, x, y-1) == true || IsSymbol(grid, x, y+1) == true {
			return true
		}
	}

	// no symbol
	return false
}

func IsSymbol(grid [][]string, x int, y int) bool {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return false
	}

	char := grid[y][x]

	if char == "." {
		return false
	}
	if _, err := strconv.Atoi(char); err == nil {
		return false
	}
	return true
}

func SolvePartTwo(data []string) int {
	grid := parseData(data)
	xLen := len(grid[0])
	yLen := len(grid)

	// make stars unique
	starId := 0
	starMap := make(map[string][]int)
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if grid[y][x] == "*" {
				starId++
				starKey := fmt.Sprintf("*%d", starId)
				grid[y][x] = starKey
				var arr []int
				starMap[starKey] = arr
			}
		}
	}

	for y := 0; y < yLen; y++ {
		numStart, numEnd := -1, -1
		numStr := ""
		for x := 0; x < xLen; x++ {
			// check if cell is a number
			if _, err := strconv.Atoi(grid[y][x]); err == nil {
				if numStr == "" {
					numStart = x
				}
				numStr += grid[y][x]
			} else {
				if numStr != "" {
					numEnd = x - 1
				}
			}

			if numEnd != -1 || (numStr != "" && x == xLen-1) {
				// handle end of line
				if x == xLen-1 {
					numEnd = x
				}

				// get number and surrounding stars
				num, _ := strconv.Atoi(numStr)
				stars := CheckForStar(grid, numStart, numEnd, y)

				// store stars in the map
				for _, star := range stars {
					starMap[star] = append(starMap[star], num)
				}

				// reset
				numStart, numEnd = -1, -1
				numStr = ""
			}
		}
	}

	total := 0
	for _, numList := range starMap {
		if len(numList) == 2 {
			total += numList[0] * numList[1]
		}
	}
	return total
}

func CheckForStar(grid [][]string, start int, end int, y int) []string {
	var stars []string

	// edges
	if IsStar(grid, start-1, y) == true {
		stars = append(stars, grid[y][start-1])
	}

	if IsStar(grid, end+1, y) == true {
		stars = append(stars, grid[y][end+1])
	}

	// above and below
	for x := start - 1; x <= end+1; x++ {
		if IsStar(grid, x, y-1) == true {
			stars = append(stars, grid[y-1][x])
		}
		if IsStar(grid, x, y+1) == true {
			stars = append(stars, grid[y+1][x])
		}
	}

	return stars
}

func IsStar(grid [][]string, x int, y int) bool {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return false
	}

	if grid[y][x][0] == '*' {
		return true
	}
	return false
}
