package _10

import (
	"adventofcode2024/challenges/util"
	"log"
	"strconv"
)

type Coordinate struct {
	x int
	y int
}

type Trail struct {
	trail []Coordinate
}

type TrailHead struct {
	start  Coordinate
	trails []Trail
	score  int // counts the trails that reach a height-level of 9 outgoing from this trailhead
	rating int // rating counts all trails that start from this trailhead, no matter if they reach the same height
}

type TrailMap struct {
	mp         [][]int
	trailHeads []*TrailHead
}

// parse reads the string input of the topographic map as an array of strings
// and creates a 2D integer array with the height-profile that is stored in the map
func (m *TrailMap) parse(input []string) error {
	m.mp = make([][]int, len(input))

	// parse input string
	for i, line := range input {
		row := make([]int, len(line))
		// Parse the line
		for j, heightString := range line {
			// parse the height into an integer
			height, err := strconv.Atoi(string(heightString))
			if err != nil {
				return err
			}
			row[j] = height
		}
		m.mp[i] = row
	}

	// find all trailheads
	for y, row := range m.mp {
		for x, height := range row {
			if height == 0 {
				m.trailHeads = append(m.trailHeads, &TrailHead{
					Coordinate{x: x, y: y},
					nil,
					0,
					0})
			}
		}
	}

	log.Println("Trail map:")
	util.PrintIntMap(m.mp)

	return nil
}

// parseTrails actually walks all possible trails from each trailhead to identify the score and rating for each
func (m *TrailMap) parseTrails() {
	for _, trailHead := range m.trailHeads {
		// find trails
		peaks, uniqueTrails := m.walkTrails(trailHead.start)
		trailHead.score = len(peaks)
		trailHead.rating = uniqueTrails
	}
}

// walkTrails finds all trails starting from a given trailhead
// a trail always starts at height 0
// trails must continuously increase in height at each step (up, down, left, right) by exactly 1
func (m *TrailMap) walkTrails(pos Coordinate) ([]Coordinate, int) {
	// check if the trail has reached a height of 9 and increment the score and stop the trail
	// for each direction, check if the next step is valid
	// if it is, check if there are more than one possible next step, which means the trail is splitting.
	// if there is only one possibility, continue the trail in that direction
	// if the trail is splitting, create a second trail with the previous steps and continue all trails
	peaks := make([]Coordinate, 0)

	// if height is 9, return 1
	currentHeight := m.height(pos)
	if currentHeight == 9 {
		peaks = append(peaks, pos)
		return peaks, 1
	}

	uniqueTrails := 0
	// find next steps
	up := Coordinate{x: pos.x, y: pos.y - 1}
	down := Coordinate{x: pos.x, y: pos.y + 1}
	left := Coordinate{x: pos.x - 1, y: pos.y}
	right := Coordinate{x: pos.x + 1, y: pos.y}
	if m.isInMap(up) && m.height(up) == currentHeight+1 {
		// continue this trail
		newPeaks, rating := m.walkTrails(up)
		peaks = appendUnique(peaks, newPeaks)
		uniqueTrails += rating
	}
	if m.isInMap(down) && m.height(down) == currentHeight+1 {
		// continue this trail
		newPeaks, rating := m.walkTrails(down)
		peaks = appendUnique(peaks, newPeaks)
		uniqueTrails += rating
	}
	if m.isInMap(left) && m.height(left) == currentHeight+1 {
		// continue this trail
		newPeaks, rating := m.walkTrails(left)
		peaks = appendUnique(peaks, newPeaks)
		uniqueTrails += rating
	}
	if m.isInMap(right) && m.height(right) == currentHeight+1 {
		// continue this trail
		newPeaks, rating := m.walkTrails(right)
		peaks = appendUnique(peaks, newPeaks)
		uniqueTrails += rating
	}
	return peaks, uniqueTrails
}

// height returns the height of a given coordinate in the map
func (m *TrailMap) height(pos Coordinate) int {
	return m.mp[pos.y][pos.x]
}

// isInMap checks if a given coordinate is within the bounds of the map
func (m *TrailMap) isInMap(pos Coordinate) bool {
	if pos.y < 0 || pos.y >= len(m.mp) {
		return false
	}
	if pos.x < 0 || pos.x >= len(m.mp[0]) {
		return false
	}
	return true
}

func NewTrailMap() *TrailMap {
	return &TrailMap{
		mp:         make([][]int, 0),
		trailHeads: make([]*TrailHead, 0),
	}
}

func HoofIt(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	trailMap := NewTrailMap()
	if err = trailMap.parse(input); err != nil {
		return 0, 0, err
	}

	trailMap.parseTrails()

	sumScore := 0
	sumRating := 0
	for _, trailHead := range trailMap.trailHeads {
		sumScore += trailHead.score
		sumRating += trailHead.rating
	}

	return sumScore, sumRating, nil
}

// appendUnique returns only unique coordinates from two slices
func appendUnique(arr1, arr2 []Coordinate) []Coordinate {
	// add all coordinates from arr2 that are not already in unique
	for _, c := range arr2 {
		if !containsCoordinate(arr1, c) {
			arr1 = append(arr1, c)
		}
	}
	return arr1
}

// containsCoordinate checks if a coordinate is already in an array
func containsCoordinate(array []Coordinate, c Coordinate) bool {
	for _, coord := range array {
		if coord.x == c.x && coord.y == c.y {
			return true
		}
	}
	return false
}
