package main

import (
	"fmt"
	"strings"
)

func two(line string, blinks int) {
	// Compared to Puzzle-1, more blinks means exponential increase in number of stones. So we cannot maintain the full array of individual stones without overloading the memory.
	// We can still maintain the cache of successors to speed up each iteration. To replace the stones array, create another cache that maintains the number of stones with the same number. We only process that unique stone and multiply each successor of that stone with the number of occurrences of predecessor stone.
	successors_cache := make(map[string][]string)
	count_cache := make(map[string]int)
	line = strings.TrimSpace(line)
	stones := strings.Fields(line)
	for _, s := range stones {
		count_cache[s] = 1
	}
	for i := 0; i < blinks; i++ {
		temp_count_cache := make(map[string]int)
		for s, count := range count_cache {
			changed := []string{}
			if v, ok := successors_cache[s]; ok {
				//fmt.Println("CACHE HIT 1: ", s, "->", v)
				changed = append(changed, v...)
			} else {
				result := change(s, &successors_cache)
				successors_cache[s] = result
				changed = append(changed, result...)
			}
			for _, c := range changed {
				temp_count_cache[c] += count
			}
		}
		count_cache = temp_count_cache
		//fmt.Println(count_cache)
		fmt.Println("BLINK:", i+1, "COUNT:", count(count_cache))
	}
	fmt.Println(count(count_cache))
}

func count(cache map[string]int) (c int) {
	for _, v := range cache {
		c += v
	}
	return
}
