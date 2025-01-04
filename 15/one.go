package main

import "fmt"

func one(lines []string) {
	g := input(lines)
	fmt.Println(g)
	state := g.warehouse
	position := g.robot
	for _, direction := range g.moves {
		switch direction {
		case '<':
			state, position = left(state, position)
		case '>':
			state, position = right(state, position)
		case '^':
			state, position = up(state, position)
		case 'v':
			state, position = down(state, position)
		}
	}
	for _, line := range state {
		fmt.Println(line)
	}
	fmt.Println(gpsSum(state))
}

func gpsSum(warehouse [][]string) int {
	sum := 0
	for y := 0; y < len(warehouse); y++ {
		for x := 0; x < len(warehouse[y]); x++ {
			if warehouse[y][x] == "O" {
				sum += (100 * y) + x
			}
		}
	}
	return sum
}
