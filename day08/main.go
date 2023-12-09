package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 8

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

func parseData(data []string) map[string][]string {
	ruleMap := make(map[string][]string)

	for _, line := range data {
		regexStr := "^([\\w\\d]+) = \\(([\\w\\d]+), ([\\w\\d]+)\\)$"
		re := regexp.MustCompile(regexStr)
		match := re.FindSubmatch([]byte(line))

		if match != nil {
			ruleMap[string(match[1])] = []string{string(match[2]), string(match[3])}
		}
	}

	return ruleMap
}

func SolvePartOne(data []string) int {
	path := data[0]
	ruleMap := parseData(data)

	i := 0
	steps := 0
	node := "AAA"
	for {
		if path[i] == 'L' {
			node = ruleMap[node][0]
		} else if path[i] == 'R' {
			node = ruleMap[node][1]
		}

		i = (i + 1) % len(path)
		steps++

		if node == "ZZZ" {
			return steps
		}
	}
}

func SolvePartTwo(data []string) int {
	path := data[0]
	ruleMap := parseData(data)

	var nodes []string
	for node := range ruleMap {
		if node[2] == 'A' {
			nodes = append(nodes, node)
		}
	}

	var cycleLen []int
	for _, node := range nodes {
		i := 0
		steps := 0
		for {

			if path[i] == 'L' {
				node = ruleMap[node][0]
			} else if path[i] == 'R' {
				node = ruleMap[node][1]
			}

			i = (i + 1) % len(path)
			steps++

			if node[2] == 'Z' {
				cycleLen = append(cycleLen, steps)
				break
			}

		}

	}

	return utils.LCM(cycleLen)
}
