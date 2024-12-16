package _16

import (
	"adventofcode2024/challenges/util"
	"log"
	"strconv"
)

type Kind int
type Direction int

const (
	// Kind types
	Wall Kind = iota
	Empty
	Start
	End

	// Direction types
	North Direction = iota
	East
	South
	West
)

type Tile struct {
	x, y    int
	kind    Kind
	visited bool
}

type Position struct {
	x, y int
	dir  Direction
}

type Maze struct {
	grid          [][]*Tile
	reindeer      Position
	height, width int
	end           Position
}

func (d *Direction) toString() string {
	switch *d {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	default:
		return "Unknown"
	}
}

func (m *Maze) print() {
	for _, row := range m.grid {
		line := ""
		for _, tile := range row {
			switch tile.kind {
			case Wall:
				line += "#"
			case Empty:
				line += "."
			case Start:
				line += "S"
			case End:
				line += "E"
			default:
				// nothing
				log.Printf("Invalid tile kind: %v", tile.kind)
			}
		}
		log.Println(line)
	}
	log.Printf("Reindeer position: {%d, %d}, %s", m.reindeer.x, m.reindeer.y, m.reindeer.dir.toString())
}

func ReindeerMaze(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}

	maze := Maze{}
	parseMaze(input, &maze)

	maze.print()
	lowestScore := walkMaze(&maze)

	return lowestScore, 0, nil
}

func walkMaze(m *Maze) int {
	return bfs(m)
}

type Node struct {
	pos   Position
	score int
}

type Visited struct {
	nodes []Node
}

func (v *Visited) add(pos Position, score int) {
	v.nodes = append(v.nodes, Node{pos, score})
}

func (v *Visited) find(find Position) (int, bool) {
	for _, visited := range v.nodes {
		if visited.pos.x == find.x && visited.pos.y == find.y {
			return visited.score, true
		}
	}
	return -1, false
}

func (v *Visited) getAll() []Node {
	return v.nodes
}

// bfs checks all possible next steps
// it uses a bfs algorithm (breath-first search) to find the shortest path
func bfs(m *Maze) int {

	// create queue of notes to visit
	queue := []Node{{m.reindeer, 0}}
	visited := Visited{nodes: []Node{}}
	visited.add(m.reindeer, 0)

	directions := []Direction{North, East, South, West}

	cheapestScore := -1

	// keep on walking until the queue is empty
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // remove first node from queue

		// check if we reached the end node
		if m.grid[current.pos.y][current.pos.x].kind == End {
			log.Printf("End found at {%d, %d}", current.pos.x, current.pos.y)
			// update score
			if cheapestScore == -1 {
				cheapestScore = current.score
			} else if current.score < cheapestScore {
				cheapestScore = current.score
			}

			continue
		}

		// check possible directions and add to queue
		for _, dir := range directions {
			newPos := Position{
				x:   current.pos.x,
				y:   current.pos.y,
				dir: dir,
			}
			switch dir {
			case North:
				newPos.y--
			case East:
				newPos.x++
			case South:
				newPos.y++
			case West:
				newPos.x--
			}

			// check if the new position is within the maze
			if newPos.x >= 0 && newPos.x < m.width && newPos.y >= 0 && newPos.y < m.height {
				tile := m.grid[newPos.y][newPos.x]
				newScore := current.score + rotationScore(current.pos.dir, dir) + 1
				// check if it's a wall or already visited
				if tile.kind != Wall {
					// check if the new score is lower than the old score
					// this accounts for the fact that we can visit the same node multiple times, but the later visit might have a lower score
					if oldScore, found := visited.find(newPos); !found || newScore < oldScore {
						// add to visited
						visited.add(newPos, newScore)
						// append to queue
						queue = append(queue, Node{newPos, newScore})
					}
				}
			}
		}
	}

	log.Println("Visited nodes with score:")
	// print visited nodes in a map according to their position
	pathMap := make([][]int, m.height)
	for y, row := range pathMap {
		row = make([]int, m.width)
		pathMap[y] = row
	}
	for _, node := range visited.getAll() {
		pathMap[node.pos.y][node.pos.x] = node.score
	}
	for _, row := range pathMap {
		line := ""
		for _, score := range row {
			line += " " + strconv.Itoa(score)
		}
		log.Println(line)
	}

	return cheapestScore // end not found
}

func rotationScore(dir Direction, newDirection Direction) int {
	// can only rotate 90 degrees in both directions
	// each rotation costs 1000 points
	// if the new direction is the same as the current direction, the cost is 0
	clockwiseCost := 0
	counterClockwiseCost := 0

	temp := dir
	for i := 0; i < 4; i++ {
		if temp == newDirection {
			clockwiseCost = i * 1000
			break
		}
		temp = rotate90DegreesClockwise(temp)
	}

	temp = dir
	for i := 0; i < 4; i++ {
		if temp == newDirection {
			counterClockwiseCost = i * 1000
			break
		}
		temp = rotate90DegreesCounterClockwise(temp)
	}

	if clockwiseCost < counterClockwiseCost {
		return clockwiseCost
	}
	return counterClockwiseCost
}

func rotate90DegreesCounterClockwise(dir Direction) Direction {
	switch dir {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	case West:
		return South
	default:
		return -1
	}
}

func rotate90DegreesClockwise(dir Direction) Direction {
	switch dir {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		return -1
	}
}

func parseMaze(input []string, m *Maze) {
	m.height = len(input)
	m.width = len(input[0])
	m.grid = make([][]*Tile, m.height)
	for y, line := range input {
		row := make([]*Tile, m.width)

		for x, char := range line {
			var kind Kind
			switch char {
			case '#':
				kind = Wall
			case '.':
				kind = Empty
			case 'S':
				kind = Start
			case 'E':
				kind = End
			default:
				panic("invalid character in maze")
			}
			row[x] = &Tile{x, y, kind, false}

			// set reindeer position
			if kind == Start {
				m.reindeer = Position{x, y, East}
			}
			if kind == End {
				m.end = Position{x, y, East}
			}
		}

		// add row
		m.grid[y] = row
	}
}
