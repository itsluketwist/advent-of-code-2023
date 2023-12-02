package main

import (
	"flag"
	"fmt"
	"github.com/itsluketwist/advent-of-code-2023/utils"
	"strconv"
)

const Day = 0 // actually 2020 day01, used for practice/setup

func main() {
	part := flag.Int("part", 0, "Which parts to try?")
	try := flag.Int("try", 0, "Whether to try the real input?")
	flag.Parse()

	fmt.Println("Running day", Day, "( part:", *part, ", try:", *try, ")")

	example, _ := utils.ReadFileToArray(Day, "example", false)
	input, _ := utils.ReadFileToArray(Day, "input", false)

	if *part == 0 || *part == 1 {
		solutionOneExample := SolvePartOne(example)
		fmt.Println("Solution to part 1 (example):", solutionOneExample)

		if *try == 1 {
			SolutionOneInput := SolvePartOne(input)
			fmt.Println("Solution to part 1 (input):", SolutionOneInput)
		}
	}

	if *part == 0 || *part == 2 {
		SolutionTwoExample := SolvePartTwo(example)
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
	parsed := parseData(data)

	for _, x := range parsed {
		for _, y := range parsed {
			if x+y == 2020 {
				return x * y
			}
		}
	}

	return 0
}

func SolvePartTwo(data []string) int {
	parsed := parseData(data)

	for _, x := range parsed {
		for _, y := range parsed {
			for _, z := range parsed {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}

	return 0
}
