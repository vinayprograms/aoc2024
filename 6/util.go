package main

import (
	"fmt"
	g "global"
)

func find(m g.Map[string], sprite string) g.Point {
	start := g.Point{Coordinates: []int{-1, -1}} // Start with invalid position
	for p, v := range m.Positions {
		if v == sprite {
			// convert string "[x, y]" to Point{coordinates: [x, y]}
			var x, y int
			fmt.Sscanf(p, "[%d %d]", &x, &y)
			start = g.Point{Coordinates: []int{x, y}}
			break
		}
	}

	return start
}

func moveOne(from g.Point, direction g.Vector) (g.Point, error) {
	return g.Move(from, direction)
}

func hasObstruction(m g.Map[string], p g.Point) bool {
	if v, err := m.ValueAt(p); err == nil && v == "#" {
		return true
	} else {
		return false
	}
}

// Rotate vector clockwise by 90 degrees
func turn(v g.Vector) g.Vector {
	if len(v.Deltas) != 2 {
		panic("rotate90 only works on 2D vectors")
	}
	return g.Vector{Deltas: []int{
		(0 * v.Deltas[0]) + (-1 * v.Deltas[1]),
		(1 * v.Deltas[0]) + (0 * v.Deltas[1]),
	}}
}
