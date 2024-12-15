package main

import (
	"fmt"
	"global"
	"os"
)

func main() {
	fmt.Println("Day 8: Resonant Collinearity")
	if len(os.Args) != 2 {
		fmt.Println("Must supply input file.")
		os.Exit(-1)
	}
	file := os.Args[1]
	lines := global.Load(file)
	fmt.Println("--- Puzzle 1 ---")
	one(lines)
	fmt.Println("--- Puzzle 2 ---")
	two(lines)
}
