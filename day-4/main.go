package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	const fileName = "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()

	chars := make([][]byte, 140)
	lineNum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars[lineNum] = append([]byte(nil), scanner.Bytes()...)
		lineNum++
	}
	fmt.Println(lineNum)
	xmasCount := 0
	rows := lineNum
	cols := len(chars[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if chars[row][col] == 'X' {
				//Left
				if col > 2 &&
					chars[row][col-1] == 'M' &&
					chars[row][col-2] == 'A' &&
					chars[row][col-3] == 'S' {
					xmasCount++
				}
				//Right
				if col < cols-3 &&
					chars[row][col+1] == 'M' &&
					chars[row][col+2] == 'A' &&
					chars[row][col+3] == 'S' {
					xmasCount++
				}
				//Up
				if row > 2 &&
					chars[row-1][col] == 'M' &&
					chars[row-2][col] == 'A' &&
					chars[row-3][col] == 'S' {
					xmasCount++
				}
				//Down
				if row < rows-3 &&
					chars[row+1][col] == 'M' &&
					chars[row+2][col] == 'A' &&
					chars[row+3][col] == 'S' {
					xmasCount++
				}
				//NE
				if row > 2 &&
					col < cols-3 &&
					chars[row-1][col+1] == 'M' &&
					chars[row-2][col+2] == 'A' &&
					chars[row-3][col+3] == 'S' {
					xmasCount++
				}
				//SE
				if row < rows-3 &&
					col < cols-3 &&
					chars[row+1][col+1] == 'M' &&
					chars[row+2][col+2] == 'A' &&
					chars[row+3][col+3] == 'S' {
					xmasCount++
				}
				//SW
				if row < rows-3 &&
					col > 2 &&
					chars[row+1][col-1] == 'M' &&
					chars[row+2][col-2] == 'A' &&
					chars[row+3][col-3] == 'S' {
					xmasCount++
				}
				//NW
				if row > 2 &&
					col > 2 &&
					chars[row-1][col-1] == 'M' &&
					chars[row-2][col-2] == 'A' &&
					chars[row-3][col-3] == 'S' {
					xmasCount++
				}
			}
		}
	}
	fmt.Println("XMAS count: ", xmasCount)
	xmasCount = 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if chars[row][col] == 'A' {
				if row < 1 ||
					col < 1 ||
					row > rows-2 ||
					col > cols-2 {
					continue
				}
				if chars[row-1][col-1] == 'M' &&
					chars[row-1][col+1] == 'M' &&
					chars[row+1][col-1] == 'S' &&
					chars[row+1][col+1] == 'S' {
					xmasCount++
				}
				if chars[row-1][col-1] == 'S' &&
					chars[row-1][col+1] == 'S' &&
					chars[row+1][col-1] == 'M' &&
					chars[row+1][col+1] == 'M' {
					xmasCount++
				}
				if chars[row-1][col-1] == 'M' &&
					chars[row-1][col+1] == 'S' &&
					chars[row+1][col-1] == 'M' &&
					chars[row+1][col+1] == 'S' {
					xmasCount++
				}
				if chars[row-1][col-1] == 'S' &&
					chars[row-1][col+1] == 'M' &&
					chars[row+1][col-1] == 'S' &&
					chars[row+1][col+1] == 'M' {
					xmasCount++
				}
			}
		}
	}
	fmt.Println("X-MAS count: ", xmasCount)
}
