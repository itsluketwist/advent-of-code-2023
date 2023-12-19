package utils

import "math"

func GCD(a, b int) int {
	// find Greatest Common Divisor (GCD) via Euclidean algorithm
	// https://en.wikipedia.org/wiki/Euclidean_algorithm
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func LCM(nums []int) int {
	// find Least Common Multiple (LCM) via GCD
	a, b, nums := nums[0], nums[1], nums[2:]

	lcm := (a * b) / GCD(a, b)

	if len(nums) == 0 {
		return lcm
	} else {
		return LCM(append(nums, lcm))
	}
}

func AbsInt(number int) int {
	return int(math.Abs(float64((number))))
}

func PointsInside(coordinates [][]int, rightAngles bool) int {
	// Use the shoelace formula and Pick's theorem to calculate the points inside a polygon.
	// Coordinates provided must either contain all points in the boundary, or only have right angles.
	area := 0
	boundary := 0

	// Shoelace formula for polygonal area
	// https://en.wikipedia.org/wiki/Shoelace_formula
	for i := 0; i < len(coordinates); i++ {
		one := coordinates[i]
		two := coordinates[(i+1)%len(coordinates)]

		area += ((one[0] * two[1]) - (one[1] * two[0]))
		if rightAngles {
			boundary += AbsInt(one[0] - two[0] + one[1] - two[1])
		} else {
			boundary++
		}
	}
	A := AbsInt(area / 2) // area

	// Pick's theorem for points in polygon: A = i + b/2 - 1
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	b := boundary // boundary points

	return A - (b / 2) + 1 // = i (points inside)
}
