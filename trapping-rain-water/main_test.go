package main

import (
	"fmt"
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

	testFunc := func(tc TestCase) func(*testing.T) {
		return func(t *testing.T) {
			if actual := trap(tc.height); actual != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, actual)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i), testFunc(tc))
	}
}
