package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	const fileName = "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		mulRe, err := regexp.Compile(`mul\((\d+)\,(\d+)\)`)
		if err != nil {
			log.Fatal("Failed to compile mul regex")
		}
		matches := mulRe.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			first, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal("Failed to parse first")
			}
			second, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal("Failed to parse second")
			}
			sum += first * second
		}
	}
	fmt.Println("Sum Part 1: ", sum)

	file.Close()

	file, err = os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()

	scanner2 := bufio.NewScanner(file)
	do := 1
	sum2 := 0
	for scanner2.Scan() {
		line := scanner2.Text()
		mulRe := regexp.MustCompile(`mul\((\d+)\,(\d+)\)`)
		doRe := regexp.MustCompile(`do\(\)`)
		dontRe := regexp.MustCompile(`don\'t\(\)`)
		mulMatches := mulRe.FindAllStringSubmatchIndex(line, -1)
		doMatches := doRe.FindAllStringIndex(line, -1)
		dontMatches := dontRe.FindAllStringIndex(line, -1)
		toggler := make([]int, len(line))
		for _, doMatch := range doMatches {
			toggler[doMatch[0]] = 1
		}
		for _, dontMatch := range dontMatches {
			toggler[dontMatch[0]] = -1
		}
		j := 0
		for _, mIndeces := range mulMatches {
			start := mIndeces[0]
			for {
				if toggler[j] != 0 {
					do = toggler[j]
				}
				if j == start {
					break
				}
				j++
			}
			if do == 1 {
				first, err := strconv.Atoi(line[mIndeces[2]:mIndeces[3]])
				if err != nil {
					log.Fatal("Error parsing first str: ", err)
				}
				second, err := strconv.Atoi(line[mIndeces[4]:mIndeces[5]])
				if err != nil {
					log.Fatal("Error parsing second str", err)
				}
				sum2 += first * second
			}
		}
	}
	fmt.Println("Sum part 2: ", sum2)
}
