package main

import (
	"fmt"
	"strings"
)

func two(lines []string, maxX int, maxY int, iterations int) {
	robots := getRobots(lines)
	for i := 0; i < iterations; i++ {
		final := move(robots, maxX, maxY, i)
		m := render(final, maxX, maxY)
		if look(m) == false {
			continue
		}
		for x := 0; x < maxX; x++ {
			for y := 0; y < maxY; y++ {
				fmt.Print(m[x][y])
			}
			fmt.Print("\n")
		}
		fmt.Println("++++++++++++", i, "++++++++++++")
		return
	}
}

func look(grid [][]string) bool {
	for y := 0; y < len(grid); y++ {
		line := strings.Join(grid[y], "")
		if strings.Contains(line, "##########") {
			return true
		}
	}
	return false
}

func render(robots []robot, maxX int, maxY int) [][]string {
	_map := [][]string{}
	for x := 0; x < maxX; x++ {
		line := []string{}
		for y := 0; y < maxY; y++ {
			line = append(line, " ")
		}
		_map = append(_map, line)
	}
	for _, robot := range robots {
		_map[robot.position[0]][robot.position[1]] = "#"
	}
	return _map
}
