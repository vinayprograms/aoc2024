package main

import (
	"fmt"
	"slices"
)

func two(reports []string) {
	safeCount := 0
	for _, r := range reports {
		numbers := levels(r)
		if len(numbers) == 0 {
			continue
		}
		if test(numbers) == false {
			for _, d := range dampen(numbers) {
				if test(d) == true {
					safeCount++
					break
				}
			}
		} else {
			safeCount++
		}
	}
	fmt.Println("SAFE:", safeCount)
}

func dampen(arr []int) [][]int {
	dampened := [][]int{}
	for i := 0; i < len(arr); i++ {
		dArr := slices.Clone(arr)
		dArr = slices.Delete(dArr, i, i+1)
		dampened = append(dampened, dArr)
	}
	return dampened
}
