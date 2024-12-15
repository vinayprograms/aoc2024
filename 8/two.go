package main

import (
	"fmt"
	g "global"
)

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
	grps := buildAntennaGroups(m)
	antinodes := findAntiNodes(m, grps, false)
	uniqueNodes := []g.Point{}
	for _, v := range antinodes {
		for _, n := range v {
			if !contains(n, uniqueNodes) {
				uniqueNodes = append(uniqueNodes, n)
			}
		}
	}

	freeAntennasCount := 0
	// Filter out overlaps with antenna position
	for _, v := range grps {
		for _, a := range v {
			if !contains(a, uniqueNodes) {
				freeAntennasCount++
			}
		}
	}
	fmt.Println("NO. OF RESONANT NODES:", len(uniqueNodes)+freeAntennasCount)
}
