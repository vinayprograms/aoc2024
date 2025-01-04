package main

import (
	"fmt"
	"global"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 14: Restroom Redoubt")
	if len(os.Args) < 5 {
		fmt.Println("Must supply input file.")
		os.Exit(-1)
	}
	file := os.Args[1]
	lines := global.Load(file)
	var maxX, maxY, iterations int
	if val, err := strconv.Atoi(os.Args[2]); err == nil {
		maxX = val
	} else {
		panic(err)
	}
	if val, err := strconv.Atoi(os.Args[3]); err == nil {
		maxY = val
	} else {
		panic(err)
	}
	if val, err := strconv.Atoi(os.Args[4]); err == nil {
		iterations = val
	} else {
		panic(err)
	}
	fmt.Println("--- Puzzle 1 ---")
	one(lines, maxX, maxY)
	fmt.Println("--- Puzzle 2 ---")
	two(lines, maxX, maxY, iterations)
}
