package _19

import (
	"adventofcode2024/challenges/util"
	"log"
	"strings"
)

func LinenLayout(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	towels, designs := parseInput(input)
	log.Printf("towels: %v\ndesigns: %v\n", towels, designs)

	possibleDesigns, allPatterns := findPossibleDesigns(towels, designs)

	return possibleDesigns, allPatterns, nil
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

func findPossibleDesigns(towels []string, designs []string) (int, int) {
	// The design is a combination of towels, defined by the collor-pattern of the towels.
	// to find possible designs, we can iterate over all towels and check if the design can be made with a combination of towels.
	// we can use a backtracking approach to find all possible designs.

	possibleDesigns := make([]string, 0)
	allPatterns := 0
	for _, design := range designs {
		if isPossible, patterns := isPossibleDesign(towels, design); isPossible {
			possibleDesigns = append(possibleDesigns, design)
			allPatterns += patterns
		}
	}

	return len(possibleDesigns), allPatterns
}

// isPossibleDesign uses dynamic programming to solve the problem of finding:
// 1. Whether it's possible to create the target design using the given towels
// 2. The total number of different ways to create the design
func isPossibleDesign(towels []string, design string) (bool, int) {

	// The dp array stores the number of ways to create each prefix of the target design:
	// - dp[i] represents the number of ways to create the first i characters of the design
	// - dp[i] = 0 means it's impossible to create that prefix
	// - dp[i] > 0 gives the number of different valid combinations for that prefix
	dp := make([]int, len(design)+1)

	// Empty string has exactly one way to make it
	dp[0] = 1

	// Try to build the design string incrementally from left to right
	// The algorithm works by:
	// 1. Starting with an empty string (dp[0] = 1, as there's one way to make empty string)
	// 2. For each position i in the design:
	//    - Try each towel as a possible ending piece
	//    - If a towel can be placed at position i:
	//      * Add the number of ways to build the prefix before this towel
	// 3. The final value dp[len(design)] gives both:
	//    - Whether it's possible (> 0)
	//    - The total number of different valid combinations
	for i := 1; i <= len(design); i++ {
		// For each position, try every towel as a possible end piece
		for _, towel := range towels {
			if i >= len(towel) && // 1. Current position is long enough to fit this towel
				dp[i-len(towel)] > 0 && // 2. We could build the string up to the start of where this towel would go
				design[i-len(towel):i] == towel { // 3. The towel exactly matches the substring ending at current position

				// Add the number of ways to build the prefix to current position
				dp[i] += dp[i-len(towel)]
			}
		}
	}

	// Return whether it's possible (dp[len(design)] > 0) and the number of ways
	return dp[len(design)] > 0, dp[len(design)]
}
