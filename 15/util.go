package main

import (
	"regexp"
	"strings"
)

type game struct {
	warehouse [][]string
	moves     string
	robot     [2]int
}

func input(lines []string) game {
	g := game{}
	for i := 0; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" {
			// consolidate the moves
			g.moves = strings.Join(lines[i+1:], "")
			break
		} else {
			cells := strings.Split(lines[i], "")
			for j := 0; j < len(cells); j++ {
				if cells[j] == "@" {
					g.robot = [2]int{i, j}
				}
			}
			g.warehouse = append(g.warehouse, strings.Split(lines[i], ""))
		}
	}
	return g
}

func left(state [][]string, position [2]int) ([][]string, [2]int) {
	object := state[position[0]][position[1]-1]
	switch object {
	case "O":
		state, position = pushLeft(state, position)
	case ".":
		state[position[0]][position[1]-1] = "@"
		state[position[0]][position[1]] = "."
		position[1] = position[1] - 1
	}
	return state, position
}

func pushLeft(state [][]string, position [2]int) ([][]string, [2]int) {
	line := strings.Join(state[position[0]], "")
	re := regexp.MustCompile("\\.(O)+@")
	loc := re.FindStringIndex(line)
	if len(loc) > 0 && line[loc[0]:loc[1]] != "" {
		for i := loc[0]; i < loc[1]-1; i++ {
			state[position[0]][i] = state[position[0]][i+1]
		}
		state[position[0]][loc[1]-1] = "."
		position[1] = position[1] - 1
		return state, position
	}
	return state, position
}

func right(state [][]string, position [2]int) ([][]string, [2]int) {
	object := state[position[0]][position[1]+1]
	switch object {
	case "O":
		state, position = pushRight(state, position)
	case ".":
		state[position[0]][position[1]+1] = "@"
		state[position[0]][position[1]] = "."
		position[1] = position[1] + 1
	}
	return state, position
}

// Check right. Should be one or more "O" followed by "."
func pushRight(state [][]string, position [2]int) ([][]string, [2]int) {
	line := strings.Join(state[position[0]], "")
	re := regexp.MustCompile("@(O)+\\.")
	loc := re.FindStringIndex(line)
	if len(loc) > 0 && line[loc[0]:loc[1]] != "" {
		for i := loc[1] - 1; i > loc[0]; i-- {
			state[position[0]][i] = state[position[0]][i-1]
		}
		state[position[0]][loc[0]] = "."
		position[1] = position[1] + 1
		return state, position
	}
	return state, position
}

func up(state [][]string, position [2]int) ([][]string, [2]int) {
	object := state[position[0]-1][position[1]]
	switch object {
	case "O":
		state, position = pushUp(state, position)
	case ".":
		state[position[0]-1][position[1]] = "@"
		state[position[0]][position[1]] = "."
		position[0] = position[0] - 1
	}
	return state, position
}

func pushUp(state [][]string, position [2]int) ([][]string, [2]int) {
	line := ""
	for y := 0; y < len(state); y++ {
		line += state[y][position[1]]
	}
	re := regexp.MustCompile("\\.(O)+@")
	loc := re.FindStringIndex(line)
	if len(loc) > 0 && line[loc[0]:loc[1]] != "" {
		for i := loc[0]; i < loc[1]-1; i++ {
			state[i][position[1]] = state[i+1][position[1]]
		}
		state[loc[1]-1][position[1]] = "."
		position[0] = position[0] - 1
		return state, position
	}
	return state, position
}

func down(state [][]string, position [2]int) ([][]string, [2]int) {
	object := state[position[0]+1][position[1]]
	switch object {
	case "O":
		state, position = pushDown(state, position)
	case ".":
		state[position[0]+1][position[1]] = "@"
		state[position[0]][position[1]] = "."
		position[0] = position[0] + 1
	}
	return state, position
}

func pushDown(state [][]string, position [2]int) ([][]string, [2]int) {
	line := ""
	for y := 0; y < len(state); y++ {
		line += state[y][position[1]]
	}
	re := regexp.MustCompile("@(O)+\\.")
	loc := re.FindStringIndex(line)
	if len(loc) > 0 && line[loc[0]:loc[1]] != "" {
		for i := loc[1] - 1; i > loc[0]; i-- {
			state[i][position[1]] = state[i-1][position[1]]
		}
		state[loc[0]][position[1]] = "."
		position[0] = position[0] + 1
		return state, position
	}
	return state, position
}
