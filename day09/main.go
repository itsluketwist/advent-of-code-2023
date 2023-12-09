package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 9

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
	var parsed [][]int
	for _, line := range data {
		var numbers []int
		for _, strNum := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(strNum)
			numbers = append(numbers, int(number))
		}
		parsed = append(parsed, numbers)
	}
	return parsed
}

func SolvePartOne(data []string) int {
	seqs := parseData(data)
	total := 0

	for _, seq := range seqs {
		var subSeqs [][]int
		subSeqs = append(subSeqs, seq)
		nextSeq := seq
		for {
			hasNonZero := false
			var newSeq []int
			for i := 0; i < (len(nextSeq) - 1); i++ {
				next := nextSeq[i+1] - nextSeq[i]
				newSeq = append(newSeq, next)
				if next != 0 {
					hasNonZero = true
				}
			}
			subSeqs = append(subSeqs, newSeq)
			nextSeq = newSeq

			if !hasNonZero {
				break
			}
		}

		addition := 0
		for j := 0; j < len(subSeqs); j++ {
			addition += subSeqs[j][len(subSeqs[j])-1]
		}

		total += addition
	}

	return total
}

func SolvePartTwo(data []string) int {
	seqs := parseData(data)
	total := 0

	for _, seq := range seqs {
		var subSeqs [][]int
		subSeqs = append(subSeqs, seq)
		nextSeq := seq
		for {
			hasNonZero := false
			var newSeq []int
			for i := 0; i < (len(nextSeq) - 1); i++ {
				next := nextSeq[i+1] - nextSeq[i]
				newSeq = append(newSeq, next)
				if next != 0 {
					hasNonZero = true
				}
			}
			subSeqs = append(subSeqs, newSeq)
			nextSeq = newSeq

			if !hasNonZero {
				break
			}
		}

		previous := 0
		for j := len(subSeqs) - 2; j >= 0; j-- {
			previous = subSeqs[j][0] - previous
		}

		total += previous
	}

	return total
}
