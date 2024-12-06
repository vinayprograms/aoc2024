package main

import (
	"strconv"
	"strings"
)

func getSections(lines []string) ([][2]int, [][]int) {
	section1 := [][2]int{}
	section2 := [][]int{}

	startSection2 := false
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			startSection2 = true
			continue
		}
		if startSection2 {
			sequence := []int{}
			parts := strings.Split(l, ",")

			for _, p := range parts {
				if value, err := strconv.Atoi(p); err != nil {
					panic(err)
				} else {
					sequence = append(sequence, value)
				}
			}
			section2 = append(section2, sequence)
		} else {
			parts := strings.Split(l, "|")
			var a, b int
			if value, err := strconv.Atoi(parts[0]); err != nil {
				panic(err)
			} else {
				a = value
			}
			if value, err := strconv.Atoi(parts[1]); err != nil {
				panic(err)
			} else {
				b = value
			}
			section1 = append(section1, [2]int{a, b})
		}
	}
	return section1, section2
}

func build(section1 [][2]int) Graph[int] {
	g := Graph[int]{} // This is a graph of pages with page-number only.
	for _, entry := range section1 {
		g.Insert(entry[0], entry[1])
	}
	g.Consolidate()
	return g
}

func inOrder(g Graph[int], sequence []int) bool {
	for i := 0; i < len(sequence)-1; i++ {
		if n, ok := g.nodeIndex[sequence[i]]; ok {
			if contains(sequence[i+1], n.Successors()) {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
