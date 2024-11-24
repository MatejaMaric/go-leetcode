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

	testFunc := func(tc TestCase, trap func([]int) int) func(*testing.T) {
		return func(t *testing.T) {
			if actual := trap(tc.height); actual != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, actual)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Trap V1: Example %d", i), testFunc(tc, trap_V1))
		t.Run(fmt.Sprintf("Trap V2: Example %d", i), testFunc(tc, trap_V2))
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
				t.Errorf("Array subtracted correctly: Expected: %v Got: %v", tc.expected, res)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i), testFunc(tc))
	}
}
