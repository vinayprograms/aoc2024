package main

import (
	"fmt"
	"strings"
)

func one(lines []string) {
	runningTotal := 0
	for _, l := range lines {
		//l := lines[len(lines)-1]
		parts := strings.Split(strings.TrimSpace(l), ":")
		items := strings.Split(strings.TrimSpace(parts[1]), " ")
		fmt.Println("RESULT:", parts[0], "EQN:", items)
		operatorCount := len(items) - 1
		operators := []string{"+", "*"}
		opsMap := map[string](func(int, int) int){
			"+": func(x, y int) int { return x + y },
			"*": func(x, y int) int { return x * y },
		}
		ops := generateOperators(operators, operatorCount)
		equations := assemble(items, ops)
		if isAMatch, result := check(parts[0], equations, opsMap); isAMatch {
			runningTotal += result
		}
	}
	fmt.Println(runningTotal)
}
