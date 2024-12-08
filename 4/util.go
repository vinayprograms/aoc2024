package main

import (
	"fmt"
	g "global"
)

func findStarts(m g.Map[string], start string) []g.Point {
	starts := []g.Point{}
	for p, v := range m.Positions {
		if v == start {
			// convert string "[x, y]" to Point{coordinates: [x, y]}
			var x, y int
			fmt.Sscanf(p, "[%d %d]", &x, &y)
			starts = append(starts, g.Point{Coordinates: []int{x, y}})
		}
	}
	return starts
}
