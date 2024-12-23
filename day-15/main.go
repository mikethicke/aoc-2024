package main

import (
	"bufio"
	"fmt"
	"os"
)

type coords struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var board [][]byte
	var moves []byte
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			lineCopy := make([]byte, len(line))
			copy(lineCopy, line)
			board = append(board, lineCopy)
		} else {
			moves = append(moves, line...)
		}
	}
	//printBoard(board)

	var robot coords
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[j]); i++ {
			if board[j][i] == '@' {
				robot.y = j
				robot.x = i
			}
		}
	}

	for _, move := range moves {
		//printBoard(board)
		var moveDir coords
		switch move {
		case '^':
			moveDir = coords{x: 0, y: -1}
		case '>':
			moveDir = coords{x: 1, y: 0}
		case 'v':
			moveDir = coords{x: 0, y: 1}
		case '<':
			moveDir = coords{x: -1, y: 0}
		}

		nextPos := robot
		var endPos coords
		movable := false
		for {
			nextPos.x = nextPos.x + moveDir.x
			nextPos.y = nextPos.y + moveDir.y
			if board[nextPos.y][nextPos.x] == '.' {
				movable = true
				endPos = nextPos
				break
			}
			if board[nextPos.y][nextPos.x] == '#' {
				movable = false
				break
			}
		}

		nextPos = endPos
		if movable {
			done := false
			for {
				board[nextPos.y][nextPos.x] = board[nextPos.y-moveDir.y][nextPos.x-moveDir.x]
				if board[nextPos.y][nextPos.x] == '@' {
					done = true
					robot = nextPos
				}
				nextPos.x -= moveDir.x
				nextPos.y -= moveDir.y
				board[nextPos.y][nextPos.x] = '.'
				if done {
					break
				}
			}
		}
	}
	fmt.Println("Part 1: ", scoreBoard(board))
}

func printBoard(board [][]byte) {
	for _, line := range board {
		fmt.Println(string(line))
	}
}

func scoreBoard(board [][]byte) int {
	score := 0
	for j, line := range board {
		for i, char := range line {
			if char == 'O' {
				score += 100*j + i
			}
		}
	}
	return score
}
