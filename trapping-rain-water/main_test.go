package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestTrap(t *testing.T) {
	type TestCase struct {
		height   []int
		expected int
	}

	testCases := []TestCase{
		{
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
	}

	implementations := []func([]int) int{
		trap_V1,
		trap_V2,
		trap_V3,
		trap_V4,
	}

	testFunc := func(tc TestCase, trap func([]int) int) func(*testing.T) {
		return func(t *testing.T) {
			if actual := trap(tc.height); actual != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, actual)
			}
		}
	}

	testImpl := func(impl func([]int) int) func(*testing.T) {
		return func(t *testing.T) {
			for i, tc := range testCases {
				t.Run(fmt.Sprintf("Example %d", i), testFunc(tc, impl))
			}
		}
	}

	for i, impl := range implementations {
		t.Run(fmt.Sprintf("Trap V%d", i), testImpl(impl))
	}
}

func TestTrimZeros(t *testing.T) {
	type TestCase struct {
		arr      []int
		expected []int
	}

	testCases := []TestCase{
		{
			arr:      []int{},
			expected: []int{},
		},
		{
			arr:      []int{0},
			expected: []int{},
		},
		{
			arr:      []int{1},
			expected: []int{1},
		},
		{
			arr:      []int{0, 0, 0, 0, 0},
			expected: []int{},
		},
		{
			arr:      []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			arr:      []int{0, 0, 1, 0, 1, 0, 0},
			expected: []int{1, 0, 1},
		},
		{
			arr:      []int{1, 0, 1, 0, 0},
			expected: []int{1, 0, 1},
		},
		{
			arr:      []int{0, 0, 1, 0, 1},
			expected: []int{1, 0, 1},
		},
	}

	testFunc := func(tc TestCase) func(*testing.T) {
		return func(t *testing.T) {
			if res := trimZeros(tc.arr); !slices.Equal(res, tc.expected) {
				t.Errorf("Array not trimmed correctly: Expected: %v Got: %v", tc.expected, res)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i), testFunc(tc))
	}
}

func TestCountZeros(t *testing.T) {
	type TestCase struct {
		arr      []int
		expected int
	}
	testCases := []TestCase{
		{
			arr:      []int{},
			expected: 0,
		},
		{
			arr:      []int{0},
			expected: 1,
		},
		{
			arr:      []int{1},
			expected: 0,
		},
		{
			arr:      []int{0, 0, 0, 0, 0},
			expected: 5,
		},
		{
			arr:      []int{1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			arr:      []int{0, 0, 1, 0, 1, 0, 0},
			expected: 5,
		},
		{
			arr:      []int{1, 0, 1, 0, 0},
			expected: 3,
		},
		{
			arr:      []int{0, 0, 1, 0, 1},
			expected: 3,
		},
	}
	testFunc := func(tc TestCase) func(*testing.T) {
		return func(t *testing.T) {
			if res := countZeros(tc.arr); res != tc.expected {
				t.Errorf("Count not correct: Expected: %d Got: %d", tc.expected, res)
			}
		}
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i), testFunc(tc))
	}
}

func TestSubtractOne(t *testing.T) {
	type TestCase struct {
		arr      []int
		expected []int
	}

	testCases := []TestCase{
		{
			arr:      []int{},
			expected: []int{},
		},
		{
			arr:      []int{0},
			expected: []int{0},
		},
		{
			arr:      []int{1},
			expected: []int{0},
		},
		{
			arr:      []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			arr:      []int{1, 1, 1, 1, 1},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			arr:      []int{0, 0, 5, 0, 4, 0, 0},
			expected: []int{0, 0, 4, 0, 3, 0, 0},
		},
		{
			arr:      []int{1, 0, 1, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			arr:      []int{3, 0, 6},
			expected: []int{2, 0, 5},
		},
	}

	testFunc := func(tc TestCase) func(*testing.T) {
		return func(t *testing.T) {
			if res := subtractOne(tc.arr); !slices.Equal(res, tc.expected) {
				t.Errorf("Array subtracted incorrectly: Expected: %v Got: %v", tc.expected, res)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i), testFunc(tc))
	}
}

func TestMaxLeftArr(t *testing.T) {
	type TestCase struct {
		height   []int
		expected []int
	}

	testCases := []TestCase{
		{
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: []int{0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3},
		},
		{
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: []int{4, 4, 4, 4, 4, 5},
		},
	}

	testFunc := func(tc TestCase, maxLeftFunc func([]int) []int) func(*testing.T) {
		return func(t *testing.T) {
			if res := maxLeftFunc(tc.height); !slices.Equal(res, tc.expected) {
				t.Errorf("Expected: %v Got: %v", tc.expected, res)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Recursive maxLeftArr: Example %d", i), testFunc(tc, maxLeftArr))
		t.Run(fmt.Sprintf("Iterative maxLeftArr: Example %d", i), testFunc(tc, iterativeMaxLeftArr))
	}
}

func TestMaxRightArr(t *testing.T) {
	type TestCase struct {
		height   []int
		expected []int
	}

	testCases := []TestCase{
		{
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: []int{3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 1},
		},
		{
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: []int{5, 5, 5, 5, 5, 5},
		},
	}

	testFunc := func(tc TestCase, maxRightFunc func([]int) []int) func(*testing.T) {
		return func(t *testing.T) {
			if res := maxRightFunc(tc.height); !slices.Equal(res, tc.expected) {
				t.Errorf("Expected: %v Got: %v", tc.expected, res)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Recursive maxRightArr: Example %d", i), testFunc(tc, maxRightArr))
		t.Run(fmt.Sprintf("Iterative maxRightArr: Example %d", i), testFunc(tc, iterativeMaxRightArr))
	}
}

func TestMinLR(t *testing.T) {
	type TestCase struct {
		left     []int
		right    []int
		expected []int
	}

	testCases := []TestCase{
		{
			left:     []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			right:    []int{3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 1},
			expected: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		},
		{
			left:     []int{4, 2, 0, 3, 2, 5},
			right:    []int{5, 5, 5, 5, 5, 5},
			expected: []int{4, 2, 0, 3, 2, 5},
		},
	}

	testFunc := func(tc TestCase, minLRFunc func([]int, []int) []int) func(*testing.T) {
		return func(t *testing.T) {
			if res := minLRFunc(tc.left, tc.right); !slices.Equal(res, tc.expected) {
				t.Errorf("Expected: %v Got: %v", tc.expected, res)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Recursive minLR: Example %d", i), testFunc(tc, minLR))
		t.Run(fmt.Sprintf("Iterative minLR: Example %d", i), testFunc(tc, iterativeMinLR))
	}
}
