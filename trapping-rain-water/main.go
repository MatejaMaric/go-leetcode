package main

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	if height[0] == 0 {
		return trap(height[1:])
	}

	if height[len(height)-1] == 0 {
		return trap(height[:len(height)-1])
	}

	return countZeros(height) + trap(subtractOne(height))
}

func subtractOne(height []int) []int {
	if len(height) == 0 {
		return []int{}
	}

	if height[0] > 0 {
		return append([]int{height[0] - 1}, subtractOne(height[1:])...)
	}

	return append([]int{height[0]}, subtractOne(height[1:])...)
}

func countZeros(height []int) int {
	if len(height) == 0 {
		return 0
	}

	if height[0] == 0 {
		return 1 + countZeros(height[1:])
	}

	return countZeros(height[1:])
}
