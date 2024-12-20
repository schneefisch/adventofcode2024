package _16

import (
	"adventofcode2024/challenges/util"
	"container/heap"
	"log"
	"strconv"
)

type Node struct {
	pos   util.Position
	score int
}

type Visited struct {
	nodes []Node
}

func (v *Visited) add(pos util.Position, score int) {
	v.nodes = append(v.nodes, Node{pos, score})
}

func (v *Visited) find(find util.Position) (int, bool) {
	for _, visited := range v.nodes {
		if visited.pos.X == find.X && visited.pos.Y == find.Y {
			return visited.score, true
		}
	}
	return -1, false
}

func (v *Visited) getAll() []Node {
	return v.nodes
}

// PriorityQueue is a Priority-queue sorted by the score of the nodes. The lowest score is at the top of the queue.
type PriorityQueue []Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].score < pq[j].score }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func ReindeerMaze(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}

	maze := util.Maze{}
	maze.ParseMaze(input)

	maze.Print()
	lowestScore := walkMaze(&maze)

	return lowestScore, 0, nil
}

func walkMaze(m *util.Maze) int {
	return dijkstra(m)
}

// dijkstra calculates the cheapest path in a dijkstra approach with weighted edges
func dijkstra(m *util.Maze) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{m.Position, 0})

	visited := Visited{nodes: []Node{}}
	directions := []util.Direction{util.North, util.East, util.South, util.West}

	// ToDo: for part two, I need all of the cheapest paths.
	// update the loop to get continue until all paths are found that have the same score, as soon as the score
	// starts rising over the lowest score, we can stop the loop
	// also return the Visited nodes, so we can see the path and get the best places to sit on

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Node)

		// check if we reached the End node
		if m.TileAt(current.pos).Kind == util.End {
			log.Printf("End found at {%d, %d} with score %d", current.pos.X, current.pos.Y, current.score)
			return current.score
		}

		// skip if already Visited with a cheaper score
		if oldScore, found := visited.find(current.pos); found && current.score >= oldScore {
			continue
		}
		visited.add(current.pos, current.score)

		// explore neighbours
		for _, dir := range directions {
			newPos := nextPos(current, dir)

			// skip walls and off-Grid tiles
			if !m.IsInGrid(newPos) || m.TileAt(newPos).Kind == util.Wall {
				continue
			}

			// calculate cost
			newScore := current.score + rotationScore(current.pos.Dir, dir) + 1
			if oldScore, found := visited.find(newPos); !found || newScore < oldScore {
				heap.Push(pq, Node{newPos, newScore})
			}
		}
	}

	return -1
}

// dfs calculates the shortest path in a depth-first search approach
func dfs(m *util.Maze) int {
	stack := []Node{{m.Position, 0}}
	visited := Visited{nodes: []Node{}}
	visited.add(m.Position, 0)

	directions := []util.Direction{util.North, util.East, util.South, util.West}
	cheapestScore := -1

	for len(stack) > 0 {
		// put the "last" entry from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if m.Grid[current.pos.Y][current.pos.X].Kind == util.End {
			if cheapestScore == -1 || current.score < cheapestScore {
				cheapestScore = current.score
			}
			log.Printf("End found at {%d, %d} with score %d", current.pos.X, current.pos.Y, current.score)
			continue
		}

		for _, dir := range directions {
			newPos := nextPos(current, dir)

			if m.IsInGrid(newPos) {
				tile := m.Grid[newPos.Y][newPos.X]
				newScore := current.score + rotationScore(current.pos.Dir, dir) + 1
				if tile.Kind != util.Wall {
					if oldScore, found := visited.find(newPos); !found || newScore < oldScore {
						visited.add(newPos, newScore)
						stack = append(stack, Node{newPos, newScore})
					}
				}
			}
		}
	}

	return cheapestScore
}

// bfs calculates the shortest path in a breadth-first search approach
// it's a modified bfs, because the rotations are so extremely expensive.
// NOTE: bfs is not the wanted result, it takes too long to find the shortest path
func bfs(m *util.Maze) int {

	// create queue of notes to visit
	queue := []Node{{m.Position, 0}}
	visited := Visited{nodes: []Node{}}
	visited.add(m.Position, 0)

	directions := []util.Direction{util.North, util.East, util.South, util.West}

	cheapestScore := -1

	// keep on walking until the queue is empty
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // remove first node from queue

		// check if we reached the End node
		if m.Grid[current.pos.Y][current.pos.X].Kind == util.End {
			log.Printf("End found at {%d, %d}", current.pos.X, current.pos.Y)
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
			newPos := nextPos(current, dir)

			// check if the new Position is within the maze
			if m.IsInGrid(newPos) {
				tile := m.TileAt(newPos)
				newScore := current.score + rotationScore(current.pos.Dir, dir) + 1
				// check if it's a wall or already Visited
				if tile.Kind != util.Wall {
					// check if the new score is lower than the old score
					// this accounts for the fact that we can visit the same node multiple times, but the later visit might have a lower score
					if oldScore, found := visited.find(newPos); !found || newScore < oldScore {
						// add to Visited
						visited.add(newPos, newScore)
						// append to queue
						queue = append(queue, Node{newPos, newScore})
					}
				}
			}
		}
	}

	log.Println("Visited nodes with score:")
	// Print Visited nodes in a map according to their Position
	pathMap := make([][]int, m.Height)
	for y, row := range pathMap {
		row = make([]int, m.Width)
		pathMap[y] = row
	}
	for _, node := range visited.getAll() {
		pathMap[node.pos.Y][node.pos.X] = node.score
	}
	for _, row := range pathMap {
		line := ""
		for _, score := range row {
			line += " " + strconv.Itoa(score)
		}
		log.Println(line)
	}

	return cheapestScore // End not found
}

func nextPos(current Node, dir util.Direction) util.Position {
	newPos := util.Position{
		X:   current.pos.X,
		Y:   current.pos.Y,
		Dir: dir,
	}
	switch dir {
	case util.North:
		newPos.Y--
	case util.East:
		newPos.X++
	case util.South:
		newPos.Y++
	case util.West:
		newPos.X--
	}
	return newPos
}

func rotationScore(dir util.Direction, newDirection util.Direction) int {
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

func rotate90DegreesCounterClockwise(dir util.Direction) util.Direction {
	switch dir {
	case util.North:
		return util.West
	case util.East:
		return util.North
	case util.South:
		return util.East
	case util.West:
		return util.South
	default:
		return -1
	}
}

func rotate90DegreesClockwise(dir util.Direction) util.Direction {
	switch dir {
	case util.North:
		return util.East
	case util.East:
		return util.South
	case util.South:
		return util.West
	case util.West:
		return util.North
	default:
		return -1
	}
}
