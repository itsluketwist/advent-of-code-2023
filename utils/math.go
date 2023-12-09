package utils

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
