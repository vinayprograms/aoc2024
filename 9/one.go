package main

import (
	"fmt"
	"os"
)

func one(lines []string) {
	line := merge(lines)
	blocks := metadata(line)
	//fmt.Println(blocks)
	diskmap := representation(blocks)
	//fmt.Println(diskmap)
	defrag(&diskmap)
	//fmt.Println(diskmap)
	fmt.Println("CHECKSUM:", checksum(diskmap))
}

func defrag(diskmap *[]string) {
	length := int(len(*diskmap))
	var reverseIndex int = length - 1
	//fmt.Println(length, reverseIndex)

	// output defrag output to a log file instead of STDOUT
	file, err := os.Create("defrag.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := int(0); i < length; i++ {
		if i >= reverseIndex { // Stop once indices overlap or cross over, since all blocks after that are known to be free spaces
			fmt.Fprintln(file, "STOPPED @", reverseIndex, "/", i, "LEN:", len(*diskmap), "VALUE @ stopped:", (*diskmap)[reverseIndex+1])
			break
		}
		if (*diskmap)[i] == "." {
			fmt.Fprintln(file, "EMPTY @", i)
			if (*diskmap)[reverseIndex] == "." {
				fmt.Fprintln(file, "REVERSE EMPTY @", reverseIndex)
				reverseIndex = searchBackward(diskmap, reverseIndex)
				fmt.Fprintln(file, "REVERSE MOVED BACK TO", reverseIndex)
				if i >= reverseIndex {
					fmt.Fprintln(file, "OVERSHOOT FROM MOVING BACK. I'M DONE DEFRAGGING")
					break
				}
			}
			fmt.Fprintf(file, "%s @ %d <- %s @ %d\n", (*diskmap)[i], i, (*diskmap)[reverseIndex], reverseIndex)
			(*diskmap)[i] = (*diskmap)[reverseIndex]
			(*diskmap)[reverseIndex] = "."
			reverseIndex--
		}
	}
}

func searchBackward(diskmap *[]string, idx int) int {
	for i := idx; i >= 0; i-- {
		if (*diskmap)[i] != "." {
			return i
		}
	}
	return 0
}
