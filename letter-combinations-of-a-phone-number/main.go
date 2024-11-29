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

	result := []string{}
	for _, letter := range letters {
		for _, r := range rest {
			result = append(result, letter+r)
		}
	}
	return result
}
