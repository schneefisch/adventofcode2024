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
