package main

import (
	"fmt"
)

func two(lines []string, lenX int, lenY int) {
	obstacles := getObstacles(lines)
	for i := 0; i < len(obstacles); i++ {
		_map := buildMap(obstacles, lenX, lenY, i)
		nodeCosts := make(map[[2]int]int)
		nodeCosts[[2]int{0, 0}] = 0
		potentials := make(map[[2]int]path)
		potentials[[2]int{0, 0}] = path{cost: 0, steps: [][2]int{{0, 0}}}
		_, err := next(_map, potentials, nodeCosts, 0, [2]int{0, 0}, [2]int{lenX - 1, lenY - 1})
		if err != nil {
			fmt.Printf("❌ %d\n", i)
			fmt.Printf("ERROR: %s. Failed when obstacle number %d dropped @ location [%v]\n", err, i, lines[i-1])
			return
		} else {
			fmt.Printf("✅ %d,", i)
		}
	}
}
