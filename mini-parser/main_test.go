package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNestedIntegerToArray(t *testing.T) {
	ni1 := &NestedInteger{}
	ni2 := &NestedInteger{}
	ni3 := &NestedInteger{}
	ni4 := &NestedInteger{}
	ni5 := &NestedInteger{}
	ni6 := &NestedInteger{}

	ni1.SetInteger(123)
	ni2.SetInteger(456)
	ni3.SetInteger(789)

	ni4.Add(*ni1)
	ni5.Add(*ni2)
	ni6.Add(*ni3)

	ni5.Add(*ni6)
	ni4.Add(*ni5)

	expected1 := 123
	expected2 := []any{
		123,
		[]any{
			456,
			[]any{
				789,
			},
		},
	}

	if NestedIntegerToArray(ni1) != expected1 {
		t.Errorf("Example 1: Expected %v but got %v", expected1, NestedIntegerToArray(ni1))
	}

	if !reflect.DeepEqual(NestedIntegerToArray(ni4), expected2) {
		t.Errorf("Example 2: Expected %v but got %v", expected2, NestedIntegerToArray(ni4))
	}
}

func TestDeserialize(t *testing.T) {
	type TestCase struct {
		input    string
		expected any
	}

	testCases := []TestCase{
		{
			input:    "123",
			expected: 123,
		},
		{
			input:    "[123,[456,[789]]]",
			expected: []any{123, []any{456, []any{789}}},
		},
	}

	testFunc := func(tc TestCase) func(*testing.T) {
		return func(t *testing.T) {
			got := NestedIntegerToArray(deserialize(tc.input))
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, got)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i+1), testFunc(tc))
	}
}
