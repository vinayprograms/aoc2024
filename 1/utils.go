package main

import (
	"strconv"
	"strings"
)

// Generate 'left' and 'right' lists from input lines.
// Each line contains a pair of numbers separated by one or more spaces.
// The first number is the left element and the second one the right.
func getlists(lines []string) ([]int, []int) {
	var left, right []int
	for _, line := range lines {
		elements := strings.Fields(line)
		if len(elements) != 2 {
			continue
		}
		l, err := strconv.Atoi(elements[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(elements[1])
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}
