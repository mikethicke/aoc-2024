package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

type trailhead struct {
	coords
	score int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening input file")
	}
	scanner := bufio.NewScanner(file)
	var topo [][]int
	var trailheads []trailhead
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		var row []int
		for n, char := range chars {
			num, _ := strconv.Atoi(char)
			row = append(row, num)
			if num == 0 {
				trailheads = append(trailheads, trailhead{coords{n, lineNumber}, 0})
			}
		}
		topo = append(topo, row)
		lineNumber++
	}

	//Part 1
	score := 0
	part2Score := 0
	for _, trailhead := range trailheads {
		endpoints := make(map[coords]bool)
		currentPos := make([]coords, 1, 10)
		currentPos[0] = coords{trailhead.x, trailhead.y}
		for step := 1; step <= 9; step++ {
			nextSteps := make([]coords, 0, 10)
			for _, pos := range currentPos {
				if pos.x > 0 && topo[pos.y][pos.x-1] == step {
					nextSteps = append(nextSteps, coords{pos.x - 1, pos.y})
				}
				if pos.x < len(topo[pos.y])-1 && topo[pos.y][pos.x+1] == step {
					nextSteps = append(nextSteps, coords{pos.x + 1, pos.y})
				}
				if pos.y > 0 && topo[pos.y-1][pos.x] == step {
					nextSteps = append(nextSteps, coords{pos.x, pos.y - 1})
				}
				if pos.y < len(topo)-1 && topo[pos.y+1][pos.x] == step {
					nextSteps = append(nextSteps, coords{pos.x, pos.y + 1})
				}
			}
			if len(nextSteps) == 0 {
				break
			}
			currentPos = make([]coords, len(nextSteps))
			copy(currentPos, nextSteps)
			if step == 9 {
				part2Score += len(currentPos)
				for _, pos := range currentPos {
					endpoints[pos] = true
				}
			}
		}
		score += len(endpoints)
	}
	fmt.Println("Part 1: ", score)
	fmt.Println("Part 2: ", part2Score)

}
