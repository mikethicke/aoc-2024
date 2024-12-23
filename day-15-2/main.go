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

type entity struct {
	kind byte
}

type board map[coords]entity

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	board := make(map[coords]entity)
	var moves []coords
	var width, height int
	lineNum := 0
	var robotPos coords
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		i := 0
		for _, char := range line {
			switch char {
			case '#', '.':
				board[coords{lineNum, i}] = entity{kind: char}
				i++
				board[coords{lineNum, i}] = entity{kind: char}
				width = i + 1
			case 'O':
				board[coords{lineNum, i}] = entity{kind: '['}
				i++
				board[coords{lineNum, i}] = entity{kind: ']'}
			case '@':
				board[coords{lineNum, i}] = entity{kind: '@'}
				robotPos = coords{lineNum, i}
				i++
				board[coords{lineNum, i}] = entity{kind: '.'}
			case '^':
				moves = append(moves, coords{-1, 0})
			case '>':
				moves = append(moves, coords{0, 1})
			case 'v':
				moves = append(moves, coords{1, 0})
			case '<':
				moves = append(moves, coords{0, -1})
			}
			i++
		}
		lineNum++
	}
	height = lineNum
	//printBoard(board, height, width)
	for i, move := range moves {
		moveable, moved := evaluateMove(board, move, robotPos)
		if moveable {
			robotPos = doMove(board, move, moved)
			fmt.Println("Move ", i)
			//printBoard(board, height, width)
		}
	}
	score := scoreBoard(board, height, width)
	//printBoard(board, height, width)
	fmt.Println("Part 2: ", score)
}

func printBoard(board board, height int, width int) {
	for j := 0; j < height-1; j++ {
		fmt.Print(j, " ")
		for i := 0; i < width; i++ {
			fmt.Print(string(board[coords{j, i}].kind))
		}
		fmt.Println()
	}
	fmt.Print("  ")
	for i := 0; i < width; i++ {
		fmt.Print(i % 10)
	}
	fmt.Println()
}

func evaluateMove(theBoard board, moveDir coords, startPos coords) (movable bool, moved []coords) {
	nextPos := coords{
		y: startPos.y + moveDir.y,
		x: startPos.x + moveDir.x,
	}
	moved = append(moved, startPos)
	if theBoard[nextPos].kind == '#' {
		return false, nil
	}
	if theBoard[nextPos].kind == '.' {
		return true, moved
	}
	if moveDir == (coords{0, 1}) || moveDir == (coords{0, -1}) {
		nextMovable, nextMoved := evaluateMove(theBoard, moveDir, nextPos)
		moved = append(moved, nextMoved...)
		return nextMovable, moved
	}
	switch theBoard[nextPos].kind {
	case '[':
		leftMovable, leftMoved := evaluateMove(theBoard, moveDir, nextPos)
		rightMovable, rightMoved := evaluateMove(theBoard, moveDir, coords{nextPos.y, nextPos.x + 1})
		if leftMovable && rightMovable {
			moved = append(moved, leftMoved...)
			moved = append(moved, rightMoved...)
			return true, moved
		}
		return false, nil
	case ']':
		rightMovable, rightMoved := evaluateMove(theBoard, moveDir, nextPos)
		leftMovable, leftMoved := evaluateMove(theBoard, moveDir, coords{nextPos.y, nextPos.x - 1})
		if rightMovable && leftMovable {
			moved = append(moved, rightMoved...)
			moved = append(moved, leftMoved...)
			return true, moved
		}
		return false, nil
	}
	//We should never get here
	return false, nil
}

func doMove(theBoard board, moveDir coords, toMove []coords) (robotPos coords) {
	moved := make(map[coords]bool)
	for i := len(toMove) - 1; i >= 0; i-- {
		if moved[toMove[i]] {
			continue
		}
		moved[toMove[i]] = true
		nextPos := coords{
			y: toMove[i].y + moveDir.y,
			x: toMove[i].x + moveDir.x,
		}
		theBoard[nextPos] = theBoard[toMove[i]]
		if theBoard[nextPos].kind == '@' {
			robotPos = nextPos
		}
		theBoard[toMove[i]] = entity{kind: '.'}
	}
	return robotPos
}

func scoreBoard(theBoard board, height int, width int) (score int) {
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if theBoard[coords{j, i}].kind == '[' {
				score += 100 * j
				score += i
			}
		}
	}
	return score
}
