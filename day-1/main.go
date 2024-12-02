package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	const numLines = 1000

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open input file: " + fileName)
	}
	defer file.Close()

	difference := int(0)
	lines := 0
	firstList := make([]int, numLines)
	secondList := make([]int, numLines)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}
		re := regexp.MustCompile(`\d+`)
		numbers := re.FindAll(line, 2)
		num1Int, err := strconv.Atoi(string(numbers[0]))
		if err != nil {
			log.Fatal("Failed to convert num1 to int")
		}
		firstList[lines] = num1Int
		num2Int, err := strconv.Atoi(string(numbers[1]))
		if err != nil {
			log.Fatal("Failed to convert num2 to int")
		}
		secondList[lines] = num2Int
		lines++
	}
	sort.Ints(firstList)
	sort.Ints(secondList)
	for i := 0; i < numLines; i++ {
		if firstList[i] > secondList[i] {
			difference += firstList[i] - secondList[i]
		} else {
			difference += secondList[i] - firstList[i]
		}
	}
	fmt.Println("Total lines: ", lines)
	fmt.Println("Total difference: ", difference)

	firstIndex := 0
	secondIndex := 0
	similarity := 0
	firstCurrent := 0
	similarityCurrent := 0
	for {
		if secondList[secondIndex] < firstList[firstIndex] {
			secondIndex++
			continue
		}
		if secondList[secondIndex] == firstList[firstIndex] {
			similarity += firstList[firstIndex]
			similarityCurrent = similarity
			firstCurrent = firstList[firstIndex]
			secondIndex++
			continue
		}
		firstIndex++
		if firstIndex >= numLines || secondIndex >= numLines {
			break
		}
		if firstList[firstIndex] == firstCurrent {
			similarity += similarityCurrent
		}
	}
	fmt.Println("Similarity: ", similarity)
}
