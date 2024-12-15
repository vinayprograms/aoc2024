package main

import (
	"fmt"
	g "global"
)

func find(m g.Map[string], sprite string) g.Point {
	start := g.Point{Coordinates: []int{-1, -1}} // Start with invalid position
	for p, v := range m.Positions {
		if v == sprite {
			// convert string "[x, y]" to Point{coordinates: [x, y]}
			var x, y int
			fmt.Sscanf(p, "[%d %d]", &x, &y)
			start = g.Point{Coordinates: []int{x, y}}
			break
		}
	}

	return start
}

func buildAntennaGroups(m g.Map[string]) map[string][]g.Point {
	groups := make(map[string][]g.Point)
	for k, v := range m.Positions {
		if v == "." {
			continue
		}
		if _, present := groups[v]; !present {
			groups[v] = []g.Point{}
		}
		groups[v] = append(groups[v], g.PointFromString(k))
	}
	return groups
}

func findAntiNodes(m g.Map[string], groups map[string][]g.Point, adjacentOnly bool) map[string][]g.Point {
	nodes := make(map[string][]g.Point)
	for c, grp := range groups {
		//fmt.Println("-----", c, "pair -----")
		for i := 0; i < len(grp); i++ {
			for j := i + 1; j < len(grp); j++ {
				//fmt.Println(i, j)
				for _, n := range getAntinodesForPair(m, grp[i], grp[j], adjacentOnly) {
					if !contains(n, nodes[c]) {
						nodes[c] = append(nodes[c], n)
					}
				}
			}
		}
	}
	return nodes
}

func getAntinodesForPair(m g.Map[string], antenna1 g.Point, antenna2 g.Point, adjacentOnly bool) []g.Point {
	antinodes := []g.Point{}
	//fmt.Println("PAIR:", antenna1.String(), antenna2.String())
	// In direction
	dist := g.Distance(antenna1, antenna2)
	potential := antenna2
	anCount := 0
	for {
		potential, _ = g.Move(potential, dist)
		//fmt.Println("POTENTIAL:", potential.String())
		if !m.IsInsideMap(potential) {
			break
		}
		if contains(potential, antinodes) {
			if !adjacentOnly {
				continue
			}
		} else {
			anCount++
			//fmt.Printf("P1(%s) -> P2(%s) => AN%da(%s)\n", antenna1.String(), antenna2.String(), anCount, potential.String())
			antinodes = append(antinodes, potential)
			if adjacentOnly {
				break
			}
		}
	}

	// Opposite direction
	dist = g.Distance(antenna2, antenna1)
	potential = antenna1
	anCount = 0
	for {
		potential, _ = g.Move(potential, dist)
		//fmt.Println("POTENTIAL:", potential.String(), "DIST:", dist.String())
		if !m.IsInsideMap(potential) {
			break
		}
		if contains(potential, antinodes) {
			if !adjacentOnly {
				continue
			}
		} else {
			anCount++
			//fmt.Printf("P2(%s) -> P1(%s) => AN%db(%s)\n", antenna2.String(), antenna1.String(), anCount, potential.String())
			antinodes = append(antinodes, potential)
			if adjacentOnly {
				break
			}
		}
	}

	return antinodes
}

func contains(p g.Point, points []g.Point) bool {
	for _, x := range points {
		if g.IsSamePoint(p, x) {
			return true
		}
	}
	return false
}
