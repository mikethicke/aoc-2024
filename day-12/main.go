package main

import (
	"bufio"
	"fmt"
	"os"
)

type coords struct {
	y int
	x int
}

var visited map[coords]bool
var garden [][]byte

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		newLine := make([]byte, len(line))
		copy(newLine, line)
		garden = append(garden, newLine)
	}
	totalPrice := 0
	part2Price := 0
	gardenScannerPosition := coords{0, 0}
	visited = make(map[coords]bool)
	for gardenScannerPosition.y = 0; gardenScannerPosition.y < len(garden); gardenScannerPosition.y++ {
		for gardenScannerPosition.x = 0; gardenScannerPosition.x < len(garden[gardenScannerPosition.y]); gardenScannerPosition.x++ {
			if visited[gardenScannerPosition] {
				continue
			}
			area, perimiter, corners := mapRegion(gardenScannerPosition)
			totalPrice += area * perimiter
			part2Price += area * corners
		}
	}
	fmt.Println("Part 1: ", totalPrice)
	fmt.Println("Part 2: :", part2Price)
}

func mapRegion(startPosition coords) (area, perimiter, corners int) {
	area, perimiter, corners = 0, 0, 0
	regionChar := garden[startPosition.y][startPosition.x]
	positionQueue := make([]coords, 1)
	positionQueue[0] = startPosition
	for len(positionQueue) > 0 {
		currentPosition := positionQueue[0]
		positionQueue = positionQueue[1:]
		if visited[currentPosition] {
			continue
		}
		visited[currentPosition] = true
		area++
		perimiter += 4
		if currentPosition.y > 0 && garden[currentPosition.y-1][currentPosition.x] == regionChar {
			perimiter--
			if !visited[coords{currentPosition.y - 1, currentPosition.x}] {
				positionQueue = append(positionQueue, coords{currentPosition.y - 1, currentPosition.x})
			}
		}
		if currentPosition.y < len(garden)-1 && garden[currentPosition.y+1][currentPosition.x] == regionChar {
			perimiter--
			if !visited[coords{currentPosition.y + 1, currentPosition.x}] {
				positionQueue = append(positionQueue, coords{currentPosition.y + 1, currentPosition.x})
			}
		}
		if currentPosition.x > 0 && garden[currentPosition.y][currentPosition.x-1] == regionChar {
			perimiter--
			if !visited[coords{currentPosition.y, currentPosition.x - 1}] {
				positionQueue = append(positionQueue, coords{currentPosition.y, currentPosition.x - 1})
			}
		}
		if currentPosition.x < len(garden[currentPosition.y])-1 && garden[currentPosition.y][currentPosition.x+1] == regionChar {
			perimiter--
			if !visited[coords{currentPosition.y, currentPosition.x + 1}] {
				positionQueue = append(positionQueue, coords{currentPosition.y, currentPosition.x + 1})
			}
		}

		//Outie corners
		if (currentPosition.y == 0 || garden[currentPosition.y-1][currentPosition.x] != regionChar) &&
			(currentPosition.x == len(garden[currentPosition.y])-1 || garden[currentPosition.y][currentPosition.x+1] != regionChar) {
			corners++
		}
		if (currentPosition.y == len(garden)-1 || garden[currentPosition.y+1][currentPosition.x] != regionChar) &&
			(currentPosition.x == len(garden[currentPosition.y])-1 || garden[currentPosition.y][currentPosition.x+1] != regionChar) {
			corners++
		}
		if (currentPosition.y == 0 || garden[currentPosition.y-1][currentPosition.x] != regionChar) &&
			(currentPosition.x == 0 || garden[currentPosition.y][currentPosition.x-1] != regionChar) {
			corners++
		}
		if (currentPosition.y == len(garden)-1 || garden[currentPosition.y+1][currentPosition.x] != regionChar) &&
			(currentPosition.x == 0 || garden[currentPosition.y][currentPosition.x-1] != regionChar) {
			corners++
		}

		//Innie corners
		if currentPosition.y < len(garden)-1 &&
			currentPosition.x > 0 &&
			garden[currentPosition.y+1][currentPosition.x] != regionChar &&
			garden[currentPosition.y+1][currentPosition.x-1] == regionChar &&
			garden[currentPosition.y][currentPosition.x-1] == regionChar {
			corners++
		}
		if currentPosition.y < len(garden)-1 &&
			currentPosition.x < len(garden[currentPosition.y])-1 &&
			garden[currentPosition.y+1][currentPosition.x] != regionChar &&
			garden[currentPosition.y+1][currentPosition.x+1] == regionChar &&
			garden[currentPosition.y][currentPosition.x+1] == regionChar {
			corners++
		}
		if currentPosition.y > 0 &&
			currentPosition.x < len(garden[currentPosition.y])-1 &&
			garden[currentPosition.y-1][currentPosition.x] != regionChar &&
			garden[currentPosition.y-1][currentPosition.x+1] == regionChar &&
			garden[currentPosition.y][currentPosition.x+1] == regionChar {
			corners++
		}
		if currentPosition.y > 0 &&
			currentPosition.x > 0 &&
			garden[currentPosition.y-1][currentPosition.x] != regionChar &&
			garden[currentPosition.y-1][currentPosition.x-1] == regionChar &&
			garden[currentPosition.y][currentPosition.x-1] == regionChar {
			corners++
		}

	}
	return area, perimiter, corners
}
