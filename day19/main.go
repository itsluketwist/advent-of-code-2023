package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 19

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

type Rule struct {
	code   string
	less   bool
	value  int
	target string
}

type Workflow struct {
	rules    []Rule
	fallback string
}

func parseData(data []string) ([]map[string]int, map[string]Workflow) {
	var parts []map[string]int
	flows := make(map[string]Workflow)

	partRegex := regexp.MustCompile("^{x=(\\d+),m=(\\d+),a=(\\d+),s=(\\d+)}$")
	ruleRegex := regexp.MustCompile("^(\\w)([<>])(\\d+):(\\w+)$")

	for _, line := range data {
		if line == "" {
			continue

		} else if line[0] == '{' {
			match := partRegex.FindSubmatch([]byte(line))
			x, _ := strconv.Atoi(string(match[1]))
			m, _ := strconv.Atoi(string(match[2]))
			a, _ := strconv.Atoi(string(match[3]))
			s, _ := strconv.Atoi(string(match[4]))
			parts = append(parts, map[string]int{"x": x, "m": m, "a": a, "s": s})

		} else {
			split := strings.Split(line, "{")
			rules := strings.Split(strings.Trim(split[1], "}"), ",")
			var parsedRules []Rule
			for _, rule := range rules {
				match := ruleRegex.FindSubmatch([]byte(rule))
				if len(match) > 0 {
					value, _ := strconv.Atoi(string(match[3]))
					less := string(match[2]) == "<"
					parsedRules = append(parsedRules, Rule{string(match[1]), less, value, string(match[4])})
				}
			}
			flows[split[0]] = Workflow{parsedRules, rules[len(rules)-1]}
		}

	}

	return parts, flows

}

func SolvePartOne(data []string) int {
	total := 0
	parts, flows := parseData(data)

	for _, part := range parts {
		current := "in"
		for {
			if current == "R" {
				fmt.Println("REJECT")
				break
			} else if current == "A" {
				fmt.Println("ACCEPT")
				for _, val := range part {
					total += val
				}
				break
			}

			flow := flows[current]
			next := ""
			for _, rule := range flow.rules {
				if (rule.less && (part[rule.code] < rule.value)) || (!rule.less && (part[rule.code] > rule.value)) {
					next = rule.target
					break
				}
			}

			if next == "" {
				current = flow.fallback
			} else {
				current = next
			}
		}
	}

	return total
}

func copyMap(init map[string][]int) map[string][]int {
	new := make(map[string][]int)
	for k, v := range init {
		new[k] = []int{v[0], v[1]}
	}
	return new
}

type RangeMap struct {
	code string
	vals map[string][]int
}

func SolvePartTwo(data []string) int {
	total := 0
	_, flows := parseData(data)

	batch := []RangeMap{{"in", map[string][]int{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}}}}

	for i := 0; i < 10; i++ {
		nextBatch := make([]RangeMap, 0)

		for _, next := range batch {
			if next.code == "A" {
				total += (next.vals["x"][1] - next.vals["x"][0] + 1) * (next.vals["m"][1] - next.vals["m"][0] + 1) * (next.vals["a"][1] - next.vals["a"][0] + 1) * (next.vals["s"][1] - next.vals["s"][0] + 1)
			} else if next.code == "R" {
				continue
			}

			flow := flows[next.code]
			for _, rule := range flow.rules {
				if next.vals[rule.code][0] < rule.value && rule.value < next.vals[rule.code][1] {
					newVals := copyMap(next.vals)
					if rule.less {
						newVals[rule.code][1] = rule.value - 1
						next.vals[rule.code][0] = rule.value
					} else {
						newVals[rule.code][0] = rule.value + 1
						next.vals[rule.code][1] = rule.value
					}
					nextBatch = append(nextBatch, RangeMap{rule.target, newVals})
				}
			}
			newVals := copyMap(next.vals)
			nextBatch = append(nextBatch, RangeMap{flow.fallback, newVals})
		}

		if len(nextBatch) == 0 {
			break
		} else {
			batch = nextBatch
		}
	}

	return total
}
