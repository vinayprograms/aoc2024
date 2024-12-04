package main

import (
	"fmt"
)

func two(lines []string) {
	left, right := getlists(lines)
	rightMap := buildMap(right)
	similarity_score := calculate(left, rightMap)
	fmt.Println(similarity_score)
}

func buildMap(list []int) map[int]int {
	numMap := make(map[int]int)
	for _, element := range list {
		if _, ok := numMap[element]; ok {
			numMap[element]++
		} else {
			numMap[element] = 1
		}
	}
	return numMap
}

func calculate(left []int, rightMap map[int]int) int {
	var score = 0
	for _, num := range left {
		if val, ok := rightMap[num]; ok {
			score += num * val
		}
	}
	return score
}
