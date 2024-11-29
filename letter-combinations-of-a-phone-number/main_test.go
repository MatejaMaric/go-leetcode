package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type TestCase struct {
		input    string
		expected []string
	}

	testCases := []TestCase{
		{input: "23", expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{input: "", expected: []string{}},
		{input: "2", expected: []string{"a", "b", "c"}},
	}

	testFunc := func(tc TestCase) func(*testing.T) {
		return func(t *testing.T) {
			got := letterCombinations(tc.input)
			if !slices.Equal(got, tc.expected) {
				t.Errorf("got: %v, expected: %v", got, tc.expected)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i+1), testFunc(tc))
	}
}
