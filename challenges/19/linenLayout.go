package _19

import (
	"adventofcode2024/challenges/util"
	"log"
	"strings"
)

func LinenLayout(filename string) (int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	towels, designs := parseInput(input)
	log.Printf("towels: %v\ndesigns: %v\n", towels, designs)

	possibleDesigns := findPossibleDesigns(towels, designs)

	return possibleDesigns, nil
}

// parseInput parses the input into towels and designs
// the first line is the available towels with their colors
// the second line is just an empty line as separator
// the following lines are the designs
func parseInput(input []string) ([]string, []string) {
	towels := make([]string, 0)
	towelsLine := input[0]
	towelsStrings := strings.Split(towelsLine, ", ")
	towels = append(towels, towelsStrings...)

	// parse designs
	designs := make([]string, 0)

	// parse designs
	designs = append(designs, input[2:]...)

	return towels, designs
}

func findPossibleDesigns(towels []string, designs []string) int {
	// The design is a combination of towels, defined by the collor-pattern of the towels.
	// to find possible designs, we can iterate over all towels and check if the design can be made with a combination of towels.
	// we can use a backtracking approach to find all possible designs.

	possibleDesigns := make([]string, 0)
	for _, design := range designs {
		if isPossibleDesign(towels, design) {
			possibleDesigns = append(possibleDesigns, design)
		}
	}

	return len(possibleDesigns)
}

func isPossibleDesign(towels []string, design string) bool {
	// dp[i] means "can we create the substring design[0:i] using our towels?"
	// For example, if design is "rgbw":
	// dp[0] = true  (empty string is always possible)
	// dp[1] = can we make "r"?
	// dp[2] = can we make "rg"?
	// dp[3] = can we make "rgb"?
	// dp[4] = can we make "rgbw"?
	dp := make([]bool, len(design)+1)

	// Empty string (length 0) is always possible to create
	dp[0] = true

	// Try to build the design string incrementally from left to right
	for i := 1; i <= len(design); i++ {
		// For each position, try every towel as a possible end piece
		for _, towel := range towels {
			// Three conditions must be met:
			if i >= len(towel) && // 1. Current position is long enough to fit this towel
				dp[i-len(towel)] && // 2. We could build the string up to the start of where this towel would go
				design[i-len(towel):i] == towel { // 3. The towel exactly matches the substring ending at current position

				// If all conditions are met, we can build the string up to position i
				dp[i] = true
				break // Found a valid towel, no need to try others
			}
		}
	}

	// Can we build the entire design string?
	return dp[len(design)]
}
