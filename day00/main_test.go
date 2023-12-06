package main

import (
	"testing"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

var resultsMap = map[string]int{
	"PartOneExample": 514579,
	"PartOneInput":   888331,
	"PartTwoExample": 241861950,
	"PartTwoInput":   130933530,
}

func TestSolvePartOneExample(t *testing.T) {
	test := "PartOneExample"
	expected := resultsMap[test]

	data, _ := utils.ReadFileToArray(Day, "example", true)
	result := SolvePartOne(data)

	if expected != result {
		t.Errorf("Day%02d %s FAIL! Expected %d but got %d", Day, test, expected, result)
		t.Fail()
	}
}

func TestSolvePartOneInput(t *testing.T) {
	test := "PartOneInput"
	expected := resultsMap[test]

	data, _ := utils.ReadFileToArray(Day, "input", true)
	result := SolvePartOne(data)

	if expected != result {
		t.Errorf("Day%02d %s FAIL! Expected %d but got %d", Day, test, expected, result)
		t.Fail()
	}
}

func TestSolvePartTwoExample(t *testing.T) {
	test := "PartTwoExample"
	expected := resultsMap[test]

	data, _ := utils.ReadFileToArray(Day, "example", true)
	result := SolvePartTwo(data)

	if expected != result {
		t.Errorf("Day%02d %s FAIL! Expected %d but got %d", Day, test, expected, result)
		t.Fail()
	}
}

func TestSolvePartTwoInput(t *testing.T) {
	test := "PartTwoInput"
	expected := resultsMap[test]

	data, _ := utils.ReadFileToArray(Day, "input", true)
	result := SolvePartTwo(data)

	if expected != result {
		t.Errorf("Day%02d %s FAIL! Expected %d but got %d", Day, test, expected, result)
		t.Fail()
	}
}
