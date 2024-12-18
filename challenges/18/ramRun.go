package _18

import (
	"adventofcode2024/challenges/util"
	"container/list"
	"log"
	"strconv"
	"strings"
)

type Cell rune
type Direction int

const (
	Empty     Cell = '.'
	Corrupted Cell = '#'

	North Direction = iota
	East
	South
	West
)

type Position struct {
	x, y int
}

func (p *Position) Equals(other Position) bool {
	return p.x == other.x && p.y == other.y
}

type MemoryGrid struct {
	width, height int
	grid          [][]Cell
}

func (g *MemoryGrid) print() {
	for y := 0; y < g.height; y++ {
		s := ""
		for x := 0; x < g.width; x++ {
			s += string(g.grid[y][x])
		}
		log.Println(s)
	}
}
func (g *MemoryGrid) isInGrid(p Position) bool {
	return p.x >= 0 && p.x < g.width && p.y >= 0 && p.y < g.height
}
func (g *MemoryGrid) isCorrupted(p Position) bool {
	return g.grid[p.y][p.x] == Corrupted
}

type Node struct {
	pos   Position
	score int
}

type Visited struct{ nodes []Node }

func (v *Visited) add(n Node) { v.nodes = append(v.nodes, n) }
func (v *Visited) exists(p Position) bool {
	for _, node := range v.nodes {
		if node.pos.x == p.x && node.pos.y == p.y {
			return true
		}
	}
	return false
}

func RamRun(filename string, width, height, numBytes int) (int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	byteList := make([][]int, len(input))
	byteList = parseInput(input, byteList)

	// create a memory of cells
	memory := initMemory(width, height)
	// mark corrupted cells
	for i := 0; i < numBytes; i++ {
		bt := byteList[i]
		x, y := bt[0], bt[1]
		memory.grid[y][x] = Corrupted
	}
	memory.print()

	// find shortest path
	start := Position{0, 0}
	end := Position{width - 1, height - 1}
	steps := shortestPath(memory, start, end)

	return steps, nil
}

// find the shortest path using a breadth-first search
func shortestPath(memory *MemoryGrid, start Position, end Position) int {

	visited := Visited{nodes: []Node{}}
	visited.add(Node{start, 0})
	directions := []Direction{North, East, South, West}

	queue := list.New()
	queue.PushBack(Node{start, 0})

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Node)
		if current.pos.Equals(end) {
			return current.score
		}
		// find next directions
		for _, dir := range directions {
			newPos := Position{current.pos.x, current.pos.y}
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
			if memory.isInGrid(newPos) && !visited.exists(newPos) && !memory.isCorrupted(newPos) {
				newNode := Node{newPos, current.score + 1}
				visited.add(newNode)
				queue.PushBack(newNode)
			}
		}
	}
	return -1
}

func initMemory(width int, height int) *MemoryGrid {
	mem := &MemoryGrid{
		width:  width,
		height: height,
		grid:   make([][]Cell, height),
	}
	for y, row := range mem.grid {
		row = make([]Cell, width)
		for x := range row {
			row[x] = Empty
		}
		mem.grid[y] = row
	}
	return mem
}

func parseInput(input []string, list [][]int) [][]int {
	for i, line := range input {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			log.Fatalf("invalid input: %s", line)
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("invalid input: %s", line)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("invalid input: %s", line)
		}
		list[i] = []int{x, y}
	}
	return list
}
