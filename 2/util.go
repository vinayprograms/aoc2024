package main

import (
	"os"
	"strconv"
	"strings"
)

// Read contents from input file and return as a set of lines
func load(file string) []string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	return lines[:len(lines)-1]
}

// Read contents from input file and return as a set of lines
func levels(report string) []int {
	levelValues := strings.Fields(report)
	var levels []int
	for _, l := range levelValues {
		level, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		levels = append(levels, level)
	}
	return levels
}

func test(arr []int) bool {
	return isMonotonic(arr) && isSafe(arr)
}

func isMonotonic(arr []int) bool {
	isMonotonic := false
	for i := 0; i < len(arr)-2; i++ {
		if ((arr[i+1]-arr[i] > 0) && (arr[i+2]-arr[i+1] > 0)) || // monotonically increasing
			((arr[i+1]-arr[i] < 0) && (arr[i+2]-arr[i+1] < 0)) { // monotonically decreasing
			isMonotonic = true
		} else {
			return false
		}
	}
	return isMonotonic
}

func isSafe(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		diff := abs(numbers[i] - numbers[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
