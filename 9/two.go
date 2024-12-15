package main

import (
	"fmt"
)

func two(lines []string) {
	line := merge(lines)
	blocks := metadata(line)
	diskmap := representation(blocks)
	//fmt.Println(diskmap, len(diskmap))
	defragged := defragFiles(diskmap)
	//fmt.Println(defragged)
	fmt.Println("CHECKSUM:", checksum(defragged))
}

func defragFiles(diskmap []string) []string {
	fileEnd := len(diskmap) - 1
	fileStart := len(diskmap) - 1
	id := ""
	searchSpace := diskmap
	for len(searchSpace) > 0 {
		// Start from end of map, move left and find the first file
		id, fileStart, fileEnd = findFile(searchSpace, fileEnd)
		// Find the lowest index, before the file start index, to accomodate it.
		freeStart, freeEnd := findFreeSpace(searchSpace[:fileStart], fileEnd-fileStart+1)
		if freeStart == -1 {
			//fmt.Println("CANNOT FIND FREE SPACE FOR FILEID", id)
		} else if freeEnd-freeStart >= fileEnd-fileStart {
			//fmt.Println("FREE SPACE @", freeStart, "-", freeEnd)
			diskmap = swap(diskmap, id, freeStart, freeEnd, fileStart, fileEnd)
			//fmt.Println(diskmap)
		} //else {
		//fmt.Println("CANNOT FIND FREE SPACE FOR FILEID", id)
		//}
		fileEnd = fileStart - 1
		fileStart--
		searchSpace = searchSpace[:fileStart+1]
	}
	return diskmap
}

func findFile(diskmap []string, end int) (string, int, int) {
	//fmt.Println("STARTING SEARCH FOR END @", end)
	for diskmap[end] == "." { // search backward for block with file ID
		end--
	}
	//fmt.Println("FILE END FOUND:", end)
	start := end
	for diskmap[start] == diskmap[end] {
		start--
		if start < 0 {
			return diskmap[0], 0, end
		}
	}
	//fmt.Println("FILE START FOUND:", start+1)
	file := diskmap[start+1 : end+1]
	//fmt.Println("FILE:", diskmap[start+1:end+1])
	return file[0], start + 1, end
}

func findFreeSpace(diskmap []string, size int) (int, int) {
	fstart := -1
	fend := -1
	//fmt.Println("FREE SPACE REQD.:", size)
	for i := 0; i < len(diskmap); i++ {
		if diskmap[i] == "." {
			fstart = i
			//fmt.Println("FREE START:", fstart)
			for j := i; j < len(diskmap); j++ {
				if diskmap[j] == "." {
					fend = j
					//fmt.Println("FREE END:", fend)
					if fend-fstart+1 == size {
						//fmt.Println("FOUND FREE SPACE:", fstart, "-", fend)
						return fstart, fend
					}
				} else {
					if fend-fstart < size {
						i = fend
						break
					}
				}
			}
		}
	}
	return fstart, fend
}

func swap(diskmap []string, id string, frstart int, frend int, flstart int, flend int) []string {
	for i := frstart; i <= frend; i++ {
		diskmap[i] = id
	}
	for i := flstart; i <= flend; i++ {
		diskmap[i] = "."
	}

	return diskmap
}
