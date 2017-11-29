package main

import (
	"unicode"
)

func RemoveEven(input []int) []int {
	slice := make([]int, 0)
	for i := range input {
		if input[i]%2 == 1 {
			slice = append(slice, input[i])
		}
	}
	return slice
}

func DifferentWordsCount(s string) int {
	allWords := make([]string, 1)
	var j = 0
	var isNotWord = false
	for i:= range s {
		if unicode.IsLetter(rune(s[i])){
			allWords[j] = allWords[j] + string(unicode.ToLower(rune(s[i])))
			isNotWord = false
		} else if !isNotWord {
			isNotWord = true
			allWords = append(allWords, "")
			j = j + 1
		}
	}
	result := make(map[string]int)
	for i := 0; i < len(allWords); i++ {
		if allWords[i] != "" {
			result[allWords[i]]++
		}
	}
	return len(result)
}

func PowerGenerator(a int) func() int{
	b := a
	return func() (pow int) {
		pow = b
		b = b * a
		return
	}
}
