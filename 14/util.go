package main

import (
	"strconv"
	"strings"
)

type robot struct {
	position [2]int
	velocity [2]int
}

func getRobots(lines []string) []robot {
	robots := []robot{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		robots = append(robots, getRobot(line))
	}
	return robots
}

func getRobot(line string) robot {
	r := robot{}
	parts := strings.Split(line, " ")
	r.position = getPoint(strings.Split(parts[0], "=")[1])
	r.velocity = getPoint(strings.Split(parts[1], "=")[1])
	return r
}

func getPoint(p string) (point [2]int) {
	values := strings.Split(p, ",")
	if v, err := strconv.Atoi(values[0]); err == nil {
		point[0] = v
	} else {
		panic(err)
	}
	if v, err := strconv.Atoi(values[1]); err == nil {
		point[1] = v
	} else {
		panic(err)
	}
	return
}

func move(robots []robot, maxX int, maxY int, seconds int) []robot {
	final := []robot{}
	for _, r := range robots {
		posX := (r.position[0] + (r.velocity[0] * seconds)) % maxX
		if posX < 0 {
			posX = maxX + posX
		}
		posY := (r.position[1] + (r.velocity[1] * seconds)) % maxY
		if posY < 0 {
			posY = maxY + posY
		}
		f := robot{
			position: [2]int{posX, posY},
			velocity: r.velocity,
		}
		final = append(final, f)
	}
	return final
}
