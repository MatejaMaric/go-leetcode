package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := IsValid(test.input); got != test.want {
				t.Errorf("for string '%s' expected '%t' got '%t'", test.input, test.want, got)
			}
		})
	}
}

func TestRecursiveIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got := RecursiveIsValid([]rune(test.input), make([]rune, 0))
			if got != test.want {
				t.Errorf("for string '%s' expected '%t' got '%t'", test.input, test.want, got)
			}
		})
	}
}

func TestChatGPT_RecursiveIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()[]{}", true},
		{"([{}])", true},
		{"([)]", false},
		{"((())", false},
		{"((){})", true},
		{"((){}[()]())", true},
		{"{[()]}", true},
		{"{[(])}", false},
		{"", true}, // an empty string is considered valid
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			got := RecursiveIsValid([]rune(test.input), make([]rune, 0))
			if got != test.want {
				t.Errorf("for string '%s' expected '%t' got '%t'", test.input, test.want, got)
			}
		})
	}
}
