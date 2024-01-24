package main

const (
	OpenParenthesis  = '('
	CloseParenthesis = ')'
	OpenBracket      = '['
	CloseBracket     = ']'
	OpenBrace        = '{'
	CloseBrace       = '}'
)

type Stack []rune

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() *rune {
	if len(*s) == 0 {
		return nil
	}

	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return &r
}

func IsValid(s string) bool {
	var stack Stack

	for _, r := range s {
		if r == OpenParenthesis || r == OpenBracket || r == OpenBrace {
			stack.Push(r)
			continue
		}
		head := stack.Pop()
		if head == nil {
			return false
		}
		if r == CloseParenthesis && *head != OpenParenthesis {
			return false
		}
		if r == CloseBracket && *head != OpenBracket {
			return false
		}
		if r == CloseBrace && *head != OpenBrace {
			return false
		}
	}

	return len(stack) == 0
}
