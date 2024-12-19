package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const ACost = 3
const BCost = 1
const TargetOffset = 10000000000000

type coords struct {
	x int
	y int
}

type clawMachine struct {
	aInc     coords
	bInc     coords
	prizeLoc coords
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	clawMachines := make([]clawMachine, 0, 0)
	stringBuffer := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			newClaw := parseClawMachine(stringBuffer)
			clawMachines = append(clawMachines, newClaw)
			stringBuffer = ""
			continue
		}
		stringBuffer += line + "\n"
	}
	totalTokens := 0
	for _, machine := range clawMachines {
		totalTokens += analyticSolveClawMachine(machine)
	}
	fmt.Println("Part 1: ", totalTokens)
}

func parseClawMachine(clawPrize string) (newClaw clawMachine) {
	aRegex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	bRegex := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	targetRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	match := aRegex.FindStringSubmatch(clawPrize)
	newClaw.aInc.x, _ = strconv.Atoi(match[1])
	newClaw.aInc.y, _ = strconv.Atoi(match[2])
	match = bRegex.FindStringSubmatch(clawPrize)
	newClaw.bInc.x, _ = strconv.Atoi(match[1])
	newClaw.bInc.y, _ = strconv.Atoi(match[2])
	match = targetRegex.FindStringSubmatch(clawPrize)
	newClaw.prizeLoc.x, _ = strconv.Atoi(match[1])
	newClaw.prizeLoc.x += TargetOffset
	newClaw.prizeLoc.y, _ = strconv.Atoi(match[2])
	newClaw.prizeLoc.y += TargetOffset

	return newClaw
}

func solveClawMachine(theClaw clawMachine) (tokens int) {
	// Strategy:
	//   - Press a until either x or y overshoots
	//   - Back off a, see if we can hit the target by pressing b
	//   - If keep backing off a and trying to hit with b
	//   - Do this all the way down to 0 a presses
	//   - If a combo hits, calculate its cost and record it as the best cost if it beats the previous best cost

	fmt.Println("Searching for ", theClaw.prizeLoc)

	xPos := 0
	yPos := 0

	bestResult := 0
	initialAPresses := 0

	for xPos < theClaw.prizeLoc.x && yPos < theClaw.prizeLoc.y {
		initialAPresses++
		xPos += theClaw.aInc.x
		yPos += theClaw.aInc.y
	}

	for trialAPresses := initialAPresses; trialAPresses >= 0; trialAPresses-- {
		trialBPresses := 0
		for {
			testLoc := coords{
				trialAPresses*theClaw.aInc.x + trialBPresses*theClaw.bInc.x,
				trialAPresses*theClaw.aInc.y + trialBPresses*theClaw.bInc.y,
			}
			if testLoc.x > theClaw.prizeLoc.x || testLoc.y > theClaw.prizeLoc.y {
				break
			}
			if testLoc.x == theClaw.prizeLoc.x &&
				testLoc.y == theClaw.prizeLoc.y {
				trialCost := trialAPresses*ACost + trialBPresses*BCost
				if bestResult == 0 || trialCost < bestResult {
					bestResult = trialCost
					break
				}
			}
			trialBPresses++
		}
	}

	fmt.Println("Found cost: ", bestResult)
	return bestResult
}

func analyticSolveClawMachine(theClaw clawMachine) (tokens int) {
	fmt.Println("Searching for ", theClaw.prizeLoc)
	var bAns int
	var aAns int
	bNum := theClaw.prizeLoc.y*theClaw.aInc.x - theClaw.aInc.y*theClaw.prizeLoc.x
	bDenom := theClaw.bInc.y*theClaw.aInc.x - theClaw.aInc.y*theClaw.bInc.x
	if bDenom != 0 {
		bAns = bNum / bDenom
		aAns = aFromB(theClaw, bAns)
		if (aAns*theClaw.aInc.x+bAns*theClaw.bInc.x != theClaw.prizeLoc.x) ||
			(aAns*theClaw.aInc.y+bAns*theClaw.bInc.y != theClaw.prizeLoc.y) {
			return 0
		}
		bestResult := bAns*BCost + aAns*ACost
		fmt.Println("Found cost: ", bestResult)
		return bestResult
	}
	abRatio := float64(theClaw.aInc.x) / float64(theClaw.bInc.x)
	preferA := abRatio > float64(ACost)/float64(BCost)
	if preferA {
		aAns = int(math.Floor(float64(theClaw.prizeLoc.x) / float64(theClaw.aInc.x)))
		for bAns = 0; aAns*theClaw.aInc.x+bAns*theClaw.bInc.x < theClaw.prizeLoc.x; bAns++ {
		}
	} else {
		bAns = int(math.Floor(float64(theClaw.prizeLoc.x) / float64(theClaw.bInc.x)))
		for aAns = 0; aAns*theClaw.aInc.x+bAns*theClaw.bInc.x < theClaw.prizeLoc.x; aAns++ {
		}
	}
	if (aAns*theClaw.aInc.x+bAns*theClaw.bInc.x != theClaw.prizeLoc.x) ||
		(aAns*theClaw.aInc.y+bAns*theClaw.bInc.y != theClaw.prizeLoc.y) {
		return 0
	}
	bestResult := bAns*BCost + aAns*ACost
	fmt.Println("Found cost: ", bestResult)
	return bestResult
}

func aFromB(theClaw clawMachine, bAns int) (aAns int) {
	return (theClaw.prizeLoc.x - bAns*theClaw.bInc.x) / theClaw.aInc.x
}
