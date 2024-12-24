package main

import (
	"fmt"
)

func one(towels []string, designs []string) {
	var possible int
	// A cache of all unmatched substrings. Avoids wasting time in searching
	cache := make(map[string]int)
	for _, design := range designs {
		fmt.Printf("Searching towels for '%s'\n", design)
		// Recursively look for sequence members
		if next(towels, design, &cache) > 0 {
			possible++
		}
	}
	fmt.Println("POSSIBLE:", possible)
}
