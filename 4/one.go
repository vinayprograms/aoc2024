package main

import (
	"fmt"
	g "global"
)

func one(lines []string) {
	values := []string{}
	for _, l := range lines {
		for _, c := range l {
			values = append(values, string(c))
		}
	}

	m := g.Map[string]{}
	dimensions := []int{len(lines[0]), len(lines)}
	m.Build(dimensions, values)
	query := []string{"X", "M", "A", "S"}
	count := 0
	for _, s := range findStarts(m, "X") {
		// Search in a 8 directions
		if moveAndFind(query, 0, m, s, g.Vector{Deltas: []int{0, -1}}) {
			count++
			//fmt.Println("FOUND:", "UP", s, count)
		}
		if moveAndFind(query, 0, m, s, g.Vector{Deltas: []int{1, -1}}) {
			count++
			//fmt.Println("FOUND:", "UP-RIGHT", s, count)
		}
		if moveAndFind(query, 0, m, s, g.Vector{Deltas: []int{1, 0}}) {
			count++
			//fmt.Println("FOUND:", "RIGHT", s, count)
		}
		if moveAndFind(query, 0, m, s, g.Vector{Deltas: []int{1, 1}}) {
			count++
			//fmt.Println("FOUND:", "DOWN-RIGHT", s, count)
		}
		if moveAndFind(query, 0, m, s, g.Vector{Deltas: []int{0, 1}}) {
			count++
			//fmt.Println("FOUND:", "DOWN", s, count)
		}
		if moveAndFind(query, 0, m, s, g.Vector{Deltas: []int{-1, 1}}) {
			count++
			//fmt.Println("FOUND:", "DOWN-LEFT", s, count)
		}
		if moveAndFind(query, 0, m, s, g.Vector{Deltas: []int{-1, 0}}) {
			count++
			//fmt.Println("FOUND:", "LEFT", s, count)
		}
		if moveAndFind(query, 0, m, s, g.Vector{Deltas: []int{-1, -1}}) {
			count++
			//fmt.Println("FOUND:", "UP-LEFT", s, count)
		}
	}
	fmt.Println("COUNT:", count)
}

func moveAndFind(query []string, queryPos int, m g.Map[string], start g.Point, direction g.Vector) bool {
	if v, err := m.ValueAt(start); err == nil && v == query[queryPos] {
		queryPos++
		if queryPos == len(query) {
			return true
		}
		next := m.Move(start, direction)
		if next.Coordinates == nil {
			return false
		} else {
			return moveAndFind(query, queryPos, m, next, direction)
		}
	}
	return false
}
