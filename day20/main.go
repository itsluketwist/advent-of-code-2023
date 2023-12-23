package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 20

func main() {
	part := flag.Int("part", 0, "Which parts to try?")
	try := flag.Int("try", 0, "Whether to try the real input?")
	flag.Parse()

	fmt.Println("Running day", Day, "( part:", *part, ", try:", *try, ")")

	exampleOne, _ := utils.ReadFileToArray(Day, "example1", false)
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
		if *try == 1 {
			SolutionTwoInput := SolvePartTwo(input)
			fmt.Println("Solution to part 2 (input):", SolutionTwoInput)
		}
	}
}

type FlipFlop struct {
	on      bool
	targets []string
}

type Conjunction struct {
	memory  map[string]bool
	targets []string
}

func parseData(data []string) ([]string, map[string]FlipFlop, map[string]Conjunction) {
	flipFlops := make(map[string]FlipFlop)
	conjunctions := make(map[string]Conjunction)
	var broadcast []string
	re := regexp.MustCompile("^([%&]?)(\\w+) -> (.*)$")
	for _, line := range data {
		match := re.FindSubmatch([]byte(line))

		label := string(match[2])
		targets := strings.Split(string(match[3]), ", ")

		if label == "broadcaster" {
			broadcast = targets
		} else if string(match[1]) == "%" {
			flipFlops[label] = FlipFlop{false, targets}
		} else {
			conjunctions[label] = Conjunction{make(map[string]bool), targets}
		}

	}

	for label, flip := range flipFlops {
		for _, target := range flip.targets {
			if _, ok := conjunctions[target]; ok {
				conjunctions[target].memory[label] = false
			}
		}
	}
	for label, con := range conjunctions {
		for _, target := range con.targets {
			if _, ok := conjunctions[target]; ok {
				conjunctions[target].memory[label] = false
			}
		}
	}

	return broadcast, flipFlops, conjunctions
}

type Pulse struct {
	target  string
	trigger string
	high    bool
}

func Solve(data []string, part int) int {
	highCount, lowCount := 0, 0
	broadcast, flipFlops, conjunctions := parseData(data)

	var loopMax int
	if part == 1 {
		loopMax = 1000
	} else {
		loopMax = 10000
	}

	queue := utils.Queue[Pulse]{}
	var partTwoNums []int
	pressed := map[string]bool{
		"kf": false, "kr": false, "qk": false, "zs": false,
	}
	for i := 0; i < loopMax; i++ {
		if len(partTwoNums) == 4 {
			return utils.LCM(partTwoNums)
		}

		lowCount++ // button pressed
		for _, target := range broadcast {
			queue.Push(Pulse{target, "broadcast", false})
		}

		for {
			if queue.Len() == 0 {
				break
			}

			pulse := queue.Pop()
			if pulse.high {
				highCount++
			} else {
				lowCount++
			}

			for key, val := range pressed {
				if !val && pulse.target == "gf" && conjunctions["gf"].memory[key] {
					pressed[key] = true
					partTwoNums = append(partTwoNums, i+1)

				}
			}

			if flip, ok := flipFlops[pulse.target]; ok {
				if !pulse.high {
					flip.on = !flip.on
					for _, t := range flip.targets {
						queue.Push(Pulse{t, pulse.target, flip.on})
					}
					flipFlops[pulse.target] = flip
				}
			} else if con, ok := conjunctions[pulse.target]; ok {
				con.memory[pulse.trigger] = pulse.high
				allHigh := true
				for _, v := range con.memory {
					allHigh = allHigh && v
				}
				for _, t := range con.targets {
					queue.Push(Pulse{t, pulse.target, !allHigh})
				}
				conjunctions[pulse.target] = con
			}
		}

	}

	return highCount * lowCount
}

func SolvePartOne(data []string) int {
	return Solve(data, 1)
}

func SolvePartTwo(data []string) int {
	return Solve(data, 2)
}
