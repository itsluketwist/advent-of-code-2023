package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 1

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
	for _, line := range data {
		var first rune
		var last rune
		for _, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				if first == 0 {
					first = char
				}
				last = char
			}
		}
		num, _ := strconv.Atoi(string(first) + string(last))
		total += num
	}

	return total
}

func SolvePartTwo(data []string) int {
	total := 0
	var numStrMap = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range data {
		foundMap := make(map[int]string)

		for k, v := range numStrMap {
			if idx := strings.Index(line, k); idx != -1 {
				foundMap[idx] = v
			}
			if idx := strings.LastIndex(line, k); idx != -1 {
				foundMap[idx] = v
			}
		}

		for i, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				foundMap[i] = string(char)
			}
		}

		keys := []int{}
		for k, _ := range foundMap {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		numStr := foundMap[keys[0]] + foundMap[keys[len(keys)-1]]
		num, _ := strconv.Atoi(numStr)
		total += num
	}

	return total
}
