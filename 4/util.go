package main

import (
	"fmt"
)

func findStarts(m Map[string], start string) []Point {
	starts := []Point{}
	for p, v := range m.positions {
		if v == start {
			// convert string "[x, y]" to Point{coordinates: [x, y]}
			var x, y int
			fmt.Sscanf(p, "[%d %d]", &x, &y)
			starts = append(starts, Point{coordinates: []int{x, y}})
		}
	}

	return starts
}

func moveAndFind(query []string, queryPos int, m Map[string], start Point, direction Vector) bool {
	if m.ValueAt(start) == query[queryPos] {
		queryPos++
		if queryPos == len(query) {
			return true
		}
		next, err := m.Move(start, direction)
		if err != nil {
			return false
		} else {
			return moveAndFind(query, queryPos, m, next, direction)
		}
	}
	return false
}
