package _4

import (
	"adventofcode2024/challenges/util"
	"log"
	"regexp"
	"strings"
)

func CeresSearch(filename string) (int, error) {
	lines, err := util.ReadLines(filename)
	if err != nil {
		return 0, err
	}

	runeMap := util.SplitLinesToCharacterMap(lines)
	util.PrintMap(runeMap)
	log.Println("Rotated map:")
	rotatedMap := util.RotateMatrix(runeMap)
	util.PrintMap(rotatedMap)

	//xmasOccurences := findOccurrencesInMap(runeMap)
	//xmasOccurences += findOccurrencesInMap(rotatedMap)
	xmasOccurences := findDiagonalOccurrencesInMap(rotatedMap)

	return xmasOccurences, nil
}

type Location struct {
	x int
	y int
}

func findDiagonalOccurrencesInMap(matrix [][]rune) int {
	reg := regexp.MustCompile("MAS")
	regReverse := regexp.MustCompile("SAM")

	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	// creating a slice of strings diagonals from right-top to left-bottom
	res := []string{}
	for d := 0; d < n+m-1; d++ {
		cur := ""
		for x := max(0, d-m+1); x < min(n, d+1); x++ {
			y := d - x
			cur += string(matrix[x][y])
		}
		res = append(res, cur)
	}

	// find all locations in this direction
	middleIndicesLeft := []Location{}
	for l, diag := range res {
		index := reg.FindAllStringIndex(diag, -1)
		index = append(index, regReverse.FindAllStringIndex(diag, -1)...)
		for _, i := range index {
			// find the location from the middle character
			stringIndex := i[0] + 1
			// since this is the center of the diagonal string from a x*x matrix, we now need
			// to find the location in the matrix
			// the first string is the bottom-left corner of the matrix with, therefore we can consider
			// the line-number
			// the last string is the top-right corner of the matrix.
			// e.g. the third string ("l") is "SAM", the middle character has index 1
			// the Location should be (in a 10x10 matrix) (8, 1)
			//
			// for a 10x10 matrix we have 19 diagonals (2*len(matrix)-1)
			// for the first 10
			// for the test-input the diagonal locations should be
			//lineSubtract := 0
			//if l > n {
			//    lineSubtract = l - n
			//}
			xcorrection := n - l - 1
			if xcorrection < 0 {
				xcorrection = 0
			}
			ycorrection := 0
			if l >= n {
				ycorrection = l - n + 1
			}
			location := Location{
				x: xcorrection + stringIndex,
				y: ycorrection + stringIndex,
			}
			middleIndicesLeft = append(middleIndicesLeft, location)
		}
		log.Println("Indices:", index)
	}
	log.Println("Middle indices left:", middleIndicesLeft)

	// create the slice of string in the opposite direction
	res = []string{}
	for d := 0; d < n+m-1; d++ {
		cur := ""
		for x := max(0, d-m+1); x < min(n, d+1); x++ {
			y := d - x
			cur += string(matrix[n-x-1][y])
		}
		res = append(res, cur)
	}
	// find locations in the other direction
	middleIndicesRight := []Location{}
	for l, diag := range res {
		index := reg.FindAllStringIndex(diag, -1)
		index = append(index, regReverse.FindAllStringIndex(diag, -1)...)
		for _, i := range index {
			stringIndex := i[0] + 1
			xcorrection := n - l - 1
			if l >= n {
				xcorrection = 0
			}
			ycorrection := n - 1
			if l >= n {
				ycorrection = ycorrection - (l - n + 1)
			}
			location := Location{
				x: xcorrection + stringIndex,
				y: ycorrection - stringIndex,
			}
			middleIndicesRight = append(middleIndicesRight, location)
		}
	}
	log.Println("Middle indices right:", middleIndicesRight)

	// count occurrences of the same locations in both slices
	occurrences := 0
	for _, left := range middleIndicesLeft {
		for _, right := range middleIndicesRight {
			if left.x == right.x && left.y == right.y {
				occurrences++
			}
		}
	}

	return occurrences
}

// reverseString reverses a string.
func reverseString(s string) string {
	// reversing the order of the string
	runes := []rune(s)
	reversed := make([]rune, len(runes))
	for i, j := range runes {
		reversed[len(runes)-1-i] = j
	}
	return string(reversed)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findOccurrencesInMap(runeMap [][]rune) int {
	xmasOccurences := 0
	for _, row := range runeMap {
		rowString := string(row)
		xmasOccurences += strings.Count(rowString, "XMAS")
		xmasOccurences += strings.Count(rowString, "SAMX")
	}
	return xmasOccurences
}
