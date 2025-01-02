package main

import "fmt"

func one(lines []string) {
	_map := build(lines)
	perimeter := make(map[[2]int]int)
	for y := 0; y < len(_map); y++ {
		for x := 0; x < len(_map[y]); x++ {
			perimeter[[2]int{y, x}] = getFenceCount(_map, x, y)
			//fmt.Printf("%s @ [%d,%d]: %d\n", _map[y][x], x, y, perimeter[[2]int{y, x}])
		}
	}

	groups := getGroups(_map)
	price := 0
	for _, g := range groups {
		boundary := 0
		for _, point := range g {
			boundary += perimeter[point]
		}
		//fmt.Printf("%d * %d = %d\n", len(g), boundary, boundary*len(g))
		price += boundary * len(g)
	}
	fmt.Println("PRICE:", price)
}

func build(lines []string) [][]string {
	_map := [][]string{}
	for y := 0; y < len(lines); y++ {
		line := []string{}
		for x := 0; x < len(lines[y]); x++ {
			line = append(line, string(lines[y][x]))
		}
		_map = append(_map, line)
	}
	return _map
}

func getFenceCount(_map [][]string, x int, y int) int {
	count := 4
	if checkNorth(_map, x, y) {
		count--
	}
	if checkSouth(_map, x, y) {
		count--
	}
	if checkEast(_map, x, y) {
		count--
	}
	if checkWest(_map, x, y) {
		count--
	}
	return count
}
