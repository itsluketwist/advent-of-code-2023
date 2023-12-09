package main

import (
	"flag"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 7

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

var handTypeStrength = map[string]int{
	"five":      7,
	"four":      6,
	"fullhouse": 5,
	"three":     4,
	"twopair":   3,
	"onepair":   2,
	"high":      1,
}

var cardValue = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var cardValueWithJoker = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

type Hand struct {
	cards    string
	handType string
	value    int
}

func parseData(data []string, joker bool) []Hand {
	var hands []Hand
	for _, line := range data {
		lineSplit := strings.Split(line, " ")
		value, _ := strconv.Atoi(lineSplit[1])

		cardCount := make(map[rune]int)
		for _, card := range lineSplit[0] {
			cardCount[card]++
		}

		var countNums []int
		jokerAdjust := 0
		for key, count := range cardCount {
			if joker && key == 'J' {
				jokerAdjust = count
			} else {
				countNums = append(countNums, count)
			}
		}
		sort.Slice(countNums, func(i, j int) bool {
			return countNums[j] < countNums[i]
		})
		if len(countNums) == 0 {
			countNums = append(countNums, 5)
		} else {
			countNums[0] += jokerAdjust
		}

		var handType string
		switch unique := len(countNums); unique {
		case 1:
			handType = "five"
		case 2:
			if slices.Contains(countNums, 4) {
				handType = "four"
			} else {
				handType = "fullhouse"
			}
		case 3:
			if slices.Contains(countNums, 3) {
				handType = "three"
			} else {
				handType = "twopair"
			}
		case 4:
			handType = "onepair"
		case 5:
			handType = "high"

		}

		hand := Hand{
			value:    value,
			cards:    lineSplit[0],
			handType: handType,
		}

		hands = append(hands, hand)

	}

	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		if handTypeStrength[a.handType] != handTypeStrength[b.handType] {
			return handTypeStrength[a.handType] < handTypeStrength[b.handType]
		}

		var valueMap map[string]int
		if joker {
			valueMap = cardValueWithJoker
		} else {
			valueMap = cardValue
		}
		for i := 0; i < 5; i++ {
			if valueMap[string(a.cards[i])] != valueMap[string(b.cards[i])] {
				return valueMap[string(a.cards[i])] < valueMap[string(b.cards[i])]
			}
		}

		return false
	})

	return hands
}

func SolvePartOne(data []string) int {
	hands := parseData(data, false)

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.value
	}

	return total
}

func SolvePartTwo(data []string) int {
	hands := parseData(data, true)

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.value
	}

	return total
}
