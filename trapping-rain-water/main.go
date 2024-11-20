package main

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	trimedHeight := trimZeros(height)

	return countZeros(trimedHeight) + trap(subtractOne(trimedHeight))
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
