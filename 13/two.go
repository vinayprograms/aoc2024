package main

import "fmt"

func two(lines []string) {
	total := 0
	fixed := fixGames(getGames(lines))

	// For puzzle 2, we'll use Euler's formula to solve simultaneous equations
	for _, g := range fixed {
		matrix := [2][2]int{
			{g.a[0], g.b[0]},
			{g.a[1], g.b[1]},
		}
		inverse := invert(matrix)
		var x, y int
		remainder := ((inverse[0][0] * g.prize[0]) + (inverse[0][1] * g.prize[1])) % determinant(matrix)
		if remainder == 0 {
			x = ((inverse[0][0] * g.prize[0]) + (inverse[0][1] * g.prize[1])) / determinant(matrix)
		} else {
			//fmt.Println("No solution:", g)
			continue
		}
		remainder = ((inverse[1][0] * g.prize[0]) + (inverse[1][1] * g.prize[1])) % determinant(matrix)
		if remainder == 0 {
			y = ((inverse[1][0] * g.prize[0]) + (inverse[1][1] * g.prize[1])) / determinant(matrix)
		} else {
			//fmt.Println("No Solution:", g)
			continue
		}
		//fmt.Println("SOLUTION:", g, "->", x, ",", y)
		total += (x * 3) + y
	}
	fmt.Println("TOTAL COST:", total)
}

func determinant(matrix [2][2]int) int {
	return (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
}

func invert(matrix [2][2]int) [2][2]int {
	inv := [2][2]int{}
	inv[0][0] = matrix[1][1]
	inv[1][1] = matrix[0][0]
	inv[0][1] = -matrix[0][1]
	inv[1][0] = -matrix[1][0]
	return inv
}
