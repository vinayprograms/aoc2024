package main

import (
	"fmt"
)

func one(lines []string) {
	text := merge(lines)
	tokens := tokenize(text, `mul\(\d+,\d+\)`)
	result := calculate(tokens)
	fmt.Println("RESULT:", result)
}
