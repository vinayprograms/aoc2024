package main

import (
	"fmt"
	"strconv"
	"strings"
)

func two(lines []string) {
	runningTotal := 0
	for _, l := range lines {
		parts := strings.Split(strings.TrimSpace(l), ":")
		items := strings.Split(strings.TrimSpace(parts[1]), " ")
		fmt.Println("RESULT:", parts[0], "EQN:", items)
		operatorCount := len(items) - 1
		operators := []string{"+", "*", "||"}
		opsMap := map[string](func(int, int) int){
			"+": func(x, y int) int { return x + y },
			"*": func(x, y int) int { return x * y },
			"||": func(x, y int) int {
				xStr := strconv.Itoa(x)
				yStr := strconv.Itoa(y)
				result := fmt.Sprintf("%s%s", xStr, yStr)
				if r, err := strconv.Atoi(result); err != nil {
					panic(err)
				} else {
					return r
				}
			},
		}
		ops := generateOperators(operators, operatorCount)
		equations := assemble(items, ops)
		if isAMatch, result := check(parts[0], equations, opsMap); isAMatch {
			runningTotal += result
		}
	}
	fmt.Println(runningTotal)
}
