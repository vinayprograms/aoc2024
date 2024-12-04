package main

import (
	"global"
	"strconv"
	"strings"
)

// Read a report line and extract levels
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

// Check if report is continuously increasing or continuously decreasing.
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

// Check if levels are safe (as per logic defined in the puzzles)
func isSafe(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		diff := global.Abs(numbers[i] - numbers[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
