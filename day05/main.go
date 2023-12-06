package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 5

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

func parseData(data []string) [][][]int {
	var maps [][][]int
	var currentMap [][]int
	inMap := false
	for _, line := range data {
		if strings.Contains(line, "map") {
			// map data started
			inMap = true
			currentMap = [][]int{}

		} else if inMap && line == "" {
			// end of map, stop processing
			maps = append(maps, currentMap)
			inMap = false

		} else if inMap && line != "" {
			// process line of map info
			var mapInfo []int
			for _, numStr := range strings.Split(line, " ") {
				if num, err := strconv.Atoi(numStr); err == nil {
					mapInfo = append(mapInfo, num)
				}
			}
			currentMap = append(currentMap, mapInfo)

		}

	}
	return maps
}

func SolvePartOne(data []string) int {
	var seeds []int
	for _, numStr := range strings.Split(data[0], " ") {
		if num, err := strconv.Atoi(numStr); err == nil {
			seeds = append(seeds, num)
		}
	}

	maps := parseData(data)
	for i := 0; i < len(seeds); i++ {
		for _, mapData := range maps {
			updated := false
			for _, m := range mapData {
				if m[1] <= seeds[i] && seeds[i] < (m[1]+m[2]) && !updated {
					seeds[i] += (m[0] - m[1])
					updated = true
				}
			}
		}
	}

	min := seeds[0]
	for _, seed := range seeds {
		if seed < min {
			min = seed
		}
	}
	return min
}

func SolvePartTwo(data []string) int {
	var initSeeds []int
	for _, numStr := range strings.Split(data[0], " ") {
		if num, err := strconv.Atoi(numStr); err == nil {
			initSeeds = append(initSeeds, num)
		}
	}

	var seeds []int
	for j := 0; j < len(initSeeds)/2; j++ {
		for k := 0; k < initSeeds[2*j+1]; k++ {
			seeds = append(seeds, initSeeds[2*j]+k)
		}
	}

	maps := parseData(data)

	min := -1
	for _, next := range seeds {
		seed := next
		for _, mapData := range maps {
			updated := false
			for _, m := range mapData {
				if m[1] <= seed && seed < (m[1]+m[2]) && !updated {
					seed += (m[0] - m[1])
					updated = true
				}
			}
		}

		if min == -1 || seed < min {
			min = seed
		}

	}

	return min
}
