package _3

import (
	"adventofcode2024/challenges/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// MulNumbers is a struct that holds a list of the two numbers
type MulNumbers struct {
	n1, n2 int
}

// DaythreeMullitover reads a file, processes its content to find and multiply numbers, and returns the result.
// It takes a filename as input and returns the result of the added multiplications.
func DaythreeMullitover(filename string) (int, error) {
	lines, err := util.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	// add all lines into one string
	line := strings.Join(lines, "")
	// remove don't instructions
	line = removeDonts(line)
	log.Println(line)

	// get all mull-instances
	mulInstances := findAllMulInstances(line)
	mulMap, err := parseMulNumbers(mulInstances)
	if err != nil {
		return 0, err
	}
	log.Println(mulMap)

	// addMultiplications
	result := 0
	for _, mul := range mulMap {
		result += mul.n1 * mul.n2
	}

	return result, nil
}

// removeDonts parses through the string and removes all characters that are in between a
// "don't()" and a "do()" instruction.
// It takes a string as input and returns the modified string.
func removeDonts(line string) string {
	// split string at "don't()" and "do()"
	splitDonts := strings.Split(line, "don't()")
	// it starts with valid-expression, so the first part is always valid
	validInstructions := splitDonts[0]
	// iterate through the rest of the split
	for _, part := range splitDonts[1:] {
		// split at all "do()" instructions
		// add only the second and consecutive parts, after the "do()" instruction if it exists
		if splitDos := strings.Split(part, "do()"); len(splitDos) > 1 {
			validInstructions += strings.Join(splitDos[1:], "")
		}
	}
	return validInstructions
}

// parseMulNumbers parses the mul-instance strings and returns a list of MulNumbers.
// It takes a slice of strings as input and returns a slice of MulNumbers and an error if any.
func parseMulNumbers(mulInstances []string) ([]MulNumbers, error) {
	var mulMap []MulNumbers
	for _, mulInst := range mulInstances {
		// get the two numbers from the mul-instance
		n1, n2, err := parseMulNumbersFromInstance(mulInst)
		if err != nil {
			return nil, err
		}
		mulMap = append(mulMap, MulNumbers{n1, n2})
	}
	return mulMap, nil
}

// parseMulNumbersFromInstance parses the two numbers from a single mul-instance String.
// It takes a string as input and returns two integers and an error if any.
func parseMulNumbersFromInstance(mulInst string) (int, int, error) {
	// regex to match the numbers
	re := regexp.MustCompile("\\d{1,3}")
	// get all numbers
	nums := re.FindAllString(mulInst, -1)
	// expect exactly two numbers
	if len(nums) != 2 {
		return 0, 0, fmt.Errorf("expected 2 numbers, got %d", len(nums))
	}
	// convert the numbers to integers
	n1, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0, 0, err
	}
	n2, err := strconv.Atoi(nums[1])
	if err != nil {
		return 0, 0, err
	}
	return n1, n2, nil
}

// findAllMulInstances finds all mul-instances in a given string.
// It takes a string as input and returns a slice of strings containing all mul-instances.
func findAllMulInstances(line string) []string {
	// regular expression to match the mull-instances
	re := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	// get all mull-instances
	return re.FindAllString(line, -1)
}
