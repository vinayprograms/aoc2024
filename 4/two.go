package main

import "fmt"

func two(lines []string) {
	values := []string{}
	for _, l := range lines {
		for _, c := range l {
			values = append(values, string(c))
		}
	}

	m := Map[string]{}
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

func findX_MAS(start Point, m Map[string]) bool {
	offsets := [4]Vector{
		Vector{deltas: []int{-1, -1}}, // "M"
		Vector{deltas: []int{-1, 1}},  // "S"
		Vector{deltas: []int{1, 1}},   // "S"
		Vector{deltas: []int{1, -1}},  // "M"
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

func matchPositions(start Point, m Map[string], offsets [4]Vector) bool {
	if match(start, offsets[0], m, "M") &&
		match(start, offsets[1], m, "S") &&
		match(start, offsets[2], m, "S") &&
		match(start, offsets[3], m, "M") {
		return true
	} else {
		return false
	}
}

func match(start Point, offset Vector, m Map[string], str string) bool {
	if pos, err := m.Move(start, offset); err != nil {
		return false
	} else if m.ValueAt(pos) == str {
		return true
	} else {
		return false
	}
}

func rotateGrid(v [4]Vector) [4]Vector {
	return [4]Vector{
		rotate90(v[0]),
		rotate90(v[1]),
		rotate90(v[2]),
		rotate90(v[3]),
	}
}

// Rotate vector clockwise by 90 degrees
func rotate90(v Vector) Vector {
	if len(v.deltas) != 2 {
		panic("rotate90 only works on 2D vectors")
	}
	return Vector{deltas: []int{
		(0 * v.deltas[0]) + (-1 * v.deltas[1]),
		(1 * v.deltas[0]) + (0 * v.deltas[1]),
	}}
}
