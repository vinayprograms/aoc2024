package main

import (
	"fmt"
	"global"
	"strings"
)

func two(lines []string) {
	text := merge(lines)
	if len(strings.TrimSpace(text)) == 0 {
		return
	}
	tokens := tokenize(text, `mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	filtered := filter(tokens, func(t []string) []string {
		include := true
		filtered_tokens := []string{}
		for _, x := range t {
			switch {
			case x == "don't()":
				include = false
			case x == "do()":
				include = true
			case include == true:
				filtered_tokens = append(filtered_tokens, x)
			}
		}
		return filtered_tokens
	})
	results := calculate(filtered)
	fmt.Println("RESULT:", global.Sum(results))
}
