package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type fileBlock struct {
	start int
	end   int
	ID    int
}

func (fb fileBlock) size() int {
	return fb.end - fb.start + 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening input file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var filesystem []int
	filesystemIndex := 0
	var freespace []int
	fileID := 0
	isFileBlock := true
	var lastFileIndex int
	numFileBlocks := 0
	var freeBlocks []fileBlock
	var fileBlocks []fileBlock
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		encodedNumber := int(char - '0')
		startIndex := filesystemIndex
		for {
			if encodedNumber == 0 {
				break
			}
			if encodedNumber < 0 || encodedNumber > 9 {
				break
			}
			if isFileBlock {
				filesystem = append(filesystem, fileID)
				lastFileIndex = filesystemIndex
				numFileBlocks++
			} else {
				filesystem = append(filesystem, -1)
				freespace = append(freespace, filesystemIndex)
			}
			filesystemIndex++
			if filesystemIndex-startIndex == encodedNumber {
				break
			}
		}
		if isFileBlock {
			fileBlocks = append(fileBlocks, fileBlock{startIndex, filesystemIndex - 1, fileID})
			fileID++
		} else {
			freeBlocks = append(freeBlocks, fileBlock{startIndex, filesystemIndex - 1, -1})
		}
		isFileBlock = !isFileBlock
	}

	//Part 1
	nextFreeIndex := 0
	for i := lastFileIndex; i > 0; i-- {
		if filesystem[i] == -1 {
			continue
		}
		filesystem[freespace[nextFreeIndex]] = filesystem[i]
		filesystem[i] = -1
		//fmt.Println(filesystem)
		if i <= freespace[nextFreeIndex] {
			freespace[nextFreeIndex] = i
		} else {
			nextFreeIndex++
		}
		if nextFreeIndex >= len(freespace) {
			break
		}
		if i <= numFileBlocks {
			break
		}
	}
	checksum := 0
	for i := 0; filesystem[i] != -1; i++ {
		checksum += i * filesystem[i]
	}
	fmt.Println("Part 1: ", checksum)

	//Part 2
	for i := len(fileBlocks) - 1; i >= 0; i-- {
		for j := 0; j < len(freeBlocks); j++ {
			if fileBlocks[i].size() <= freeBlocks[j].size() && fileBlocks[i].start > freeBlocks[j].start {
				fbSize := fileBlocks[i].size()
				fileBlocks[i].start = freeBlocks[j].start
				fileBlocks[i].end = fileBlocks[i].start + fbSize - 1
				freeBlocks[j].start = fileBlocks[i].end + 1
				break
			}
		}
	}
	checksum = 0
	for _, fb := range fileBlocks {
		for i := fb.start; i <= fb.end; i++ {
			checksum += fb.ID * i
		}
	}
	fmt.Println("Part 2: ", checksum)
}
