package main

import (
	"fmt"
	"global"
	"os"
)

func main() {
	fmt.Println("Day 10: Plutonian Pebbles")
	if len(os.Args) != 2 {
		fmt.Println("Must supply input file.")
		os.Exit(-1)
	}
	file := os.Args[1]
	lines := global.Load(file)
	fmt.Println("--- Puzzle 1 ---")
	one(lines[0], 25)
	fmt.Println("--- Puzzle 2 ---")
	two(lines[0], 75)
}
