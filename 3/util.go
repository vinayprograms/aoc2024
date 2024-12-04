package main

import (
	"regexp"
	"strconv"
)

// Merge lines into a single piece of text
func merge(lines []string) string {
	text := ""
	for _, str := range lines {
		text += str
	}
	return text
}

// Uses the supplied regular expression to extract a set of phrases from input text.
func tokenize(text string, regex string) []string {
	re := regexp.MustCompile(regex)
	return re.FindAllString(text, -1)
}

// Applies the supplied filtering function to the set of provided tokens.
// Filtering function is expected to remove tokens based on some specific logic.
func filter(tokens []string, filterFunc func([]string) []string) []string {
	return filterFunc(tokens)
}

// Calculate the total by processing 'mul(a,b)' instructions
func calculate(instructions []string) []int {
	results := []int{}
	for _, instr := range instructions {
		re := regexp.MustCompile(`\d+`)
		numbers := re.FindAllString(instr, 2) // instruction must contain only two numbers
		if a, err := strconv.Atoi(numbers[0]); err != nil {
			panic(err)
		} else if b, err := strconv.Atoi(numbers[1]); err != nil {
			panic(err)
		} else {
			results = append(results, a*b)
		}
	}
	return results
}
