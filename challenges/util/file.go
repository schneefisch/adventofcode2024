package util

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadCSV reads a CSV file and returns a slice of slices of strings.
func ReadCSV(filename string) ([][]string, error) {
	// read file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// close file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ReadSpaceSeparatedData reads a space-separated file and returns a slice of slices of strings.
func ReadSpaceSeparatedData(filename string) ([][]int, error) {
	// read file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// close file after function ends
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// read data line-by-line
	data := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read line as string and split by whitespace
		lineStr := scanner.Text()
		line := strings.Split(lineStr, " ")

		// convert strings to integers
		numbers := make([]int, 0)
		for _, numStr := range line {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
		data = append(data, numbers)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// ReadLines reads a file line by line and returns a slice of strings.
func ReadLines(filename string) ([]string, error) {
	// read file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// close file after function ends
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// read data line-by-line
	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// SplitLinesToCharacterMap splits the lines into a map of characters.
func SplitLinesToCharacterMap(lines []string) [][]rune {
	runeMap := make([][]rune, 0)
	for _, line := range lines {
		runeMap = append(runeMap, []rune(line))
	}
	return runeMap
}

// RotateMatrix rotates a matrix by 90 degrees.
func RotateMatrix[T any](matrix [][]T) [][]T {
	n := len(matrix)
	if n == 0 {
		return matrix
	}
	m := len(matrix[0])
	rotated := make([][]T, m)
	for i := range rotated {
		rotated[i] = make([]T, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rotated[j][n-i-1] = matrix[i][j]
		}
	}
	return rotated
}

// PrintIntMap prints a map of integers.
func PrintIntMap(matrix [][]int) {
	for _, row := range matrix {
		log.Println(row)
	}
}

// PrintRuneMap prints a map of characters.
func PrintRuneMap(matrix [][]rune) {
	for _, row := range matrix {
		log.Println(string(row))
	}
}
