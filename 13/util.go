package main

import (
	"strconv"
	"strings"
)

type game struct {
	a     [2]int
	b     [2]int
	prize [2]int
}

func getGames(lines []string) (games []game) {
	for i := 0; i < len(lines); i += 4 {
		g := game{
			a:     getValues(lines[i]),
			b:     getValues(lines[i+1]),
			prize: getValues(lines[i+2]),
		}
		games = append(games, g)
	}
	return
}

func getValues(line string) [2]int {
	parts := strings.Split(line, ":")
	parts[1] = strings.TrimSpace(parts[1])
	coords := strings.Split(parts[1], ",")
	x, err := strconv.Atoi(strings.TrimSpace(coords[0])[2:])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(strings.TrimSpace(coords[1])[2:])
	if err != nil {
		panic(err)
	}
	return [2]int{x, y}
}

func fixGames(games []game) (fixed []game) {
	for _, g := range games {
		g.prize[0] = 10000000000000 + g.prize[0]
		g.prize[1] = 10000000000000 + g.prize[1]
		fixed = append(fixed, g)
	}
	return fixed
}
