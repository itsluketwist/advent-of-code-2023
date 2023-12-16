package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 12

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

type SpringRecord struct {
	record         string
	damagedGroups  []int
	unknownIndexes []int
	missing        int
}

func parseData(data []string, repeat int) []SpringRecord {
	var springs []SpringRecord

	for _, line := range data {
		split := strings.Split(line, " ")

		record := split[0]
		numbers := split[1]
		for i := 0; i < repeat; i++ {
			record += "?"
			record += split[0]
			numbers += ","
			numbers += split[1]
		}
		record += "."

		var damagedGroups []int
		damagedTotal := 0
		strNums := strings.Split(numbers, ",")
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			damagedGroups = append(damagedGroups, num)
			damagedTotal += num
		}

		damagedKnown := 0
		var unknownIndexes []int
		for i, spr := range record {
			if spr == '#' {
				damagedKnown++
			} else if spr == '?' {
				unknownIndexes = append(unknownIndexes, i)
			}
		}

		springs = append(springs, SpringRecord{
			record:         record,
			damagedGroups:  damagedGroups,
			unknownIndexes: unknownIndexes,
			missing:        damagedTotal - damagedKnown,
		})

	}
	return springs
}

type State struct {
	position int // string index to consider from
	matches  int // count of currently matched spring groups
	current  int // current length of group being processed
}

func Permutations(sta State, rec SpringRecord, cache map[State]int) int {
	if num, ok := cache[sta]; ok {
		return num // use cached result
	}

	if sta.position == len(rec.record)-1 {
		// either have already matched all groups
		// or finishing to match the last group
		if (sta.matches == len(rec.damagedGroups) && sta.current == 0) ||
			(sta.matches == len(rec.damagedGroups)-1 && sta.current == rec.damagedGroups[len(rec.damagedGroups)-1]) {
			cache[sta] = 1
			return 1
		} else {
			cache[sta] = 0
			return 0
		}

	}

	hashResult := 0
	if rec.record[sta.position] == '#' || rec.record[sta.position] == '?' {
		if sta.current == 0 {
			// start new group
			if sta.matches == len(rec.damagedGroups) {
				hashResult = 0
			} else {
				hashResult = Permutations(State{sta.position + 1, sta.matches, 1}, rec, cache)
			}

		} else {
			// continue current group
			if sta.current+1 > rec.damagedGroups[sta.matches] {
				hashResult = 0
			} else {
				hashResult = Permutations(State{sta.position + 1, sta.matches, sta.current + 1}, rec, cache)
			}
		}

	}

	dotResult := 0
	if rec.record[sta.position] == '.' || rec.record[sta.position] == '?' {
		if sta.current == 0 {
			// no group, just continue
			dotResult = Permutations(State{sta.position + 1, sta.matches, 0}, rec, cache)
		} else {
			// end of a group
			if sta.current == rec.damagedGroups[sta.matches] {
				dotResult = Permutations(State{sta.position + 1, sta.matches + 1, 0}, rec, cache)
			} else {
				dotResult = 0
			}
		}

	}

	// cache and return the result
	result := dotResult + hashResult
	cache[sta] = result
	return result

}

func Solve(data []string, repeat int) int {
	parsed := parseData(data, repeat)

	total := 0
	for _, record := range parsed {
		blankCache := make(map[State]int)
		total += Permutations(State{0, 0, 0}, record, blankCache)
	}

	return total
}

func SolvePartOne(data []string) int {
	return Solve(data, 0)
}

func SolvePartTwo(data []string) int {
	return Solve(data, 4)
}
