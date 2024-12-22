package main

import (
	"fmt"
)

func one(lines []string, lenX int, lenY int, count int) {
	obstacles := getObstacles(lines)
	_map := buildMap(obstacles, lenX, lenY, count)
	nodeCosts := make(map[[2]int]int)
	nodeCosts[[2]int{0, 0}] = 0
	potentials := make(map[[2]int]path)
	potentials[[2]int{0, 0}] = path{cost: 0, steps: [][2]int{{0, 0}}}
	result, err := next(_map, potentials, nodeCosts, 0, [2]int{0, 0}, [2]int{lenX - 1, lenY - 1})
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("COST:", result.cost, "ERROR:", err)
	}
}
