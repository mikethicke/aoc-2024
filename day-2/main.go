package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	const fileName = "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()
	lines := make([][]int, 5, 5)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numStr := strings.Fields(line)
		nums := make([]int, len(numStr))
		for i, s := range numStr {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Failed to parse line")
			}
			nums[i] = n
		}
		lines = append(lines, nums)
	}

	safeCount := 0
	for _, l := range lines {
		asc := true
		failIndex := -1
		if len(l) == 0 {
			continue
		}
		ln := make([]int, len(l))
		copy(ln, l)
	trialLoop:
		for trial := 0; trial < 4; trial++ {
			if trial == 1 {
				ln = append([]int{}, l[1:]...)
			}
			if trial == 2 {
				ln = append([]int{}, l[:failIndex]...)
				ln = append(ln, l[failIndex+1:]...)
			}
			if trial == 3 {
				if failIndex+3 > len(l) {
					ln = l[:len(l)-1]
				} else {
					ln = append(l[:failIndex+1], l[failIndex+2:]...)
				}
			}
			for i := 0; i < len(ln); i++ {
				if i == 0 {
					asc = ln[1] > ln[0]
				}
				if i == len(ln)-1 {
					safeCount++
					break trialLoop
				}
				if (asc && ln[i] > ln[i+1]) ||
					(!asc && ln[i] < ln[i+1]) ||
					(ln[i] == ln[i+1]) ||
					(ln[i+1]-ln[i] > 3) ||
					(ln[i]-ln[i+1] > 3) {
					if trial == 0 {
						failIndex = i
					}
					break
				}
			}
		}
	}
	println("Safe count: ", safeCount)
}
