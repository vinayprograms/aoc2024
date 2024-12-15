package main

import (
	"fmt"
	"strconv"
)

// Merge lines into a single piece of text
func merge(lines []string) string {
	text := ""
	for _, str := range lines {
		text += str
	}
	return text
}

type blockType string

const (
	freespace blockType = "free"
	file      blockType = "file"
)

type block struct {
	length int
	typ    blockType
	id     string
}

func metadata(layout string) (blocks []block) {
	idCounter := 0
	for i := int(0); i < int(len(layout)); i++ {
		if length, err := strconv.Atoi(string(layout[i])); err != nil {
			panic(err)
		} else {
			if length == 0 {
				continue
			}
			blk := block{length: length}
			if i%2 == 0 {
				blk.typ = file
				blk.id = fmt.Sprintf("%d", idCounter)
				idCounter++
			} else {
				blk.typ = freespace
			}
			blocks = append(blocks, blk)
		}
	}
	return
}

func representation(blocks []block) (diskMap []string) {
	for _, blk := range blocks {
		var char string
		if blk.typ == file {
			char = blk.id
		} else {
			char = "."
		}
		for i := 0; i < blk.length; i++ {
			diskMap = append(diskMap, char)
		}
	}
	return diskMap
}

func checksum(diskmap []string) (sum int) {
	sum = int(0)
	for i := int(0); i < int(len(diskmap)); i++ {
		if diskmap[i] == "." {
			continue
		}
		id, _ := strconv.Atoi(diskmap[i])
		sum += i * id
	}
	return sum
}
