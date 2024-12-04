package main

import (
	"fmt"
	"global"
	"sort"
)

func one(lines []string) {
	left, right := getlists(lines)
	left, right = pair(left, right)
	diffs := diff(left, right)
	total := global.Sum(diffs)
	fmt.Println(total)
}

// Pair numbers from smallest to largest on both lists
func pair(left []int, right []int) ([]int, []int) {
	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

// Calculate the difference between adjacent elements of 'left' and 'right' lists.
func diff(left []int, right []int) []int {
	diffs := []int{}

	for i := range len(left) {
		diffs = append(diffs, global.Abs(left[i]-right[i]))
	}

	return diffs
}
