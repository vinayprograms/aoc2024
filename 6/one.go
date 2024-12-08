package main

import (
	"fmt"
	g "global"
)

func one(lines []string) {
	// Extract each position into a single collection
	values := []string{}
	for _, l := range lines {
		for _, c := range l {
			values = append(values, string(c))
		}
	}

	m := g.Map[string]{}
	dimensions := []int{len(lines[0]), len(lines)}
	m.Build(dimensions, values)

	start := find(m, "^")
	if !m.IsInsideMap(start) {
		fmt.Println("Guard not found! You are free to explore :)")
	} else {
		fmt.Println(start, "^")
	}

	current := start
	direction := g.Vector{Deltas: []int{0, -1}}
	fmt.Print(current, direction, "; ")
	count := 0
	for {
		newPos, err := moveOne(current, direction)
		if err != nil || m.IsInsideMap(newPos) == false {
			fmt.Println("\n!! GUARD HAS LEFT THE BUILDING !!")
			break
		}
		if hasObstruction(m, newPos) {
			direction = turn(direction)
			fmt.Println(current, direction, "; ")
		} else {
			current = newPos
			fmt.Print(current, direction, "; ")
			if m.Positions[current.String()] != "X" {
				count++ // only capture distinctly visited positions
			}
			m.Positions[current.String()] = "X"
		}
	}
	fmt.Println("TOTAL LOCATIONS VISITED:", count)
}
