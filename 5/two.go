package main

import "fmt"

func two(lines []string) {
	section1, section2 := getSections(lines)
	g := build(section1)
	//fmt.Println(g)
	total := 0
	for _, update := range section2 {
		if inOrder(g, update) == false {
			sorted := g.Sort(update, isPrevious)
			//fmt.Println(sorted)
			middle := int(len(sorted) / 2)
			//fmt.Printf("MIDDLE[%d]: %d\n", middle, sorted[middle])
			total += sorted[middle]
		}
	}
	fmt.Println("TOTAL:", total)
}
