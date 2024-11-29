package main

import "strings"

func letterCombinations(digits string) []string {
	lettersMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	digitsArr := strings.Split(digits, "")

	if len(digitsArr) == 0 {
		return []string{}
	}

	firstDigit := digitsArr[0]
	letters := strings.Split(lettersMap[firstDigit], "")

	if len(digitsArr) == 1 {
		return letters
	}

	rest := letterCombinations(strings.Join(digitsArr[1:], ""))

	return FlatMap(letters, func(letter string) []string {
		return Map(rest, func(r string) string {
			return letter + r
		})
	})
}

func FlatMap[A, B any](input []A, f func(A) []B) []B {
	var result []B
	for _, v := range input {
		result = append(result, f(v)...)
	}
	return result
}

func Map[A, B any](input []A, f func(A) B) []B {
	var result []B
	for _, v := range input {
		result = append(result, f(v))
	}
	return result
}

func FlatMapRecursive[A, B any](input []A, f func(A) []B) []B {
	if len(input) == 0 {
		return []B{}
	}
	return append(f(input[0]), FlatMapRecursive(input[1:], f)...)
}

func MapRecursive[A, B any](input []A, f func(A) B) []B {
	if len(input) == 0 {
		return []B{}
	}
	return append([]B{f(input[0])}, MapRecursive(input[1:], f)...)
}
