package main

import (
	"fmt"
	g "global"
)

func two(lines []string) {
	values := []string{}
	for _, l := range lines {
		for _, c := range l {
			values = append(values, string(c))
		}
	}

	m := g.Map[string]{}
	dimensions := []int{len(lines[0]), len(lines)}
	m.Build(dimensions, values)
	count := 0
	for _, s := range findStarts(m, "A") {
		if findX_MAS(s, m) {
			count++
			//fmt.Println("AT:", s, "COUNT:", count)
		}
	}
	fmt.Println("COUNT:", count)
}

func findX_MAS(start g.Point, m g.Map[string]) bool {
	offsets := [4]g.Vector{
		g.Vector{Deltas: []int{-1, -1}}, // "M"
		g.Vector{Deltas: []int{-1, 1}},  // "S"
		g.Vector{Deltas: []int{1, 1}},   // "S"
		g.Vector{Deltas: []int{1, -1}},  // "M"
	}
	if matchPositions(start, m, offsets) {
		return true
	}
	offsets = rotateGrid(offsets)
	if matchPositions(start, m, offsets) {
		return true
	}
	offsets = rotateGrid(offsets)
	if matchPositions(start, m, offsets) {
		return true
	}
	offsets = rotateGrid(offsets)
	if matchPositions(start, m, offsets) {
		return true
	}
	return false
}

func matchPositions(start g.Point, m g.Map[string], offsets [4]g.Vector) bool {
	if match(start, offsets[0], m, "M") &&
		match(start, offsets[1], m, "S") &&
		match(start, offsets[2], m, "S") &&
		match(start, offsets[3], m, "M") {
		return true
	} else {
		return false
	}
}

func match(start g.Point, offset g.Vector, m g.Map[string], str string) bool {
	if pos := m.Move(start, offset); len(pos.Coordinates) == 0 {
		return false
	} else {
		if v, err := m.ValueAt(pos); err == nil && v == str {
			return true
		} else {
			return false
		}
	}
}

func rotateGrid(v [4]g.Vector) [4]g.Vector {
	return [4]g.Vector{
		rotate90(v[0]),
		rotate90(v[1]),
		rotate90(v[2]),
		rotate90(v[3]),
	}
}

// Rotate vector clockwise by 90 degrees
func rotate90(v g.Vector) g.Vector {
	if len(v.Deltas) != 2 {
		panic("rotate90 only works on 2D vectors")
	}
	return g.Vector{Deltas: []int{
		(0 * v.Deltas[0]) + (-1 * v.Deltas[1]),
		(1 * v.Deltas[0]) + (0 * v.Deltas[1]),
	}}
}
