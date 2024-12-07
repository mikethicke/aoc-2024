package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	West  byte = 1 << 0 // 0b0001
	East  byte = 1 << 1 // 0b0010
	North byte = 1 << 2 // 0b0100
	South byte = 1 << 3 // 0b1000

)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}
	world := make([][]byte, 0, 130)
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := append([]byte(nil), scanner.Bytes()...)
		world = append(world, line)
		lineNum++
	}

	// Make a clean copy of world for Part 2
	world2 := make([][]byte, len(world))
	for i := range world {
		world2[i] = make([]byte, len(world[i]))
		copy(world2[i], world[i])
	}

	//fmt.Println(string(world[0]))
	lineLength := len(world[0])
	totalLines := lineNum
	fmt.Println(lineLength, totalLines)
	//Find starting position
	guardX := -1
	guardY := -1
findLoop:
	for y := 0; y < totalLines; y++ {
		for x := 0; x < lineLength; x++ {
			if world[y][x] == '^' {
				guardX = x
				guardY = y
				break findLoop
			}
		}
	}
	fmt.Println("Starting guard pos: ", guardY, ", ", guardX)
	startX := guardX
	startY := guardY

	// Part 1
	visitedSquares := 1
	guardDir := 'N'
	for {
		if world[guardY][guardX] == '.' {
			visitedSquares++
			world[guardY][guardX] = 'X'
		}
		var nextX, nextY int
		switch guardDir {
		case 'N':
			nextX = guardX
			nextY = guardY - 1
		case 'E':
			nextX = guardX + 1
			nextY = guardY
		case 'S':
			nextX = guardX
			nextY = guardY + 1
		case 'W':
			nextX = guardX - 1
			nextY = guardY
		}
		if nextX < 0 || nextY < 0 || nextX == lineLength || nextY == totalLines {
			break
		}
		if world[nextY][nextX] == '#' {
			switch guardDir {
			case 'N':
				guardDir = 'E'
			case 'E':
				guardDir = 'S'
			case 'S':
				guardDir = 'W'
			case 'W':
				guardDir = 'N'
			}
			continue
		}
		guardX = nextX
		guardY = nextY
		//printMap(world)
		//fmt.Println("()()()()()()()()")
	}
	//printMap(world)
	fmt.Println("Visited squares: ", visitedSquares)

	// Part 2
	guardDir = 'N'
	guardX = startX
	guardY = startY
	loopCounter := 0
	//world2[guardY][guardX] = North
	for bpY := 0; bpY < totalLines; bpY++ {
		for bpX := 0; bpX < lineLength; bpX++ {
			// Copy world2 back to world
			for i := range world2 {
				world[i] = make([]byte, len(world2[i]))
				copy(world[i], world2[i])
			}

			if world[bpY][bpX] == '#' {
				continue
			}
			if bpY == startY && bpX == startX {
				continue
			}
			world[bpY][bpX] = '#'
			guardDir = 'N'
			guardX = startX
			guardY = startY
		travelLoop:
			for {
				if world[guardY][guardX] == '.' || world[guardY][guardX] == '^' {
					switch guardDir {
					case 'N':
						world[guardY][guardX] = North
					case 'E':
						world[guardY][guardX] = East
					case 'S':
						world[guardY][guardX] = South
					case 'W':
						world[guardY][guardX] = West
					}
				} else {
					switch guardDir {
					case 'N':
						if world[guardY][guardX]&North != 0 {
							loopCounter++
							//fmt.Println(loopCounter)
							//printMap(world)
							break travelLoop
						}
						world[guardY][guardX] = world[guardY][guardX] | North
					case 'E':
						if world[guardY][guardX]&East != 0 {
							loopCounter++
							//fmt.Println(loopCounter)
							break travelLoop
						}
						world[guardY][guardX] = world[guardY][guardX] | East
					case 'S':
						if world[guardY][guardX]&South != 0 {
							loopCounter++
							//fmt.Println(loopCounter)
							break travelLoop
						}
						world[guardY][guardX] = world[guardY][guardX] | South
					case 'W':
						if world[guardY][guardX]&West != 0 {
							loopCounter++
							//fmt.Println(loopCounter)
							break travelLoop
						}
						world[guardY][guardX] = world[guardY][guardX] | West
					}
				}
				var nextX, nextY int
				switch guardDir {
				case 'N':
					nextX = guardX
					nextY = guardY - 1
				case 'E':
					nextX = guardX + 1
					nextY = guardY
				case 'S':
					nextX = guardX
					nextY = guardY + 1
				case 'W':
					nextX = guardX - 1
					nextY = guardY
				}
				if nextX < 0 || nextY < 0 || nextX == lineLength || nextY == totalLines {
					break
				}
				if world[nextY][nextX] == '#' {
					switch guardDir {
					case 'N':
						guardDir = 'E'
					case 'E':
						guardDir = 'S'
					case 'S':
						guardDir = 'W'
					case 'W':
						guardDir = 'N'
					}
					continue
				}
				guardX = nextX
				guardY = nextY
				//printMap(world)
				//fmt.Println("()()()()()()()()")
			}
			//printMap(world)
		}
	}
	fmt.Println("Total loops: ", loopCounter)
}

func printMap(world [][]byte) {
	fmt.Println("()()()()()()()()()()")
	for _, line := range world {
		for _, char := range line {
			switch char {
			case North:
				fmt.Printf("^")
			case East:
				fmt.Printf(">")
			case South:
				fmt.Printf("v")
			case West:
				fmt.Printf("<")
			case North | East, North | South, North | West, East | South, East | West, South | West:
				fmt.Printf("*")
			default:
				fmt.Printf("%c", char)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("()())()()()()(()()()()")
}
