package _15

import (
	"adventofcode2024/challenges/util"
	"log"
)

type Direction rune
type Element rune

const (
	Empty Element = '.'
	Wall  Element = '#'
	Box   Element = 'O'
	Robot Element = '@'

	Up    Direction = '^'
	Right Direction = '>'
	Down  Direction = 'v'
	Left  Direction = '<'
)

type Position struct {
	x, y int
}

type Warehouse struct {
	width, height int
	grid          [][]Element
	robotPosition *Position
}

func (w *Warehouse) moveRobot(direction Direction) {
	robot := w.robotPosition
	// check if there is free space between the robot, boxes and the wall
	// jump over boxes, since they can be moved
	if freeSpace := w.moveNext(robot, direction); freeSpace {
		// move robot should be moved
		//w.print()
		//log.Printf("Robot moved to %v\n", w.robotPosition)
	} else {
		//log.Printf("No free space to move robot to %v\n", direction.String())
	}
}

func (w *Warehouse) moveNext(pos *Position, direction Direction) bool {
	nextPos, nextElement := w.getNextGrid(pos, direction)
	isRobot := w.grid[pos.y][pos.x] == Robot
	// if it's a wall, return false
	switch nextElement {
	case Wall:
		return false
	case Empty:
		// move the current element to the next position
		w.switchGrids(pos, nextPos)
		if isRobot {
			// update robot position
			w.robotPosition.x = nextPos.x
			w.robotPosition.y = nextPos.y
		}
		// returning the current position wich is now empty
		return true
	case Box:
		// if the next element is a box, then call recursively
		if hasFreeSpace := w.moveNext(nextPos, direction); hasFreeSpace {
			// move the current element to the next position
			w.switchGrids(pos, nextPos)
			if isRobot {
				// update robot position
				w.robotPosition.x = nextPos.x
				w.robotPosition.y = nextPos.y
			}
			// returning the current position wich is now empty
			return true
		}
		// no free space, returning false
		return false
	default:
		log.Fatalf("Unknown element: %v", nextElement.String())
		return false
	}
}

func (w *Warehouse) switchGrids(pos1, pos2 *Position) {
	temp := w.grid[pos1.y][pos1.x]
	w.grid[pos1.y][pos1.x] = w.grid[pos2.y][pos2.x]
	w.grid[pos2.y][pos2.x] = temp
}

func (w *Warehouse) getNextGrid(pos *Position, direction Direction) (*Position, Element) {
	var next *Position
	switch direction {
	case Up:
		next = &Position{x: pos.x, y: pos.y - 1}
	case Right:
		next = &Position{x: pos.x + 1, y: pos.y}
	case Down:
		next = &Position{x: pos.x, y: pos.y + 1}
	case Left:
		next = &Position{x: pos.x - 1, y: pos.y}
	default:
		log.Fatalf("Unknown direction: %v", direction.String())
	}
	return next, w.grid[next.y][next.x]
}

func (w *Warehouse) print() {
	for _, row := range w.grid {
		for _, element := range row {
			print(string(element))
		}
		println()
	}
}

func (d *Direction) String() string {
	switch *d {
	case Up:
		return "Up"
	case Right:
		return "Right"
	case Down:
		return "Down"
	case Left:
		return "Left"
	default:
		return "Unknown"
	}
}

func (e *Element) String() string {
	switch *e {
	case Empty:
		return "Empty"
	case Wall:
		return "Wall"
	case Box:
		return "Box"
	case Robot:
		return "Robot"
	default:
		return "Unknown"
	}
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		width:  0,
		height: 0,
	}
}

func WarehouseWoes(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	warehouse := NewWarehouse()
	directions := make([]Direction, 0)
	warehouse, directions = parseWarehouse(input, warehouse)

	// move robot
	moveAll(warehouse, directions)

	// calculate GPS coordinate sum
	sum := gpsCoordinateSum(warehouse)

	return sum, 0, nil
}

func gpsCoordinateSum(warehouse *Warehouse) int {
	// iterate through the map and calculate the sum of the GPS coordinates of the boxes
	sum := 0
	for h, line := range warehouse.grid {
		for w, element := range line {
			if element == Box {
				// calculate the GPS Coordinate (Good Positioning System) by
				// multiplying x*100 + y
				sum += (h * 100) + w
			}
		}
	}
	return sum
}

func moveAll(warehouse *Warehouse, directions []Direction) {
	warehouse.print()
	for _, direction := range directions {
		warehouse.moveRobot(direction)
	}
}

func parseWarehouse(input []string, warehouse *Warehouse) (*Warehouse, []Direction) {
	height := 0
	for h, line := range input {
		// parse until we find an empty line
		if len(line) == 0 {
			log.Printf("Found empty line at %d\n", h)
			break
		}
		height++
	}

	// parse grid
	warehouse.height = height
	warehouse.width = len(input[0])
	warehouse.grid = make([][]Element, warehouse.height)
	for h := 0; h < warehouse.height; h++ {
		warehouse.grid[h] = make([]Element, warehouse.width)
		for w, char := range input[h] {
			warehouse.grid[h][w] = Element(char)
			// link the robot position
			if Element(char) == Robot {
				warehouse.robotPosition = &Position{
					x: w,
					y: h,
				}
			}
		}
	}

	// parse moves
	moves := make([]Direction, 0)
	for _, line := range input[warehouse.height:] {
		for _, char := range line {
			moves = append(moves, Direction(char))
		}
	}

	return warehouse, moves
}
