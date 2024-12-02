package _1

import (
	"adventofcode2024/util"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func dayOneHistorianHysteria(filename string) (int, int, error) {
	// read input
	csv, err := util.ReadCSV(filename)
	if err != nil {
		return 0, 0, err
	}
	log.Println(csv)

	// get all numbers from the first column and convert to integers
	left, right, err := getLeftRightLists(csv)
	if err != nil {
		return 0, 0, err
	}

	// sort left slice
	err2 := sortSlices(left, right)
	if err2 != nil {
		return 0, 0, err2
	}

	// calculate distances
	distance := calculateDistance(left, right)

	// convert into weighted map
	leftMap := make(map[int]int)
	rightMap := make(map[int]int)
	convertToWeightedMap(left, leftMap)
	convertToWeightedMap(right, rightMap)

	similarityScore := calculateSimilarityScore(leftMap, rightMap)

	return distance, similarityScore, nil
}

func calculateSimilarityScore(leftMap map[int]int, rightMap map[int]int) int {
	score := 0
	for key, occurencesLeft := range leftMap {
		// check if the key is in the right map and how often
		if occurencesRight, ok := rightMap[key]; ok {
			score += key * occurencesRight * occurencesLeft
		}
	}
	return score
}

func convertToWeightedMap(left []int, leftMap map[int]int) {
	for _, nr := range left {
		// check if number is already in map
		if _, ok := leftMap[nr]; ok {
			leftMap[nr] += 1
		} else {
			leftMap[nr] = 1
		}
	}
}

func calculateDistance(left []int, right []int) int {
	distance := 0
	for i := 0; i < len(left); i++ {
		leftNr := left[i]
		rightNr := right[i]
		dist := rightNr - leftNr
		if dist < 0 {
			dist = -dist
		}
		distance += dist
	}
	return distance
}

func sortSlices(left []int, right []int) error {
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	// sort right slice
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	// check if left and right slices have the same length
	if len(left) != len(right) {
		return fmt.Errorf("left and right slices have different lengths")
	}
	return nil
}

func getLeftRightLists(csv [][]string) ([]int, []int, error) {
	var left []int
	var right []int
	for _, row := range csv {
		nr, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, nr)

		nr, err = strconv.Atoi(row[1])
		if err != nil {
			return nil, nil, err
		}
		right = append(right, nr)
	}
	return left, right, nil
}
