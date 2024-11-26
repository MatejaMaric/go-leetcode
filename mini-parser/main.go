package main

import (
	"strconv"
)

type NestedInteger struct {
	isInt bool
	value int
	list  []*NestedInteger
}

func (n NestedInteger) IsInteger() bool {
	return n.isInt
}

func (n NestedInteger) GetInteger() int {
	return n.value
}

func (n *NestedInteger) SetInteger(value int) {
	n.isInt = true
	n.value = value
}

func (n *NestedInteger) Add(elem NestedInteger) {
	n.isInt = false
	n.list = append(n.list, &elem)
}

func (n NestedInteger) GetList() []*NestedInteger {
	if n.isInt {
		return []*NestedInteger{}
	}

	return n.list
}

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}

	if (s[0] >= '0' && s[0] <= '9') || s[0] == '-' {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		res := &NestedInteger{}
		res.SetInteger(num)
		return res
	}

	if s[0] == '[' && s[len(s)-1] == ']' {
		res := &NestedInteger{}
		for _, str := range breakdown(s[1 : len(s)-1]) {
			if parsed := deserialize(str); parsed != nil {
				res.Add(*parsed)
			}
		}
		return res
	}

	return nil
}

func breakdown(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	depth := 0
	start := 0

	res := []string{}
	for i, v := range s {
		if v == '[' {
			depth++
		}
		if v == ']' {
			depth--
		}
		if v == ',' && depth == 0 {
			res = append(res, s[start:i])
			start = i + 1
		}
	}
	res = append(res, s[start:])

	return res
}

func rec_breakdown(s string, index int, depth int) []string {
	if len(s) == 0 {
		return []string{}
	}

	if index >= len(s)-1 {
		return []string{s}
	}

	if s[index] == '[' {
		return rec_breakdown(s, index+1, depth+1)
	}

	if s[index] == ']' {
		return rec_breakdown(s, index+1, depth-1)
	}

	if s[index] == ',' && depth == 0 {
		return append([]string{s[:index]}, rec_breakdown(s[index+1:], 0, 0)...)
	}

	return rec_breakdown(s, index+1, depth)
}

func NestedIntegerToArray(n *NestedInteger) any {
	if n == nil {
		return nil
	}

	if n.IsInteger() {
		return n.GetInteger()
	}

	res := []any{}
	for _, v := range n.GetList() {
		res = append(res, NestedIntegerToArray(v))
	}
	return res
}
