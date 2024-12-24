package main

import (
	"fmt"
)

func two(towels []string, designs []string) {
	// A cache of all unmatched substrings. Avoids wasting time in searching
	cache := make(map[string]int)
	count := 0
	for _, design := range designs {
		fmt.Printf("Searching towels for '%s'\n", design)
		// Recursively look for sequence members
		count += next(towels, design, &cache)
		fmt.Println(count)
	}
	fmt.Println("TOTAL:", count)
}
