package main

import (
	"fmt"
	"global"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 12: RAM Run")
	if len(os.Args) < 2 {
		fmt.Println("Must supply input file.")
		os.Exit(-1)
	}
	file := os.Args[1]
	lines := global.Load(file)
	var lenX, lenY, count int
	if val, err := strconv.Atoi(os.Args[2]); err == nil {
		lenX = val
	} else {
		panic(err)
	}
	if val, err := strconv.Atoi(os.Args[3]); err == nil {
		lenY = val
	} else {
		panic(err)
	}
	if val, err := strconv.Atoi(os.Args[4]); err == nil {
		count = val
	} else {
		panic(err)
	}

	fmt.Println("--- Puzzle 1 ---")
	one(lines, lenX, lenY, count)
	fmt.Println("--- Puzzle 2 ---")
	two(lines, lenX, lenY)
}
