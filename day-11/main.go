package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const startingBlinks = 75

type blinkStart struct {
	num             int
	remainingBlinks int
}

var blinkResults map[blinkStart]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Failed to open input file")
	}
	currentLine := make([]int, 0, 0)
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		numStrings := strings.Fields(line)
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal("Error converting numbers to ints")
			}
			currentLine = append(currentLine, num)
		}
	}
	blinkResults = make(map[blinkStart]int)
	total := 0
	for _, num := range currentLine {
		total += calcBlinkResult(blinkStart{num, startingBlinks})
	}
	fmt.Println("Final line length: ", total)
}

func calcBlinkResult(start blinkStart) int {
	var nextLine []int
	numStr := strconv.Itoa(start.num)
	if start.num == 0 {
		nextLine = append(nextLine, 1)
	} else if len(numStr)%2 == 0 {
		first, _ := strconv.Atoi(numStr[:len(numStr)/2])
		second, _ := strconv.Atoi(numStr[len(numStr)/2:])
		nextLine = append(nextLine, first, second)
	} else {
		nextLine = append(nextLine, start.num*2024)
	}
	if start.remainingBlinks == 1 {
		result := len(nextLine)
		blinkResults[start] = result
		return result
	}

	nextResult := 0
	for _, nextStartingNum := range nextLine {
		nextStart := blinkStart{nextStartingNum, start.remainingBlinks - 1}
		cachedResult := blinkResults[nextStart]
		if cachedResult != 0 {
			nextResult += cachedResult
			continue
		}
		nextResult += calcBlinkResult(nextStart)
	}
	blinkResults[start] = nextResult
	return nextResult
}
