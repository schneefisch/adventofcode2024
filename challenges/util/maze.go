package util

import "log"

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
	X, Y    int
	Kind    Kind
	Visited bool
}

type Position struct {
	X, Y int
	Dir  Direction
}

type Maze struct {
	Height, Width        int
	Grid                 [][]*Tile
	Position, Start, End Position
}

func (d *Direction) ToString() string {
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

func (m *Maze) Print() {
	for _, row := range m.Grid {
		line := ""
		for _, tile := range row {
			switch tile.Kind {
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
				log.Printf("Invalid tile Kind: %v", tile.Kind)
			}
		}
		log.Println(line)
	}
	log.Printf("Reindeer Position: {%d, %d}, %s", m.Position.X, m.Position.Y, m.Position.Dir.ToString())
}

func (m *Maze) IsInGrid(newPos Position) bool {
	return newPos.X >= 0 && newPos.X < m.Width && newPos.Y >= 0 && newPos.Y < m.Height
}

func (m *Maze) TileAt(pos Position) *Tile {
	return m.Grid[pos.Y][pos.X]
}

func (m *Maze) Parse(input []string) {
	m.Height = len(input)
	m.Width = len(input[0])
	m.Grid = make([][]*Tile, m.Height)
	for y, line := range input {
		row := make([]*Tile, m.Width)

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

			// set Position Position
			if kind == Start {
				m.Position = Position{x, y, East}
				m.Start = Position{x, y, East}
			}
			if kind == End {
				m.End = Position{x, y, East}
			}
		}

		// add row
		m.Grid[y] = row
	}
}
