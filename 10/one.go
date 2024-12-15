package main

import (
	"fmt"
	"strconv"
)

func one(lines []string) {
	_map := build(lines)
	heads := findPoints(lines, "0")
	//fmt.Println(heads)
	ends := findPoints(lines, "9")
	//fmt.Println(ends)
	headScores := []int{}
	for i := 0; i < len(heads); i++ {
		headScores = append(headScores, findNext(&_map, [2]int{len(lines[0]), len(lines)}, heads[i], 0, check_reachability))
		// reset the "X"s back to "9"s so that other trailheads can use it.
		for _, p := range ends {
			_map[p] = "9"
		}
	}
	fmt.Println("SCORES:", headScores)
	fmt.Println("SUM OF SCORES:", addTrailHeads(headScores))
}

// Check for a single unique path to the target
func check_reachability(_map *map[[2]int]string, bounds [2]int, point [2]int, height int) int {
	if (*_map)[point] == "." {
		return 0
	}

	if (*_map)[point] == strconv.Itoa(height) {
		if (*_map)[point] == "9" {
			(*_map)[point] = "X" // disable the destination to stop other trails from counting this one.
			return 1
		} else {
			return findNext(_map, bounds, point, height, check_reachability)
		}
	}
	return 0
}
