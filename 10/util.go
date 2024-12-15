package main

func build(lines []string) (_map map[[2]int]string) {
	_map = make(map[[2]int]string)
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			point := [2]int{x, y}
			_map[point] = string(lines[y][x])
		}
	}
	return
}

func findPoints(lines []string, sprite string) [][2]int {
	points := [][2]int{}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if string(lines[y][x]) == sprite {
				points = append(points, [2]int{x, y})
			}
		}
	}
	return points
}

func addTrailHeads(heads []int) int {
	sum := 0
	for _, x := range heads {
		sum += x
	}
	return sum
}

type checker func(*map[[2]int]string, [2]int, [2]int, int) int

func findNext(_map *map[[2]int]string, bounds [2]int, point [2]int, height int, chk checker) (pathCount int) {
	north := [2]int{point[0], point[1] - 1}
	if north[1] >= 0 {
		pathCount += chk(_map, bounds, north, height+1)
		//fmt.Println("C:", pathCount, "P:", point, "H:", height)
	}
	south := [2]int{point[0], point[1] + 1}
	if south[1] < bounds[1] {
		pathCount += chk(_map, bounds, south, height+1)
		//fmt.Println("C:", pathCount, "P:", point, "H:", height)
	}
	east := [2]int{point[0] + 1, point[1]}
	if east[0] < bounds[0] {
		pathCount += chk(_map, bounds, east, height+1)
		//fmt.Println("C:", pathCount, "P:", point, "H:", height)
	}
	west := [2]int{point[0] - 1, point[1]}
	if west[0] >= 0 {
		pathCount += chk(_map, bounds, west, height+1)
		//fmt.Println("C:", pathCount, "P:", point, "H:", height)
	}
	return
}
