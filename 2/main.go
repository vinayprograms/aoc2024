package main

import (
	"fmt"
	"global"
	"os"
)

func main() {
	fmt.Println("DAY 2: Red Nose Reports")
	if len(os.Args) != 2 {
		fmt.Println(`Must supply input file.
		File must contain N lines with each line containing a space-separated list of numbers`)
		os.Exit(-1)
	}
	file := os.Args[1]
	reports := global.Load(file)
	fmt.Println("--- Puzzle 1 ---")
	one(reports)
	fmt.Println("--- Puzzle 2 ---")
	two(reports)
}
