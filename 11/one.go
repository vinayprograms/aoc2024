package main

import (
	"fmt"
	"strconv"
	"strings"
)

func one(line string, blinks int) {
	cache := make(map[string][]string)
	line = strings.TrimSpace(line)
	stones := strings.Split(line, " ")
	for i := 0; i < blinks; i++ {
		changed := []string{}
		for _, s := range stones {
			s = strings.TrimSpace(s)
			if v, ok := cache[s]; ok {
				//fmt.Println("CACHE HIT 1: ", s, "->", v)
				changed = append(changed, v...)
			} else {
				result := change(s, &cache)
				cache[s] = result
				changed = append(changed, result...)
			}
		}
		fmt.Println("BLINK:", i, "COUNT:", len(stones))
		stones = changed
	}
	fmt.Println(len(stones))
	//fmt.Println(strings.TrimSpace(line))
}

func change(s string, cache *map[string][]string) []string {
	if v, present := (*cache)[s]; present {
		//fmt.Println("CACHE HIT: ", s, "->", v)
		return v
	}
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	switch {
	case s == "0":
		return []string{"1"}
	case len(s)%2 == 0:
		s1 := s[:len(s)/2]
		intval, _ := strconv.Atoi(s1)
		s1 = fmt.Sprintf("%d", intval)
		s2 := s[len(s)/2:]
		intval, _ = strconv.Atoi(s2)
		s2 = fmt.Sprintf("%d", intval)
		return []string{s1, s2}
	default:
		intval, _ := strconv.Atoi(s)
		intval = intval * 2024
		return []string{fmt.Sprintf("%d", intval)}
	}
}
