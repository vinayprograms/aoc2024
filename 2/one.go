package main

import (
	"fmt"
)

func one(reports []string) {
	safeCount := 0
	for _, r := range reports {
		numbers := levels(r)
		if len(numbers) == 0 {
			continue
		}
		if !isMonotonic(numbers) {
			continue
		}
		if isSafe(numbers) {
			safeCount++
		}
	}
	fmt.Println("SAFE:", safeCount)
}
