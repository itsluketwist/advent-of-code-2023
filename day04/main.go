package main

import (
	"flag"
	"fmt"
	"github.com/itsluketwist/advent-of-code-2023/utils"
	"slices"
	"strconv"
	"strings"
)

const Day = 4

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

func SolvePartOne(data []string) int {
	total := 0
	for _, card := range data {
		value := 0
		splitCard := strings.Split(card, ":")
		splitNums := strings.Split(splitCard[1], "|")

		var winning []int
		for _, numStr := range strings.Split(splitNums[0], " ") {
			if num, err := strconv.Atoi(numStr); err == nil {
				winning = append(winning, num)
			}
		}

		for _, numStr := range strings.Split(splitNums[1], " ") {
			if num, err := strconv.Atoi(numStr); err == nil {
				if slices.Contains(winning, num) {
					if value == 0 {
						value = 1
					} else {
						value *= 2
					}
				}
			}
		}

		total += value
	}

	return total
}

func SolvePartTwo(data []string) int {
	total := 0
	cardCount := make([]int, len(data))

	for id, card := range data {
		cardCount[id] += 1
		total += cardCount[id]

		splitCard := strings.Split(card, ":")
		splitNums := strings.Split(splitCard[1], "|")

		var winning []int
		for _, numStr := range strings.Split(splitNums[0], " ") {
			if num, err := strconv.Atoi(numStr); err == nil {
				winning = append(winning, num)
			}
		}

		winners := 0
		for _, numStr := range strings.Split(splitNums[1], " ") {
			if num, err := strconv.Atoi(numStr); err == nil {
				if slices.Contains(winning, num) {
					winners += 1
				}
			}
		}

		for i := 0; i < winners; i++ {
			cardCount[id+i+1] += cardCount[id]
		}
	}

	return total
}
