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

// ------------------------------------------------

func trap_V3(height []int) int {
	left := maxLeftArr(height)
	right := maxRightArr(height)
	lrmin := minLR(left, right)

	return ReduceTwoIntArrays(height, lrmin, func(height int, lrmin int) int {
		if lrmin > height {
			return lrmin - height
		}
		return 0
	})
}

func ReduceTwoIntArrays(arr1, arr2 []int, callback func(int, int) int) int {
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}
	return callback(arr1[0], arr2[0]) + ReduceTwoIntArrays(arr1[1:], arr2[1:], callback)
}

func minLR(left, right []int) []int {
	if len(left) == 0 || len(right) == 0 {
		return []int{}
	}

	if left[0] < right[0] {
		return append([]int{left[0]}, minLR(left[1:], right[1:])...)
	}

	return append([]int{right[0]}, minLR(left[1:], right[1:])...)
}

func maxLeftArr(height []int) []int {
	if len(height) == 0 {
		return []int{}
	}

	maxLeft := 0
	newMax := height[len(height)-1]

	arr := maxLeftArr(height[:len(height)-1])
	if len(arr) > 0 {
		maxLeft = arr[len(arr)-1]
	}

	if newMax < maxLeft {
		return append(arr, maxLeft)
	}

	return append(arr, newMax)
}

func maxRightArr(height []int) []int {
	if len(height) == 0 {
		return []int{}
	}

	maxRight := 0
	newMax := height[0]

	arr := maxRightArr(height[1:])
	if len(arr) > 0 {
		maxRight = arr[0]
	}

	if newMax < maxRight {
		return append([]int{maxRight}, arr...)
	}

	return append([]int{newMax}, arr...)
}

// ------------------------------------------------

func trap_V4(height []int) int {
	sum := 0
	lrmin := iterativeMinLR(iterativeMaxLeftArr(height), iterativeMaxRightArr(height))

	for i := 0; i < len(height); i++ {
		if lrmin[i] > height[i] {
			sum += lrmin[i] - height[i]
		}
	}

	return sum
}

func iterativeMinLR(left, right []int) []int {
	if len(left) != len(right) {
		return []int{}
	}

	res := make([]int, len(left))
	for i := 0; i < len(left); i++ {
		if left[i] < right[i] {
			res[i] = left[i]
		} else {
			res[i] = right[i]
		}
	}

	return res
}

func iterativeMaxLeftArr(height []int) []int {
	arr := make([]int, 0, len(height))
	prevMax := 0

	for _, v := range height {
		if v > prevMax {
			arr = append(arr, v)
			prevMax = v
		} else {
			arr = append(arr, prevMax)
		}
	}

	return arr
}

func iterativeMaxRightArr(height []int) []int {
	arr := make([]int, len(height))
	prevMax := 0

	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > prevMax {
			arr[i] = height[i]
			prevMax = height[i]
		} else {
			arr[i] = prevMax
		}
	}

	return arr
}

// ------------------------------------------------

func trap_V5(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	res := 0

	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			res += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			res += rightMax - height[right]
		}
	}

	return res
}

// ------------------------------------------------

func trap_V6(height []int) int {
	if len(height) == 0 {
		return 0
	}
	return rec_trap(height, height[0], height[len(height)-1])
}

func rec_trap(height []int, leftMax, rightMax int) int {
	if len(height) == 0 {
		return 0
	}

	if leftMax < rightMax {
		newLeftMax := max(leftMax, height[0])
		trapped := (newLeftMax - height[0])
		return trapped + rec_trap(height[1:], newLeftMax, rightMax)
	}

	newRightMax := max(rightMax, height[len(height)-1])
	trapped := (newRightMax - height[len(height)-1])
	return trapped + rec_trap(height[:len(height)-1], leftMax, newRightMax)
}
