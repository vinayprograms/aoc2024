package main

import (
	"fmt"
	"math"
)

func one(lines []string) {
	total := 0

	// For puzzle one, we are going with bruteforce search
	for _, g := range getGames(lines) {
		options := [][2]int{}
		//fmt.Println("....", g, "....")
		n1 := int(g.prize[0] / g.a[0])
		//fmt.Println("N1:", n1)
		for n1 > 0 {
			remaining := g.prize[0] - (n1 * g.a[0])
			n2 := remaining / g.b[0]
			//fmt.Println("N1", n1, "N2:", n2)
			if n2 == 0 {
				n1--
				//fmt.Println("NOT FOUND:", "N1=", n1, "N2=", n2)
				continue
			}
			remainder := remaining % n2
			if remainder == 0 &&
				((n1*g.a[0])+(n2*g.b[0]) == g.prize[0] &&
					(n1*g.a[1])+(n2*g.b[1]) == g.prize[1]) {
				//fmt.Println("FOUND:", n1, ",", n2)
				options = append(options, [2]int{n1, n2})
				break
			} else {
				//fmt.Println("FAILED:", "N1=", n1, "N2=", n2)
				n1--
			}
		}
		//fmt.Println("OPTIONS:", options)
		if len(options) > 0 {
			min := math.MaxInt
			//answer := options[0]
			for _, o := range options {
				cost := o[0]*3 + o[1]
				if cost < min {
					min = cost
					//answer = o
				}
			}
			//fmt.Println(min, answer)
			total += min
		}
	}
	fmt.Println("TOTAL COST:", total)
}
