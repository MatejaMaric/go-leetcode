package main

// ------------------------------------------------

func trap_V1(height []int) int {
	if len(height) == 0 {
		return 0
	}

	trimedHeight := trimZeros(height)

	return countZeros(trimedHeight) + trap_V1(subtractOne(trimedHeight))
}

func trimZeros(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	if arr[0] == 0 {
		return trimZeros(arr[1:])
	}

	if arr[len(arr)-1] == 0 {
		return trimZeros(arr[:len(arr)-1])
	}

	return arr
}

func countZeros(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if arr[0] == 0 {
		return 1 + countZeros(arr[1:])
	}

	return countZeros(arr[1:])
}

func subtractOne(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	if arr[0] > 0 {
		return append([]int{arr[0] - 1}, subtractOne(arr[1:])...)
	}

	return append([]int{arr[0]}, subtractOne(arr[1:])...)
}

// ------------------------------------------------

func trap_V2(height []int) int {
	sum := 0
	for i := 0; i < len(height); i++ {
		leftMax := trapMax(height[:i])
		rightMax := trapMax(height[i+1:])

		minLR := leftMax
		if rightMax < minLR {
			minLR = rightMax
		}

		if minLR > height[i] {
			sum += minLR - height[i]
		}
	}
	return sum
}

func trapMax(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	return max
}
