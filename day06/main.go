package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 6

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
	var times []int
	for _, strNum := range strings.Split(data[0], " ") {
		if number, err := strconv.Atoi(strNum); err == nil {
			times = append(times, number)
		}
	}

	var dists []int
	for _, strNum := range strings.Split(data[1], " ") {
		if number, err := strconv.Atoi(strNum); err == nil {
			dists = append(dists, number)
		}
	}

	result := [][]int{times, dists}
	return result
}

func SolvePartOne(data []string) int {
	parsed := parseData(data)

	final := 1
	races := len(parsed[0])
	for race := 0; race < races; race++ {
		count := 0
		for holdTime := 0; holdTime < parsed[0][race]; holdTime++ {
			if (parsed[0][race]-holdTime)*holdTime > parsed[1][race] {
				count++
			}
		}
		final *= count
	}

	return final
}

func SolvePartTwo(data []string) int {
	timeLine := strings.Split(strings.ReplaceAll(data[0], " ", ""), ":")
	time, _ := strconv.Atoi(timeLine[1])
	distLine := strings.Split(strings.ReplaceAll(data[1], " ", ""), ":")
	dist, _ := strconv.Atoi(distLine[1])

	count := 0
	for holdTime := 0; holdTime < time; holdTime++ {
		if (time-holdTime)*holdTime > dist {
			count++
		}
	}

	return count
}
