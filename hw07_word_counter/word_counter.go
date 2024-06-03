package main

import (
	"regexp"
	"strings"
)

var wordsFilterRegex = regexp.MustCompile(`[^\pL\d\s\t\n]+`)

func countWords(input string) map[string]int {
	if len(input) == 0 {
		return make(map[string]int)
	}

	fmtInput := wordsFilterRegex.ReplaceAllString(input, "")

	words := make(map[string]int)
	for _, word := range strings.Fields(fmtInput) {
		words[word]++
	}
	return words
}
