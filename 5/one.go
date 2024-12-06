package main

import "fmt"

func one(lines []string) {
	section1, section2 := getSections(lines)
	g := build(section1)
	//fmt.Println(g)
	result := 0
	for _, update := range section2 {
		if inOrder(g, update) {
			middle := int(len(update) / 2)
			//fmt.Println("IN ORDER:", update, "MIDDLE:", middle)
			result += update[middle]
		}
	}
	fmt.Println("TOTAL:", result)
}
