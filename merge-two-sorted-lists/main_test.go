package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type TestCase struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}

	testCases := []TestCase{
		{
			list1:    arrToList([]int{1, 2, 4}),
			list2:    arrToList([]int{1, 3, 4}),
			expected: arrToList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			list1:    arrToList([]int{}),
			list2:    arrToList([]int{}),
			expected: arrToList([]int{}),
		},
		{
			list1:    arrToList([]int{}),
			list2:    arrToList([]int{0}),
			expected: arrToList([]int{0}),
		},
	}

	testFunc := func(tc TestCase) func(*testing.T) {
		return func(t *testing.T) {
			if !reflect.DeepEqual(mergeTwoLists(tc.list1, tc.list2), tc.expected) {
				t.Error("not equal")
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i+1), testFunc(tc))
	}
}

func TestArrToList(t *testing.T) {
	type TestCase struct {
		arr      []int
		expected *ListNode
	}

	testCases := []TestCase{
		{
			arr:      []int{1, 2, 3},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			arr:      []int{0, 4, 9},
			expected: &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}},
		},
	}

	testFunc := func(tc TestCase) func(*testing.T) {
		return func(t *testing.T) {
			if !reflect.DeepEqual(arrToList(tc.arr), tc.expected) {
				t.Error("not equal")
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", i+1), testFunc(tc))
	}
}
