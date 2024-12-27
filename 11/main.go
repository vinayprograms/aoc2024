package main

import (
	"fmt"
	"global"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 10: Plutonian Pebbles")
	if len(os.Args) < 3 {
		fmt.Println("Must supply input file and count.")
		os.Exit(-1)
	}
	file := os.Args[1]
	lines := global.Load(file)
	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	fmt.Println("--- Puzzle 1 ---")
	one(lines[0], count)
	fmt.Println("--- Puzzle 2 ---")
	//one(file, 75)
}
