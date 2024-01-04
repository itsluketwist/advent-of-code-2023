package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/itsluketwist/advent-of-code-2023/utils"
)

const Day = 24

func main() {
	part := flag.Int("part", 0, "Which parts to try?")
	try := flag.Int("try", 0, "Whether to try the real input?")
	flag.Parse()

	fmt.Println("Running day", Day, "( part:", *part, ", try:", *try, ")")

	exampleOne, _ := utils.ReadFileToArray(Day, "example1", false)
	input, _ := utils.ReadFileToArray(Day, "input", false)

	if *part == 0 || *part == 1 {
		solutionOneExample := SolvePartOne(exampleOne, 7, 27)
		fmt.Println("Solution to part 1 (example):", solutionOneExample)

		if *try == 1 {
			SolutionOneInput := SolvePartOne(input, 200000000000000, 400000000000000)
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

type Hailstone struct {
	pos []float64
	vel []float64
}

func parseData(data []string) []Hailstone {
	var stones []Hailstone
	for _, line := range data {
		bits := strings.Split(line, " @ ")

		var pos []float64
		posStrs := strings.Split(bits[0], ", ")
		for _, str := range posStrs {
			num, _ := strconv.Atoi(str)
			pos = append(pos, float64(num))
		}

		var vel []float64
		velStrs := strings.Split(bits[1], ", ")
		for _, str := range velStrs {
			num, _ := strconv.Atoi(str)
			vel = append(vel, float64(num))
		}

		stones = append(stones, Hailstone{pos, vel})
	}
	return stones
}

func crossOver(stoneA Hailstone, stoneB Hailstone, min float64, max float64) bool {
	// linear function equation: y = grad*x + con
	var x, y float64 // calculate point of intersection

	if stoneA.vel[0] == 0 && stoneB.vel[0] == 0 {
		return false // parallel => no intersection

	} else if stoneA.vel[0] == 0 {
		x = stoneA.pos[0] // doesn't change for A

		gradB := stoneB.vel[1] / stoneB.vel[0]
		conB := stoneB.pos[1] - gradB*stoneB.pos[0]
		y = gradB*x + conB

	} else if stoneB.vel[0] == 0 {
		x = stoneB.pos[0] // doesn't change for B

		gradA := stoneA.vel[1] / stoneA.vel[0]
		conA := stoneA.pos[1] - gradA*stoneB.pos[0]
		y = gradA*x + conA

	} else {
		// calculate gradients of the lines
		gradA := stoneA.vel[1] / stoneA.vel[0]
		gradB := stoneB.vel[1] / stoneB.vel[0]

		if gradA == gradB {
			return false // parallel => no intersection
		}

		// calculate constants using known values
		conA := stoneA.pos[1] - gradA*stoneA.pos[0]
		conB := stoneB.pos[1] - gradB*stoneB.pos[0]

		// solve for x: gradA*x + conA = gradB*x + conB
		x = (conB - conA) / (gradA - gradB)
		// calculate y using one of the lines
		y = gradA*x + conA
	}

	if x < min || max < x || y < min || max < y {
		return false // intersection out of range
	}

	if ((stoneA.vel[0] > 0 && x > stoneA.pos[0]) || (stoneA.vel[0] < 0 && x < stoneA.pos[0])) && ((stoneB.vel[0] > 0 && x > stoneB.pos[0]) || (stoneB.vel[0] < 0 && x < stoneB.pos[0])) {
		return true // intersection is in the future
	}

	return false
}

func SolvePartOne(data []string, min int, max int) int {
	total := 0
	stones := parseData(data)

	// loop through each pair of hailstones
	for i := 0; i < len(stones); i++ {
		stoneA := stones[i]
		for j := i + 1; j < len(stones); j++ {
			stoneB := stones[j]

			if crossOver(stoneA, stoneB, float64(min), float64(max)) {
				total++
			}

		}
	}

	return total
}

func SolvePartTwo(data []string) int {
	stones := parseData(data)

	fmt.Println("Rock starts at position (pX, pY, pZ) and moves at velocity (vX, xY, vZ).")
	fmt.Println("Rock intersects all stones at some time.")
	fmt.Println("Can use 3 stones, and 3 points in time, to create simultaneous equations to solve.")
	fmt.Println()

	for i := 0; i < 3; i++ {
		s := stones[i]
		t := "t" + fmt.Sprint(i+1)
		fmt.Println("Intersect stone", stones[i], "at time", t)
		fmt.Println(s.pos[0], "+", t, "*", s.vel[0], "=", "pX + vX *", t)
		fmt.Println(s.pos[1], "+", t, "*", s.vel[1], "=", "pY + vY *", t)
		fmt.Println(s.pos[2], "+", t, "*", s.vel[2], "=", "pZ + vZ *", t)
		fmt.Println()
	}

	rock := Hailstone{
		pos: []float64{393358484426865, 319768494554521, 158856878271783},
		vel: []float64{-242, -49, 209},
	}

	fmt.Println("Have rock:", rock)

	return int(rock.pos[0] + rock.pos[1] + rock.pos[2])
}
