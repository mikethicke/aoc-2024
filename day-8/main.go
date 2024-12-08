package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coords struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	ants := make(map[byte][]coords)
	scanner := bufio.NewScanner(file)
	lineCounter := 0
	width := 0
	for scanner.Scan() {
		locs := scanner.Bytes()
		width = len(locs)
		for loc, char := range locs {
			if char != '.' {
				ants[char] = append(ants[char], coords{loc, lineCounter})
			}
		}
		lineCounter++
	}
	height := lineCounter

	//Part 1
	antinodes := make(map[coords]bool)
	for key := range ants {
		for i := 0; i < len(ants[key]); i++ {
			for j := 0; j < len(ants[key]); j++ {
				if i == j {
					continue
				}
				xDist := ants[key][i].x - ants[key][j].x
				yDist := ants[key][i].y - ants[key][j].y
				antinode := coords{
					x: ants[key][i].x + xDist,
					y: ants[key][i].y + yDist,
				}
				if antinode.x >= 0 && antinode.x < width && antinode.y >= 0 && antinode.y < height {
					antinodes[antinode] = true
				}
			}
		}
	}
	total := len(antinodes)
	fmt.Println("Part 1:", total)

	//Part 2
	total = 0
	antinodes = make(map[coords]bool)
	for key := range ants {
		for i := 0; i < len(ants[key]); i++ {
			for j := 0; j < len(ants[key]); j++ {
				if i == j {
					continue
				}
				xDist := ants[key][i].x - ants[key][j].x
				yDist := ants[key][i].y - ants[key][j].y
				gcd := GCD(xDist, yDist)
				xDist = xDist / gcd
				yDist = yDist / gcd
				var x, y int
				x = ants[key][i].x
				y = ants[key][i].y
				for x >= 0 && x < width && y >= 0 && y < height {
					antinode := coords{x, y}
					antinodes[antinode] = true
					x = x + xDist
					y = y + yDist
				}
				x = ants[key][i].x
				y = ants[key][i].y
				for x >= 0 && x < width && y >= 0 && y < height {
					antinode := coords{x, y}
					antinodes[antinode] = true
					x = x - xDist
					y = y - yDist
				}
			}
		}
	}
	total = len(antinodes)
	fmt.Println("Part 2:", total)
}

func GCD(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
