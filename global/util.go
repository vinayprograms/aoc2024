package global

import (
	"os"
	"strings"
)

// Read contents from input file and return as a set of lines
func Load(file string) []string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	return lines[:len(lines)-1]
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sum(list []int) int {
	var total = 0
	for i := range list {
		total += list[i]
	}
	return total
}
