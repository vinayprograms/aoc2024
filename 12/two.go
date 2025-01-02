package main

import "fmt"

func two(lines []string) {
	_map := build(lines)
	cornerMap := make(map[[2]int]int)

	groups := getGroups(_map)
	for _, grp := range groups {
		for _, point := range grp {
			cornerMap[[2]int{point[0], point[1]}] = getSidesCount(_map, point[1], point[0], grp)
			//fmt.Printf("%s @ [%d,%d]: %d\n", _map[point[0]][point[1]], point[1], point[0], cornerMap[[2]int{point[0], point[1]}])
		}
	}

	price := 0
	for _, g := range groups {
		corners := 0
		for _, point := range g {
			corners += cornerMap[point]
		}
		price += corners * len(g)
		//fmt.Printf("%d * %d = %d; %d\n", len(g), corners, corners*len(g), price)
	}
	fmt.Println("PRICE:", price)
}

func getSidesCount(_map [][]string, x int, y int, grp [][2]int) int {
	count := 0
	if hasCorner(_map, x, y, [2]int{-1, -1}, grp) {
		count++
	}
	if hasCorner(_map, x, y, [2]int{-1, 1}, grp) {
		count++
	}
	if hasCorner(_map, x, y, [2]int{1, 1}, grp) {
		count++
	}
	if hasCorner(_map, x, y, [2]int{1, -1}, grp) {
		count++
	}
	return count
}

func hasCorner(_map [][]string, x int, y int, direction [2]int, grp [][2]int) bool {
	var val1, val2, valDiagonal string
	if y+direction[0] < 0 || y+direction[0] >= len(_map) {
		val1 = "-"
		if x+direction[1] < 0 || x+direction[1] >= len(_map[0]) {
			val2 = "-"
		} else {
			val2 = _map[y][x+direction[1]]
		}
		valDiagonal = "-"
	} else if x+direction[1] < 0 || x+direction[1] >= len(_map[0]) {
		val1 = _map[y+direction[0]][x]
		val2 = "-"
		valDiagonal = "-"
	} else {
		val1 = _map[y+direction[0]][x]
		val2 = _map[y][x+direction[1]]
		if contains(grp, [2]int{y + direction[0], x + direction[1]}) {
			valDiagonal = _map[y+direction[0]][x+direction[1]]
		} else {
			valDiagonal = "-"
		}
	}
	// Convex1: Adjacents and diagonal not same as current
	// Convex2: Adjacents not same, diagonal same as current
	// Concave: Adjacents are same as current, but diagonal is not
	if (val1 != _map[y][x] && val2 != _map[y][x] && valDiagonal != _map[y][x]) || // Convex1
		(val1 != _map[y][x] && val2 != _map[y][x] && valDiagonal == _map[y][x]) || // Convex2
		(val1 == _map[y][x] && val2 == _map[y][x] && _map[y][x] != valDiagonal) { // Concave
		return true
	}
	return false
}

func contains(group [][2]int, point [2]int) bool {
	for _, p := range group {
		if p[0] == point[0] && p[1] == point[1] {
			return true
		}
	}
	return false
}
