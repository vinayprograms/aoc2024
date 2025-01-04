package main

import "fmt"

func one(lines []string, maxX int, maxY int) {
	robots := getRobots(lines)
	final := move(robots, maxX, maxY, 100)
	sf := safetyFactor(final, maxX, maxY)
	fmt.Println(sf)
}

func safetyFactor(robots []robot, maxX int, maxY int) int {
	sf := 1
	sf *= count(robots, 0, 0, (maxX/2)-1, (maxY/2)-1)
	sf *= count(robots, (maxX/2)+1, 0, maxX-1, (maxY/2)-1)
	sf *= count(robots, 0, (maxY/2)+1, (maxX/2)-1, maxY-1)
	sf *= count(robots, (maxX/2)+1, (maxY/2)+1, maxX-1, maxY-1)
	return sf
}

func count(robots []robot, xMin int, yMin int, xMax int, yMax int) int {
	c := 0
	for _, r := range robots {
		if r.position[0] >= xMin && r.position[0] <= xMax && r.position[1] >= yMin && r.position[1] <= yMax {
			c++
		}
	}
	return c
}
