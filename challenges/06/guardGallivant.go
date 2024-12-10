package _6

import (
	"adventofcode2024/challenges/util"
	"log"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var (
	input = make([]string, 0)
)

type Coordinate struct {
	x, y int
}

type Guard struct {
	pos    Coordinate
	facing Direction
}

type Visited struct {
	visited   bool
	direction Direction
}

type LabMap struct {
	width   int
	height  int
	lab     [][]rune
	visited [][]Visited
	guard   *Guard
	start   Coordinate
}

func (g *Guard) directionFromRune(r rune) Direction {
	switch r {
	case '^':
		return Up
	case '>':
		return Right
	case 'v':
		return Down
	case '<':
		return Left
	default:
		panic("Invalid guard direction")
	}
}

func (g *Guard) turnRight() {
	switch g.facing {
	case Up:
		g.facing = Right
	case Right:
		g.facing = Down
	case Down:
		g.facing = Left
	case Left:
		g.facing = Up
	}
}

func (l *LabMap) isObstacle(x, y int) bool {
	if !l.inMap(x, y) {
		return false
	}
	return l.lab[y][x] == '#'
}

func (l *LabMap) isVisited(x, y int) (bool, Direction) {
	if !l.inMap(x, y) {
		return false, Up
	}

	return l.visited[y][x].visited, l.visited[y][x].direction
}

func (l *LabMap) markVisited(old, new Coordinate, d Direction) {
	if !l.inMap(new.x, new.y) {
		log.Fatalf("Invalid position: %v", new)
	}

	l.visited[new.y][new.x] = Visited{
		visited:   true,
		direction: d,
	}
	// if the old position is not the initial position, draw a line in the map
	if old.x == -1 || old.y == -1 {
		return
	}
	// draw a line from the old position to the new position
	line := '.'
	switch d {
	case Up, Down:
		line = '|'
	case Right, Left:
		line = '-'
	}
	l.lab[old.y][old.x] = line
	// place the arrow in the right position
	arrow := '^'
	switch d {
	case Up:
		arrow = '^'
	case Right:
		arrow = '>'
	case Down:
		arrow = 'v'
	case Left:
		arrow = '<'
	}
	l.lab[new.y][new.x] = arrow
}

// inMap checks if the given coordinates are within the map
func (l *LabMap) inMap(x, y int) bool {
	return x >= 0 && x < l.width && y >= 0 && y < l.height
}

// walkGuard walks the guard in the lab until she leaves the map
// will return false if the guard left the map, true otherwise
func (l *LabMap) walkGuard() bool {
	// determine next position
	next := nextPos(l.guard.pos, l.guard.facing)
	// check if the next position out of the map
	if !l.inMap(next.x, next.y) {
		return false
	}

	// check if the next position is an obstacle
	if l.isObstacle(next.x, next.y) {
		l.guard.turnRight()
		return true
	}

	// mark the new position as visited
	l.markVisited(l.guard.pos, next, l.guard.facing)

	// move the guard to the next position
	l.guard.pos = next
	return true
}

func newLabMap(w, h int) *LabMap {
	labMap := &LabMap{
		width:   w,
		height:  h,
		lab:     make([][]rune, h),
		visited: make([][]Visited, h),
		guard: &Guard{
			pos:    Coordinate{},
			facing: Up,
		},
	}
	for i := range labMap.visited {
		labMap.visited[i] = make([]Visited, w)
	}
	return labMap
}

// GuardGallivant calculates the Day 6 challenge of the adventOfCode2024
func GuardGallivant(filename string) (int, int, error) {
	newInput, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	input = newInput
	labMap := parseMap(input)
	util.PrintRuneMap(labMap.lab)

	// walk the guard and mark visited until she leaves the map
	for labMap.walkGuard() {
		//log.Printf("Guard position: %v", labMap.guard)
		//util.PrintRuneMap(labMap.lab)
	}

	// count visited cells
	visitedCount := countVisited(labMap)

	// for all visited cells, check if the position could be an obstacle that leads the
	// guard into an endless loop
	obstacleCount := findPossibleObstacles(labMap)

	return visitedCount, obstacleCount, nil
}

// findPossibleObstacles iterates over all visited cells and checks if the position could be an obstacle
// that leads the guard into an endless loop
// returns the number of possible obstacles
func findPossibleObstacles(labMap *LabMap) int {
	// iterate over all visited cells and check if that would be a possible obstacle
	count := 0
	for y, row := range labMap.lab {
		for x := range row {
			if visited, _ := labMap.isVisited(x, y); visited {
				if possible, _ := isPossibleObstacle(x, y); possible {
					count++
				}
			}
		}
	}
	return count
}

// isPossibleObstacle creates a new obstacle in the map
// then walks the guard from the initial position until he
// a) leaves the map, then it's not a possible obstacle
// b) walks into a place he has visited before in the same direction, then it's a possible obstacle
func isPossibleObstacle(x, y int) (bool, Coordinate) {
	// new obstacle-position
	pos := Coordinate{x, y}

	modifiedMap := parseMap(input)
	// add new obstacle
	modifiedMap.lab[y][x] = '#'

	// walk the guard until he leaves the map or walks into a visited cell
	for modifiedMap.walkGuard() {
		// get next position
		next := nextPos(modifiedMap.guard.pos, modifiedMap.guard.facing)
		if !modifiedMap.inMap(next.x, next.y) {
			continue
		}
		if modifiedMap.isObstacle(next.x, next.y) {
			continue
		}

		if visited, d := modifiedMap.isVisited(next.x, next.y); visited && d == modifiedMap.guard.facing {
			// found a place where the guard has visited before in the same direction
			//log.Printf("Found possible obstacle at %v", pos)
			//util.PrintRuneMap(modifiedMap.lab)
			return true, pos
		}
	}

	return false, pos
}

// nextPos calculates the next coordinates based on the current position and the direction
func nextPos(c Coordinate, d Direction) Coordinate {
	next := Coordinate{
		x: c.x,
		y: c.y,
	}
	switch d {
	case Up:
		next.y--
	case Right:
		next.x++
	case Down:
		next.y++
	case Left:
		next.x--
	}
	return next
}

// countVisited counts the number of visited cells in the labMap
func countVisited(labMap *LabMap) int {
	count := 0
	for y, row := range labMap.lab {
		for x := range row {
			if visited, _ := labMap.isVisited(x, y); visited {
				count++
			}
		}
	}

	return count
}

// parseMap parses the input into a LabMap
func parseMap(input []string) *LabMap {
	labMap := newLabMap(len(input[0]), len(input))
	for i, line := range input {
		lineRunes := []rune(line)
		labMap.lab[i] = lineRunes
	}
	// find guard position
	for y, row := range labMap.lab {
		for x, cell := range row {
			if isGuard(cell) {
				labMap.guard.pos = Coordinate{x, y}
				labMap.guard.facing = labMap.guard.directionFromRune(cell)
				labMap.start = Coordinate{x, y}
				break
			}
		}
	}
	//log.Printf("Guard initial position: %v", labMap.guard)

	// mark first position as visited
	labMap.markVisited(Coordinate{x: -1, y: -1}, labMap.guard.pos, labMap.guard.facing)

	return labMap
}

// isGuard checks if the cell is the current guards position
func isGuard(cell rune) bool {
	return cell == '^' || cell == 'v' || cell == '<' || cell == '>'
}
