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

// solveMachineCramersRule tries to find a solution for a single machine using Cramer's rule
// see cramers-rule: https://en.wikipedia.org/wiki/Cramer%27s_rule
func solveMachineCramersRule(m Machine) (int, bool) {
	dx1, dx2 := m.ButtonA.X, m.ButtonB.X
	dy1, dy2 := m.ButtonA.Y, m.ButtonB.Y
	prizeX, prizeY := m.Prize.X, m.Prize.Y

	// using Cramer's rule to solve the system of equations
	det := dx1*dy2 - dx2*dy1
	if det == 0 {
		return 0, false // no solution
	}

	detX := prizeX*dy2 - prizeY*dx2
	detY := dx1*prizeY - dy1*prizeX

	a := detX / det
	b := detY / det

	// check if we have integer soltions
	if a*dx1+b*dx2 != prizeX || a*dy1+b*dy2 != prizeY {
		return 0, false
	}

	// check if we have a positove solution
	if a < 0 || b < 0 {
		return 0, false
	}

	return 3*a + b, true
}

func ClawContraption(filename string) (int, int) {
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

	// solving part 2
	totalTokens2 := 0

	offset := 10000000000000
	for i := range machines {
		machines[i].Prize.X += offset
		machines[i].Prize.Y += offset
	}
	for _, machine := range machines {
		if tokens, ok := solveMachineCramersRule(machine); ok {
			totalTokens2 += tokens
		}
	}

	return totalTokens, totalTokens2
}
