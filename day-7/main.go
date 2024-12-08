package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type trial struct {
	testValue int
	fields    []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var trials []trial
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		testValue, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("Failed to parse test value.")
		}
		opFields := strings.Fields(parts[1])
		opNums := make([]int, len(opFields))
		for i, oF := range opFields {
			opNums[i], _ = strconv.Atoi(oF)
		}
		trials = append(trials, trial{testValue, opNums})
	}

	//Part 1
	total := 0
	for _, currentTrial := range trials {
		//fmt.Println(currentTrial.fields)
		permutations := 1 << (len(currentTrial.fields))
		for perm := 0; perm < permutations; perm++ {
			if testPermutation(currentTrial, perm) {
				total += currentTrial.testValue
				break
			}
		}
	}
	fmt.Println("Part 1: ", total)

	//Part 2
	total = 0
	for _, currentTrial := range trials {
		permutations := generatePermutationsWithConcat(len(currentTrial.fields) - 1)
		for _, perm := range permutations {
			if testPermutationWithConcat(currentTrial, perm) {
				fmt.Println("Success: ", currentTrial.fields)
				total += currentTrial.testValue
				break
			}
		}
	}
	fmt.Println("Part 2: ", total)
}

func testPermutation(currentTrial trial, perm int) bool {
	result := currentTrial.fields[0]
	for position := 1; position < len(currentTrial.fields); position++ {
		//fmt.Println("perm: ", perm, " position: ", position)
		if perm&(1<<position) != 0 {
			result *= currentTrial.fields[position]
		} else {
			result += currentTrial.fields[position]
		}
	}
	//fmt.Println("result: ", result, " testValue: ", currentTrial.testValue)
	return result == currentTrial.testValue
}

// perm has the form [0 0 1 2] where 0 is +, 1 is *, 2 is ||
func testPermutationWithConcat(currentTrial trial, perm []int) bool {
	total := currentTrial.fields[0]
	for fieldIndex := 1; fieldIndex < len(currentTrial.fields); fieldIndex++ {
		switch perm[fieldIndex-1] {
		case 0:
			total += currentTrial.fields[fieldIndex]
		case 1:
			total *= currentTrial.fields[fieldIndex]
		case 2:
			total = concatInts(total, currentTrial.fields[fieldIndex])
		}
	}
	return total == currentTrial.testValue
}

func generatePermutationsWithConcat(length int) [][]int {
	permutations := int(math.Pow(3, float64(length)))
	result := make([][]int, permutations)
	for perm := 0; perm < permutations; perm++ {
		newSlice := make([]int, length)
		value := perm
		for i := 0; i < length; i++ {
			newSlice[i] = value % 3
			value /= 3
		}
		result[perm] = newSlice
	}
	return result
}

func concatInts(first int, second int) int {
	firstStr := strconv.Itoa(first)
	secondStr := strconv.Itoa(second)
	ccStr := firstStr + secondStr
	result, _ := strconv.Atoi(ccStr)
	return result
}
