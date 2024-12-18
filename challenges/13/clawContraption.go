package _13

import (
	"adventofcode2024/challenges/util"
	"log"
	"regexp"
	"strconv"
)

type Coordinate struct {
	X, Y int
}

type Machine struct {
	ButtonA Coordinate
	ButtonB Coordinate
	Prize   Coordinate
}

// parseInput parses the input lines into Machine structs
func parseInput(lines []string) []Machine {
	var machines []Machine
	var currentMachine Machine
	lineCount := 0

	re := regexp.MustCompile(`[-\d]+`)

	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := re.FindAllString(line, -1)
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal("Failed to parse X coordinate:", err)
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal("Failed to parse Y coordinate:", err)
		}

		switch lineCount % 3 {
		case 0: // Button A
			currentMachine.ButtonA.X = x
			currentMachine.ButtonA.Y = y
		case 1: // Button B
			currentMachine.ButtonB.X = x
			currentMachine.ButtonB.Y = y
		case 2: // Prize
			currentMachine.Prize.X = x
			currentMachine.Prize.Y = y
			machines = append(machines, currentMachine)
			currentMachine = Machine{}
		}
		lineCount++
	}
	return machines
}

// solveMachine tries to find a solution for a single machine
func solveMachine(m Machine) (int, bool) {
	// Try all combinations of button presses up to 100
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			// Check if this combination reaches the prize
			if a*m.ButtonA.X+b*m.ButtonB.X == m.Prize.X &&
				a*m.ButtonA.Y+b*m.ButtonB.Y == m.Prize.Y {
				// Calculate token cost
				return 3*a + b, true
			}
		}
	}
	return 0, false
}

func ClawContraption(filename string) int {
	lines, err := util.ReadLines(filename)
	if err != nil {
		log.Fatal(err)
	}

	machines := parseInput(lines)
	totalTokens := 0

	for _, machine := range machines {
		if tokens, ok := solveMachine(machine); ok {
			totalTokens += tokens
		}
	}

	return totalTokens
}
