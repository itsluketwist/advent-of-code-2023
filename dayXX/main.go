package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 0

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

func parseData(data []string) []int {
	var numbers []int
	for _, strNum := range data {
		number, _ := strconv.Atoi(strNum)
		numbers = append(numbers, int(number))
	}
	return numbers
}

func SolvePartOne(data []string) int {
	total := 0
	parsed := parseData(data)
	fmt.Println(parsed)

	return total
}

func SolvePartTwo(data []string) int {
	total := 0
	parsed := parseData(data)
	fmt.Println(parsed)

	return total
}
