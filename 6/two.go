package main

import (
	"fmt"
	g "global"
)

type result string

// The result of a walk.
// "outside" - Guard has moved out of the map
// "loop" - Guard has entered a loop (This is what we want to achieve)
const (
	outside result = "outside"
	loop    result = "loop"
)

// A guard's waypoint is a combination of the location on the map and the direction they are facing.
type waypoint struct {
	location  string
	direction string
} // NOTE: Go doesn't allow generics as hashmap keys. So, both members are "string".

func two(lines []string) {
	// Extract each position into a single collection
	values := []string{}
	for _, l := range lines {
		for _, c := range l {
			values = append(values, string(c))
		}
	}

	m := g.Map[string]{}
	dimensions := []int{len(lines[0]), len(lines)}
	m.Build(dimensions, values)

	start := find(m, "^")
	if !m.IsInsideMap(start) {
		fmt.Println("Guard not found! You are free to explore :)")
	}

	loopCount := 0
	direction := g.Vector{Deltas: []int{0, -1}}
	// First, get all walking points.
	waypoints, _ := walk(m, start, direction)
	fmt.Println("WAYPOINTS:", len(waypoints))
	// Brute Force: Iterate through each waypoint, place an obstacle and check for loop
	uniqueLoops := []string{}
	for k, _ := range waypoints {
		if k.location == start.String() {
			continue
		}
		oldSprite := m.Positions[k.location]
		m.Positions[k.location] = "#"
		_, result := walk(m, start, direction)
		if result == loop {
			if contains(uniqueLoops, k.location) == false {
				loopCount++ // Only count unique instances of looping
				if loopCount%100 == 0 {
					fmt.Print(loopCount)
				} else {
					fmt.Print("*")
				}
				uniqueLoops = append(uniqueLoops, k.location)
			}
		}
		m.Positions[k.location] = oldSprite
	}
	fmt.Println()
	fmt.Println("TOTAL LOOPS:", loopCount)
}

func walk(m g.Map[string], start g.Point, direction g.Vector) (map[waypoint]bool, result) {

	waypoints := make(map[waypoint]bool)
	key := waypoint{location: start.String(), direction: direction.String()}
	waypoints[key] = true
	current := start
	for {
		newPos, err := moveOne(current, direction)
		if err != nil || m.IsInsideMap(newPos) == false {
			return waypoints, outside
		}
		if hasObstruction(m, newPos) {
			direction = turn(direction)
			// If this turning was visited earlier, we have a loop
			key = waypoint{location: current.String(), direction: direction.String()}
			if yes, ok := waypoints[key]; ok && yes {
				return waypoints, loop
			} else {
				waypoints[key] = true
			}
		} else {
			current = newPos
			key := waypoint{location: current.String(), direction: direction.String()}
			if yes, ok := waypoints[key]; ok && yes {
				return waypoints, loop
			} else {
				waypoints[key] = true
			}
		}
	}
}

func contains(arr []string, s string) bool {
	for _, x := range arr {
		if x == s {
			return true
		}
	}
	return false
}
