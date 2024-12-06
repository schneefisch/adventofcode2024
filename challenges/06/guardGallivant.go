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

type Coordinate struct {
	x, y int
}

type Guard struct {
	pos    Coordinate
	facing Direction
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

type LabMap struct {
	width   int
	height  int
	lab     [][]rune
	visited [][]bool
	guard   *Guard
}

func (l *LabMap) isObstacle(x, y int) bool {
	return l.lab[y][x] == '#'
}

func (l *LabMap) isVisited(x, y int) bool {
	return l.visited[y][x]
}

func (l *LabMap) markVisited(x, y int) {
	l.visited[y][x] = true
}

func (l *LabMap) inMap(x, y int) bool {
	return x >= 0 && x < l.width && y >= 0 && y < l.height
}

// walkGuard walks the guard in the lab until she leaves the map
// will return false if the guard left the map, true otherwise
func (l *LabMap) walkGuard() bool {
	// determine next position
	nextPos := Coordinate{
		x: l.guard.pos.x,
		y: l.guard.pos.y,
	}
	switch l.guard.facing {
	case Up:
		nextPos.y--
	case Right:
		nextPos.x++
	case Down:
		nextPos.y++
	case Left:
		nextPos.x--
	}
	// check if the next position out of the map
	if !l.inMap(nextPos.x, nextPos.y) {
		return false
	}

	// check if the next position is an obstacle
	if l.isObstacle(nextPos.x, nextPos.y) {
		l.guard.turnRight()
		return true
	}

	// otherwise move the guard to the next position
	l.guard.pos = nextPos
	// mark the new position as visited
	l.markVisited(nextPos.x, nextPos.y)
	return true
}

func newLabMap(w, h int) *LabMap {
	labMap := &LabMap{
		width:   w,
		height:  h,
		lab:     make([][]rune, h),
		visited: make([][]bool, h),
		guard: &Guard{
			pos:    Coordinate{},
			facing: Up,
		},
	}
	for i := range labMap.visited {
		labMap.visited[i] = make([]bool, w)
	}
	return labMap
}

func GuardGallivant(filename string) (int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	labMap := parseMap(input)
	util.PrintMap(labMap.lab)

	// walk the guard and mark visited until she leaves the map
	for labMap.walkGuard() {
		// do nothing
		log.Printf("Guard position: %v", labMap.guard)
	}

	// count visited cells
	visitedCount := countVisited(labMap)

	return visitedCount, nil
}

func countVisited(labMap *LabMap) int {
	count := 0
	for y, row := range labMap.lab {
		for x := range row {
			if labMap.isVisited(x, y) {
				count++
			}
		}
	}

	return count
}

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
				break
			}
		}
	}
	log.Printf("Guard initial position: %v", labMap.guard)

	// mark first position as visited
	labMap.markVisited(labMap.guard.pos.x, labMap.guard.pos.y)

	return labMap
}

func isGuard(cell rune) bool {
	return cell == '^' || cell == 'v' || cell == '<' || cell == '>'
}
