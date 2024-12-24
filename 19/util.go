package main

import (
	"fmt"
	"strings"
)

func prepInput(lines []string) ([]string, []string) {
	towels := prepTowels(lines[0])
	designs := []string{}
	for i := 2; i < len(lines); i++ {
		designs = append(designs, strings.TrimSpace(lines[i]))
	}
	return towels, designs
}

func prepTowels(line string) []string {
	towels := []string{}
	parts := strings.Split(line, ",")
	for _, part := range parts {
		s := strings.TrimSpace(part)
		towels = append(towels, s)
	}
	return towels
}

// Function accepts a set of available patterns and the design that needs
// to be built out of those patterns. It returns the number of ways such a
// pattern can be broken into sub-patterns sourced from the input list of patterns. A cache structure is passed around to speed up search. Cache holds the number of ways a specific substring was split already.
func next(towels []string, design string, cache *map[string]int) int {
	// if remaining string was already processed, re-processing it will
	// give the same result. Save time by pulling it from cache.
	if v, ok := (*cache)[design]; ok {
		return v
	}

	// Degenerate case: When rest of the string was not found in the cache,
	// and that string is empty, return 1 to indicate that one instance of
	// the most the recently matched pattern has occurred once.
	if len(design) == 0 {
		(*cache)[design] = 1
		return 1
	}

	// Out of all the available patterns, only pick those that are substrings
	// of the current string (may appear at the beginning, middle or end). As
	// we continue breaking up the string, the matches list keeps narrowing
	// since we don't need to search for patterns that don't appear anymore
	// in current substring.
	matches := []string{}
	for _, towel := range towels {
		if strings.Contains(design, towel) {
			matches = append(matches, towel)
		}
	}

	// If no patterns matched current string, we cannot continue. Record that
	// in the cache and exit
	if len(matches) == 0 {
		(*cache)[design] = 0
		return 0
	}

	// Narrow matched substrings list to the ones that mark the beginning
	// of the string.
	begins := []string{}
	for _, towel := range matches {
		if strings.HasPrefix(design, towel) {
			begins = append(begins, towel)
		}
	}

	// If no beginning substring was found, it means there is no pattern
	// that can fulfill the requirement. We don't have to continue sub-string
	// match search on rest of the string.
	if len(begins) == 0 {
		fmt.Println("ADDING MISSING TOWEL", design, "TO CACHE")
		(*cache)[design] = 0
		return 0
	}
	// Iterate through the list of beginning substrings
	count := 0
	for _, b := range begins {
		// `matchCount` holds the number of ways we could break up the
		// current string. Collect this count for each beginning match
		// we have.
		count += next(matches, strings.TrimPrefix(design, b), cache)
	}
	(*cache)[design] = count
	fmt.Println("COUNT TILL NOW:", count, "FOR", design)

	return count
}
