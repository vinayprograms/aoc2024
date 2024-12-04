package main

import (
	"fmt"
	"global"
	"os"
)

func main() {
	fmt.Println("DAY 1: Historian Hysteria")
	if len(os.Args) != 2 {
		fmt.Println(`Must supply input file.
		File must contain N lines with each line containing a space-separated list of numbers`)
		os.Exit(-1)
	}
	file := os.Args[1]
	lists := global.Load(file)
	fmt.Println("--- Puzzle 1 ---")
	one(lists)
	fmt.Println("--- Puzzle 2 ---")
	two(lists)
}
