package main

import (
	"testing"
)

func TestAddTwoNumber(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		res := ToInt(AddTwoNumbers(MakeList([]int{2, 4, 3}), MakeList([]int{5, 6, 4}), 0))
		if res != 807 {
			t.Errorf("value shoud be 807 but is %d!", res)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		res := ToInt(AddTwoNumbers(MakeList([]int{0}), MakeList([]int{0}), 0))
		if res != 0 {
			t.Errorf("value shoud be 0 but is %d!", res)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		res := ToInt(AddTwoNumbers(MakeList([]int{9, 9, 9, 9, 9, 9, 9}), MakeList([]int{9, 9, 9, 9}), 0))
		if res != 10009998 {
			t.Errorf("value shoud be 10009998 but is %d!", res)
		}
	})
}

func TestIterativeAddTwoNumber(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		res := ToInt(IterativeAddTwoNumbers(MakeList([]int{2, 4, 3}), MakeList([]int{5, 6, 4}), 0))
		if res != 807 {
			t.Errorf("value shoud be 807 but is %d!", res)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		res := ToInt(IterativeAddTwoNumbers(MakeList([]int{0}), MakeList([]int{0}), 0))
		if res != 0 {
			t.Errorf("value shoud be 0 but is %d!", res)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		res := ToInt(IterativeAddTwoNumbers(MakeList([]int{9, 9, 9, 9, 9, 9, 9}), MakeList([]int{9, 9, 9, 9}), 0))
		if res != 10009998 {
			t.Errorf("value shoud be 10009998 but is %d!", res)
		}
	})
}

func TestMakeList(t *testing.T) {
	res := ToInt(MakeList([]int{7, 0, 8}))
	if res != 807 {
		t.Errorf("value shoud be 807 but is %d!", res)
	}
}

func TestToInt(t *testing.T) {
	res := ToInt(MakeList([]int{7, 0, 8}))
	if res != 807 {
		t.Errorf("value shoud be 807 but is %d!", res)
	}
}

func TestListsEqual(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		a := MakeList([]int{2, 4, 8})
		b := MakeList([]int{2, 4, 8})
		if ListsEqual(a, b) != true {
			t.Error("lists should be equal")
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		a := MakeList([]int{2, 4, 8})
		b := MakeList([]int{2, 0, 8})
		if ListsEqual(a, b) != false {
			t.Error("lists should not be equal")
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		a := MakeList([]int{2, 4, 8, 1, 1})
		b := MakeList([]int{2, 4, 8})
		if ListsEqual(a, b) != false {
			t.Error("lists should not be equal")
		}
	})

	t.Run("Example 4", func(t *testing.T) {
		a := MakeList([]int{2, 4, 8})
		b := MakeList([]int{2, 4, 8, 1, 1})
		if ListsEqual(a, b) != false {
			t.Error("lists should not be equal")
		}
	})
}
