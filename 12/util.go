package main

func checkNorth(_map [][]string, x int, y int) bool {
	switch {
	case y == 0:
		return false
	case _map[y][x] == _map[y-1][x]:
		return true
	}
	return false
}

func checkSouth(_map [][]string, x int, y int) bool {
	switch {
	case y == len(_map)-1:
		return false
	case _map[y][x] == _map[y+1][x]:
		return true
	}
	return false
}

func checkEast(_map [][]string, x int, y int) bool {
	switch {
	case x == len(_map)-1:
		return false
	case _map[y][x] == _map[y][x+1]:
		return true
	}
	return false
}

func checkWest(_map [][]string, x int, y int) bool {
	switch {
	case x == 0:
		return false
	case _map[y][x] == _map[y][x-1]:
		return true
	}
	return false
}

func getGroups(_map [][]string) (groups [][][2]int) {
	checked := make(map[[2]int]bool)

	for y := 0; y < len(_map); y++ {
		for x := 0; x < len(_map[y]); x++ {
			point := [2]int{y, x}
			if checked[point] { // skip, if point was already processed
				continue
			}
			newGrp := search(_map, x, y, &checked)
			if len(newGrp) > 0 {
				groups = append(groups, newGrp)
			}
			checked[point] = true
		}
	}
	return
}

func search(_map [][]string, x int, y int, chkd *map[[2]int]bool) [][2]int {
	group := [][2]int{}
	point := [2]int{y, x}
	if (*chkd)[point] { // if this point was already visited
		return group // return empty group
	}
	(*chkd)[point] = true
	group = append(group, point)
	if checkNorth(_map, x, y) {
		group = append(group, search(_map, x, y-1, chkd)...)
	}
	if checkSouth(_map, x, y) {
		group = append(group, search(_map, x, y+1, chkd)...)
	}
	if checkEast(_map, x, y) {
		group = append(group, search(_map, x+1, y, chkd)...)
	}
	if checkWest(_map, x, y) {
		group = append(group, search(_map, x-1, y, chkd)...)
	}
	(*chkd)[point] = true
	return group
}
