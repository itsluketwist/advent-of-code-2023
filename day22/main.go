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

const Day = 22

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

type BrickType struct {
	id         int
	locations  [][]int
	lowest     int // sort by this
	vertical   bool
	resting    []int
	supporting []int
}

func (bri *BrickType) Drop(amount int) {
	for _, loc := range bri.locations {
		loc[2] -= amount
	}
	bri.lowest -= amount
}

func parseData(data []string) []BrickType {
	var bricks []BrickType
	for i, line := range data {
		ends := strings.Split(line, "~")
		startStrs := strings.Split(ends[0], ",")

		var start []int
		for _, startStr := range startStrs {
			startInt, _ := strconv.Atoi(startStr)
			start = append(start, startInt)
		}

		endStrs := strings.Split(ends[1], ",")
		var end []int
		for _, endStr := range endStrs {
			endInt, _ := strconv.Atoi(endStr)
			end = append(end, endInt)
		}

		changeAxis := -1
		changeAmount := 1
		for i := 0; i < 3; i++ {
			if (end[i] - start[i]) != 0 {
				changeAxis = i
				changeAmount += (end[i] - start[i])
				break
			}
		}

		var locations [][]int
		if changeAxis == -1 {
			locations = append(locations, start)
		} else {
			for j := 0; j < changeAmount; j++ {
				newLoc := []int{start[0], start[1], start[2]}
				newLoc[changeAxis] += j
				locations = append(locations, newLoc)
			}
		}

		bricks = append(bricks, BrickType{
			id:         i + 1,
			locations:  locations,
			lowest:     start[2],
			vertical:   changeAxis == 2,
			resting:    []int{},
			supporting: []int{},
		})

	}
	return bricks
}

func letBricksDrop(bricks []BrickType) (map[int]BrickType, int) {
	var space [10][10][500]int
	fallenBricks := 0
	brickMap := make(map[int]BrickType)

	sort.Slice(bricks, func(i, j int) bool {
		a := bricks[i]
		b := bricks[j]
		return a.lowest <= b.lowest
	})

	for _, brick := range bricks {

		if brick.lowest != 1 {
			// process if not already on the ground
			var z int

			if brick.vertical {
				// find lowest empty spot
				x := brick.locations[0][0]
				y := brick.locations[0][1]
				z = brick.lowest
				for {
					z--

					if z == 0 || space[x][y][z] != 0 {
						// hit floor or another brick
						brick.resting = append(brick.resting, space[x][y][z])
						break
					}
				}

			} else {
				z = brick.lowest
				for {
					z--
					var resting []int
					for _, loc := range brick.locations {
						check := space[loc[0]][loc[1]][z]
						if check != 0 && !slices.Contains(resting, check) {
							resting = append(resting, check)
						}
					}

					if z == 0 || len(resting) != 0 {
						// hit floor or another brick
						brick.resting = resting
						break
					}
				}
			}

			// update brick
			dropped := brick.lowest - (z + 1)
			if dropped != 0 {
				brick.Drop(dropped)
				fallenBricks++
			}
			for _, id := range brick.resting {
				updateBrick := brickMap[id]
				updateBrick.supporting = append(updateBrick.supporting, brick.id)
				brickMap[id] = updateBrick
			}

		}

		// add brick into space
		// fmt.Println(brick)
		for _, loc := range brick.locations {
			space[loc[0]][loc[1]][loc[2]] = brick.id
		}
		brickMap[brick.id] = brick // save to map

	}

	return brickMap, fallenBricks
}

func SolvePartOne(data []string) int {
	bricks := parseData(data)
	brickMap, _ := letBricksDrop(bricks)

	total := 0
	for _, brick := range brickMap {
		if len(brick.supporting) == 0 {
			total++

		} else {
			removable := true
			for _, supported := range brick.supporting {
				hasAnotherSupport := false
				for _, rested := range brickMap[supported].resting {
					if rested != brick.id {
						hasAnotherSupport = true
						break
					}
				}
				if !hasAnotherSupport {
					removable = false
					break
				}
			}

			if removable {
				total++
			}
		}
	}

	return total
}

func SolvePartTwo(data []string) int {
	bricks := parseData(data)
	brickMap, _ := letBricksDrop(bricks)

	total := 0
	for i := 1; i <= len(brickMap); i++ {
		var newBrickList []BrickType

		for id, brick := range brickMap {
			if id != i && id != 0 {
				var locations [][]int
				for _, loc := range brick.locations {
					locations = append(locations, []int{loc[0], loc[1], loc[2]})
				}
				newBrickList = append(newBrickList, BrickType{
					id:         brick.id,
					locations:  locations,
					lowest:     brick.lowest,
					vertical:   brick.vertical,
					supporting: []int{},
					resting:    []int{},
				})
			}
		}

		_, fallenBricks := letBricksDrop(newBrickList)
		total += fallenBricks
	}

	return total
}
