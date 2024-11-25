package main

import (
	"strconv"
	"strings"
	"unicode"
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

	if unicode.IsDigit(rune(s[0])) {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		res := &NestedInteger{}
		res.SetInteger(num)
		return res
	}

	if s[0] == '[' && s[len(s)-1] == ']' {
		trimmed := s[1 : len(s)-1]

		firstBracket := len(trimmed)
		if index := strings.Index(trimmed, "["); index > -1 {
			firstBracket = index
		}

		res := &NestedInteger{}
		for _, v := range strings.Split(trimmed[:firstBracket], ",") {
			if parsed := deserialize(v); parsed != nil {
				res.Add(*parsed)
			}
		}
		if parsed := deserialize(trimmed[firstBracket:]); parsed != nil {
			res.Add(*parsed)
		}
		return res
	}

	return nil
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
