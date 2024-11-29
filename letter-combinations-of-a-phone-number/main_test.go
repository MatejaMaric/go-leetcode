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

func TestFlatMap(t *testing.T) {
	type TestCase struct {
		input    []int
		expected []int
	}

	type FlatMapFunc func([]int, func(int) []int) []int

	testCases := []TestCase{
		{input: []int{1, 2, 3}, expected: []int{1, 2, 2, 4, 3, 6}},
	}

	testFunc := func(tc TestCase, impl FlatMapFunc) func(*testing.T) {
		return func(t *testing.T) {
			got := impl(tc.input, func(a int) []int {
				return []int{a, a * 2}
			})
			if !slices.Equal(tc.expected, got) {
				t.Errorf("got: %v, expected: %v", got, tc.expected)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i+1), testFunc(tc, FlatMap))
		t.Run(fmt.Sprintf("Recursive example %d", i+1), testFunc(tc, FlatMapRecursive))
	}
}

func TestMap(t *testing.T) {
	type TestCase struct {
		input    []int
		expected []int
	}

	type MapFunc func([]int, func(int) int) []int

	testCases := []TestCase{
		{input: []int{1, 2, 3}, expected: []int{2, 4, 6}},
	}

	testFunc := func(tc TestCase, impl MapFunc) func(*testing.T) {
		return func(t *testing.T) {
			got := impl(tc.input, func(a int) int {
				return a * 2
			})
			if !slices.Equal(tc.expected, got) {
				t.Errorf("got: %v, expected: %v", got, tc.expected)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i+1), testFunc(tc, Map))
		t.Run(fmt.Sprintf("Recursive example %d", i+1), testFunc(tc, MapRecursive))
	}
}
