package main

import (
	"testing"
)

func TestAddTwoNumber(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		res := toInt(addTwoNumbers(makeList([]int{2, 4, 3}), makeList([]int{5, 6, 4}), 0))
		if res != 807 {
			t.Errorf("value shoud be 807 but is %d!", res)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		res := toInt(addTwoNumbers(makeList([]int{0}), makeList([]int{0}), 0))
		if res != 0 {
			t.Errorf("value shoud be 0 but is %d!", res)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		res := toInt(addTwoNumbers(makeList([]int{9, 9, 9, 9, 9, 9, 9}), makeList([]int{9, 9, 9, 9}), 0))
		if res != 10009998 {
			t.Errorf("value shoud be 10009998 but is %d!", res)
		}
	})
}

func TestMakeList(t *testing.T) {
	res := toInt(makeList([]int{7, 0, 8}))
	if res != 807 {
		t.Errorf("value shoud be 807 but is %d!", res)
	}
}

func TestToInt(t *testing.T) {
	res := toInt(makeList([]int{7, 0, 8}))
	if res != 807 {
		t.Errorf("value shoud be 807 but is %d!", res)
	}
}
