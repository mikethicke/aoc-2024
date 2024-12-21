package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const BoardWidth = 101
const BoardHeight = 103
const NumSeconds = 100

type coords struct {
	x int
	y int
}

type robot struct {
	position coords
	velocity coords
}

func main() {
	var robots []robot
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
		match := re.FindStringSubmatch(line)
		var position coords
		var velocity coords
		position.x, _ = strconv.Atoi(match[1])
		position.y, _ = strconv.Atoi(match[2])
		velocity.x, _ = strconv.Atoi(match[3])
		velocity.y, _ = strconv.Atoi(match[4])
		newRobot := robot{position, velocity}
		robots = append(robots, newRobot)
	}
	finalState := moveRobots(robots, NumSeconds)
	safetyFactor := robotCensus(finalState)
	//printBoard(finalState)
	fmt.Println("Part 1: ", safetyFactor)
	for i := 0; i < 10000; i++ {
		state := moveRobots(robots, i)
		sf := robotCensus(state)
		threshold := math.Floor(math.Pow(float64(len(robots)/4.0), 4.0)) / 2
		//if float64(sf) < threshold {
		if sf == 112847865 {
			fmt.Println("SF: ", sf, "Threshold: ", threshold, "Seconds: ", i)
			printBoard(state)
		}
	}
}

func printBoard(robots []robot) {
	board := make([][]byte, BoardHeight)
	for i := 0; i < BoardHeight; i++ {
		board[i] = make([]byte, BoardWidth)
		for j := 0; j < BoardWidth; j++ {
			board[i][j] = ' '
		}
	}
	for _, tr := range robots {
		board[tr.position.y][tr.position.x] = '*'
	}
	for _, row := range board {
		fmt.Println(string(row))
	}
}

func moveRobots(robots []robot, seconds int) []robot {
	var endState []robot
	for _, thisRobot := range robots {
		var endPosition coords
		endPosition.x = (thisRobot.position.x + thisRobot.velocity.x*seconds) % BoardWidth
		if endPosition.x < 0 {
			endPosition.x += BoardWidth
		}
		endPosition.y = (thisRobot.position.y + thisRobot.velocity.y*seconds) % BoardHeight
		if endPosition.y < 0 {
			endPosition.y += BoardHeight
		}
		newRobot := robot{
			position: endPosition,
			velocity: thisRobot.velocity,
		}
		endState = append(endState, newRobot)
	}
	return endState
}

func robotCensus(robots []robot) int {
	hBoundaryPosition := BoardWidth / 2
	vBoundaryPosition := BoardHeight / 2
	nwQuad := 0
	neQuad := 0
	swQuad := 0
	seQuad := 0

	for _, tr := range robots {
		if tr.position.x > hBoundaryPosition && tr.position.y > vBoundaryPosition {
			nwQuad++
		}
		if tr.position.x < hBoundaryPosition && tr.position.y > vBoundaryPosition {
			neQuad++
		}
		if tr.position.x > hBoundaryPosition && tr.position.y < vBoundaryPosition {
			swQuad++
		}
		if tr.position.x < hBoundaryPosition && tr.position.y < vBoundaryPosition {
			seQuad++
		}
	}

	return nwQuad * neQuad * swQuad * seQuad
}
