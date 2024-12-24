package main

import (
	"fmt"
	"global"
	"os"
)

func main() {
	fmt.Println("Day 19: Linen Layout")
	if len(os.Args) < 2 {
		fmt.Println("Must supply input file.")
		os.Exit(-1)
	}
	file := os.Args[1]
	lines := global.Load(file)
	towels, designs := prepInput(lines)
	fmt.Println("--- Puzzle 1 ---")
	one(towels, designs)
	fmt.Println("--- Puzzle 2 ---")
	two(towels, designs)
}
