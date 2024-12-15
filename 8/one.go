package main

import (
	"fmt"
	g "global"
)

func one(lines []string) {
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
	grps := buildAntennaGroups(m)
	antinodes := findAntiNodes(m, grps, true)
	uniqueNodes := []g.Point{}
	for _, v := range antinodes {
		for _, n := range v {
			if !contains(n, uniqueNodes) {
				uniqueNodes = append(uniqueNodes, n)
			}
		}
	}
	fmt.Println("NO. OF ANTINODES:", len(uniqueNodes))
}
