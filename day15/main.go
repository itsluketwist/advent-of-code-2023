package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 15

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

func parseData(data []string) []string {
	return strings.Split(data[0], ",")
}

func hashAlgorithm(code string) int {
	codeValue := 0
	for _, char := range code {
		ascii := int(char)
		codeValue += ascii
		codeValue *= 17
		codeValue %= 256
	}
	return codeValue
}

func SolvePartOne(data []string) int {
	total := 0
	codes := parseData(data)

	for _, code := range codes {
		total += hashAlgorithm(code)
	}

	return total
}

type Lens struct {
	label string
	value int
}

func SolvePartTwo(data []string) int {
	codes := parseData(data)

	hashMap := make(map[int][]Lens)

	for _, code := range codes {

		if strings.Contains(code, "=") {
			codeSplit := strings.Split(code, "=")
			label := codeSplit[0]
			hash := hashAlgorithm(label)
			value, _ := strconv.Atoi(codeSplit[1])

			// add lens to list
			setIndex := -1
			for i := 0; i < len(hashMap[hash]); i++ {
				if hashMap[hash][i].label == label {
					setIndex = i
					break
				}
			}
			if setIndex != -1 {
				hashMap[hash][setIndex] = Lens{label, value}
			} else {
				hashMap[hash] = append(hashMap[hash], Lens{label, value})
			}

		} else if strings.Contains(code, "-") {
			label := code[:len(code)-1]
			hash := hashAlgorithm(label)

			// replace lens in list
			for i := 0; i < len(hashMap[hash]); i++ {
				if hashMap[hash][i].label == label {
					copy(hashMap[hash][i:], hashMap[hash][i+1:])
					hashMap[hash] = hashMap[hash][:len(hashMap[hash])-1]
					break
				}
			}

		}
	}

	total := 0
	for box, lenses := range hashMap {
		for k, lens := range lenses {
			total += ((1 + box) * (1 + k) * lens.value)
		}
	}

	return total
}
