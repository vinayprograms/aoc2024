package main

import (
	"fmt"
	"global"
)

func one(lines []string) {
	text := merge(lines)
	tokens := tokenize(text, `mul\(\d+,\d+\)`)
	results := calculate(tokens)
	fmt.Println("RESULT:", global.Sum(results))
}
