package main

import (
	"fmt"
	"strconv"
)

func two(lines []string) {
	_map := build(lines)
	heads := findPoints(lines, "0")
	//fmt.Println(heads)
	ratings := []int{}
	for i := 0; i < len(heads); i++ {
		ratings = append(ratings, findNext(&_map, [2]int{len(lines[0]), len(lines)}, heads[i], 0, check_unique_trails))
	}
	fmt.Println("RATINGS:", ratings)
	fmt.Println("SUM OF RATINGS:", addTrailHeads(ratings))
}

func check_unique_trails(_map *map[[2]int]string, bounds [2]int, point [2]int, height int) int {
	if (*_map)[point] == "." {
		return 0
	}

	if (*_map)[point] == strconv.Itoa(height) {
		if (*_map)[point] == "9" {
			return 1
		} else {
			return findNext(_map, bounds, point, height, check_unique_trails)
		}
	}
	return 0
}
