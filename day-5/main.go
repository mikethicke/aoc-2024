package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rules := make(map[int][]int)
	ruleMode := true
	sum := 0
	fixedSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if ruleMode {
			if line == "" {
				ruleMode = false
				continue
			}
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				log.Fatal("Error parsing rule line")
			}
			first, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal("Error converting first rule part to Int")
			}
			second, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("Error converting second rule part to Int")
			}
			rules[first] = append(rules[first], second)
		} else {
			if line == "" {
				continue
			}
			parts := strings.Split(line, ",")
			previousParts := make(map[int]bool)
			correctOrder := true
		partsLoop:
			for _, part := range parts {
				partNum, err := strconv.Atoi(part)
				if err != nil {
					log.Fatal("Error converting " + part + "to Int")
				}
				for _, forbiddenNum := range rules[partNum] {
					if previousParts[forbiddenNum] {
						correctOrder = false
						break partsLoop
					}
				}
				previousParts[partNum] = true
			}
			if correctOrder {
				middle := len(parts) / 2
				middleNum, err := strconv.Atoi(parts[middle])
				if err != nil {
					log.Fatal("Error converting middle part to Int")
				}
				sum += middleNum
			} else {
				fixed := false
				var previousParts2 map[int]int
				for {
				partsLoop2:
					for i, part := range parts {
						if i == 0 {
							previousParts2 = make(map[int]int)
						}
						partNum, err := strconv.Atoi(part)
						if err != nil {
							log.Fatal("Error converting " + part + "to Int")
						}
						fixed = true
						for _, forbiddenNum := range rules[partNum] {
							prevIndex, exists := previousParts2[forbiddenNum]
							if exists {
								parts[i] = strconv.Itoa(forbiddenNum)
								parts[prevIndex] = strconv.Itoa(partNum)
								fixed = false
								break partsLoop2
							}
						}
						previousParts2[partNum] = i
					}
					if fixed {
						break
					}
				}
				//fmt.Println("Fixed line: ", parts)
				middle := len(parts) / 2
				middleNum, err := strconv.Atoi(parts[middle])
				if err != nil {
					log.Fatal("Error converting middle part to Int")
				}
				fixedSum += middleNum
			}

		}
	}
	fmt.Println("Part 1 sum: ", sum)
	fmt.Println("Part 2 sum: ", fixedSum)
}
