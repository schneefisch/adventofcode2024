package _7

import (
	"adventofcode2024/challenges/util"
	"log"
	"strconv"
	"strings"
)

type Equation struct {
	solution int
	numbers  []int
}

func BridgeRepair(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	// parse equations
	equations := make([]Equation, len(input))
	if err = parseEquations(input, equations); err != nil {
		return 0, 0, err
	}
	log.Printf("Equations: %v", equations)

	// solve equations
	operators := []rune{'+', '*'}
	sum := sumValidEquations(equations, operators)

	operators = []rune{'*', '+', '|'}
	sumExtended := sumValidEquations(equations, operators)

	return sum, sumExtended, nil
}

// sumValidEquations sums the solutions of all valid equations
func sumValidEquations(equations []Equation, operators []rune) int {
	sum := 0
	ch := make(chan int, len(equations))
	defer close(ch)

	for _, equation := range equations {
		go func(eq Equation) {
			if isValidEquation(eq, operators) {
				ch <- equation.solution
			} else {
				ch <- 0
			}
		}(equation)
	}
	for range equations {
		sum += <-ch
	}

	return sum
}

// isValidEquation checks if the equation is valid
// we have two operators, + and *
// check if there is any combination of the numbers that results in the solution
func isValidEquation(equation Equation, operators []rune) bool {
	combinations := generateCombinations(len(equation.numbers)-1, operators)
	//log.Printf("Combinations: %v", combinations)

	// check all combinations
	for _, combination := range combinations {
		// convert the combination to a slice of runes
		operatorRunes := []rune(combination)
		// check if the equation is valid
		if isValidCombination(equation, operatorRunes) {
			//log.Printf("Valid combination: numbers: %v, operators: %v", equation.numbers, combination)
			return true
		}
	}

	return false
}

func isValidCombination(equation Equation, operators []rune) bool {
	// calculate the result of the equation
	result := 0
	for i, number := range equation.numbers {
		if i == 0 {
			result = number
			continue
		}
		operator := operators[i-1]
		switch operator {
		case '+':
			result += number
		case '*':
			result *= number
		case '|':
			// convert number to string
			numberStr := strconv.Itoa(result) + strconv.Itoa(number)
			// convert back to int
			result, _ = strconv.Atoi(numberStr)
		}
	}
	// check if the result is the solution
	if result == equation.solution {
		return true
	}
	return false
}

// generateCombinations generates all possible combinations of operators for a given length
func generateCombinations(length int, operators []rune) []string {
	if length == 0 {
		return []string{""}
	}
	var combinations []string
	smallerCombinations := generateCombinations(length-1, operators)
	for _, operator := range operators {
		for _, combination := range smallerCombinations {
			combinations = append(combinations, string(operator)+combination)
		}
	}
	return combinations
}

func parseEquations(input []string, equations []Equation) error {
	for i, line := range input {
		//log.Printf("Parsing line: %v", line)
		// split after the colon
		parts := strings.Split(line, ":")
		// trim the whitespace
		secondPart := strings.TrimSpace(parts[1])
		// split the second part by whitespace
		secondParts := strings.Split(secondPart, " ")
		// parse int of the fist part
		key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return err
		}
		// parse the ints of the second part
		numbers := make([]int, 0)
		for _, part := range secondParts {
			number, err := strconv.Atoi(part)
			if err != nil {
				return err
			}
			numbers = append(numbers, number)
		}
		equations[i] = Equation{
			solution: key,
			numbers:  numbers,
		}
	}
	return nil
}
