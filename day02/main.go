package main

import (
	"flag"
	"fmt"
	"github.com/itsluketwist/advent-of-code-2023/utils"
	"regexp"
	"strconv"
)

const Day = 2

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

func SolvePartOne(data []string) int {
	total := 0
	var colorCount = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for i, line := range data {
		possible := true
		for color, count := range colorCount {
			regexStr := "(\\d+) " + color
			re := regexp.MustCompile(regexStr)
			matches := re.FindAllSubmatch([]byte(line), -1)

			for _, match := range matches {
				matchCount, _ := strconv.Atoi(string(match[1]))
				if matchCount > count {
					possible = false
				}
			}
		}

		if possible == true {
			total += i + 1
		}
	}
	return total
}

func SolvePartTwo(data []string) int {
	total := 0
	var colorCount = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, line := range data {
		power := 1
		for color, _ := range colorCount {
			regexStr := "(\\d+) " + color
			re := regexp.MustCompile(regexStr)
			matches := re.FindAllSubmatch([]byte(line), -1)

			max := 0
			for _, match := range matches {
				matchCount, _ := strconv.Atoi(string(match[1]))
				if matchCount > max {
					max = matchCount
				}
			}
			power *= max
		}

		total += power
	}
	return total
}
