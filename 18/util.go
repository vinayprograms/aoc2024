package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getObstacles(lines []string) [][2]int {
	obstacles := [][2]int{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ",")
		var x, y int
		if val, err := strconv.Atoi(parts[0]); err != nil {
			panic(err)
		} else {
			x = val
		}
		if val, err := strconv.Atoi(parts[1]); err != nil {
			panic(err)
		} else {
			y = val
		}
		obstacles = append(obstacles, [2]int{y, x})
	}
	return obstacles
}

func buildMap(obstacles [][2]int, lenX int, lenY int, count int) [][]string {
	_map := [][]string{}
	for y := 0; y < lenY; y++ {
		mapLine := []string{}
		for x := 0; x < lenX; x++ {
			mapLine = append(mapLine, ".")
		}
		_map = append(_map, mapLine)
	}

	for o := 0; o < count; o++ {
		_map[obstacles[o][0]][obstacles[o][1]] = "#"
	}
	return _map
}

func printMap(_map [][]string) {
	for _, line := range _map {
		fmt.Printf("%v\n", line)
	}
}

type path struct {
	cost  int
	steps [][2]int
}

func next(_map [][]string, potentials map[[2]int]path, nodeCosts map[[2]int]int, currCost int, currPoint [2]int, dest [2]int) (path, error) {
	// Add all possible successor nodes to the list of potentials
	for _, n := range getNeighbors(_map, currPoint) {
		var updated [][2]int
		if _, ok := potentials[n]; ok {
			updated = append(potentials[n].steps, n)
		} else {
			updated = append(updated, potentials[currPoint].steps...)
			updated = append(updated, n)
		}
		if v, ok := nodeCosts[n]; ok {
			if v > currCost+1 {
				nodeCosts[n] = currCost + 1
				potentials[n] = path{cost: currCost + 1, steps: updated} // update cost to reach a potential node.
			}
		} else {
			nodeCosts[n] = currCost + 1
			potentials[n] = path{cost: currCost + 1, steps: updated}
		}
	}
	delete(potentials, currPoint)
	if len(potentials) == 0 {
		return path{cost: currCost}, fmt.Errorf("Cannot reach destination")
	}
	_point, _path := lowestCost(potentials, dest)
	if _point[0] == dest[0] && _point[1] == dest[1] {
		return _path, nil
	}
	//fmt.Println("POINT:", _point, "PATH:", _path)
	return next(_map, potentials, nodeCosts, _path.cost, _point, dest)
}

func getNeighbors(_map [][]string, point [2]int) (neighbors [][2]int) {
	if point[1]-1 >= 0 && _map[point[0]][point[1]-1] != "#" { // left
		neighbors = append(neighbors, [2]int{point[0], point[1] - 1})
	}
	if point[1]+1 < len(_map) && _map[point[0]][point[1]+1] != "#" { // right
		neighbors = append(neighbors, [2]int{point[0], point[1] + 1})
	}
	if point[0]-1 >= 0 && _map[point[0]-1][point[1]] != "#" { // above
		neighbors = append(neighbors, [2]int{point[0] - 1, point[1]})
	}
	if point[0]+1 < len(_map[0]) && _map[point[0]+1][point[1]] != "#" { // below
		neighbors = append(neighbors, [2]int{point[0] + 1, point[1]})
	}
	return
}

func lowestCost(potentials map[[2]int]path, destination [2]int) (point [2]int, p path) {
	cost := math.MaxInt
	for k, v := range potentials {
		h := heuristic(k, destination)
		if v.cost+h < cost {
			cost = v.cost + h
			p = v
			point = k
		}
	}
	return
}

func heuristic(point [2]int, destination [2]int) int {
	return int(math.Abs(float64(point[0])-float64(destination[0]))) + int(math.Abs(float64(point[1])-float64(destination[1])))
}
