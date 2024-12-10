package _8

import (
	"adventofcode2024/challenges/util"
	"log"
)

type Antenna struct {
	frequency rune
	pos       Coordinate
}

type Coordinate struct {
	x int
	y int
}

type Map struct {
	width     int
	height    int
	antennas  map[rune][]Antenna
	antinodes []Coordinate
}

func (m *Map) addAntenna(cell rune, x, y int) {
	m.antennas[cell] = append(m.antennas[cell], Antenna{
		frequency: cell,
		pos:       Coordinate{x, y},
	})
}

func (m *Map) addAntinode(pos Coordinate) {
	// check if antinode is out of bounds
	if pos.x < 0 || pos.x >= m.width || pos.y < 0 || pos.y >= m.height {
		return
	}
	// check if antinode already exists
	for _, antinode := range m.antinodes {
		if antinode.x == pos.x && antinode.y == pos.y {
			return
		}
	}
	m.antinodes = append(m.antinodes, pos)
}

func NewMap(x, y int) *Map {
	return &Map{
		width:     x,
		height:    y,
		antennas:  make(map[rune][]Antenna),
		antinodes: make([]Coordinate, 0),
	}
}

func ResonantCollinearity(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}

	antennasMap := NewMap(len(input[0]), len(input))
	parseAntennas(input, antennasMap)

	calculateAntinodes(antennasMap, false)
	log.Println(antennasMap)

	sum := len(antennasMap.antinodes)
	calculateAntinodes(antennasMap, true)
	sum2 := len(antennasMap.antinodes)

	printMap(antennasMap, true)

	return sum, sum2, nil
}

func calculateAntinodes(antennasMap *Map, considerResonance bool) {
	for _, antennas := range antennasMap.antennas {
		// to have an antinode, we need at least 2 antennas
		if len(antennas) < 2 {
			continue
		}
		// iterate over all possible pairs of antennas
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				calculateAntinode(antennas[i], antennas[j], antennasMap, considerResonance)
			}
		}
	}
}

func calculateAntinode(antenna, antenna2 Antenna, antennasMap *Map, considerResonance bool) {
	// for each pair of antennas there can be two antinodes, one to each side of the line
	// the antinodes is calculated by finding the distance between the two antennas and
	// then adding the distance to the position of each antenna
	xDist, yDist := antenna2.pos.x-antenna.pos.x, antenna2.pos.y-antenna.pos.y
	if !considerResonance {
		// just calculate to each side and add to the map
		antinode1 := Coordinate{antenna.pos.x - xDist, antenna.pos.y - yDist}
		antinode2 := Coordinate{antenna2.pos.x + xDist, antenna2.pos.y + yDist}
		antennasMap.addAntinode(antinode1)
		antennasMap.addAntinode(antinode2)
		return
	}

	// calculating the antinodes with resonancy means that we need to keep adding the distance
	// and add antinodes at each step until we leave the map
	// this also includes the position of the original antennas, therefore starting at 0
	for i := 0; ; i++ {
		antinode1 := Coordinate{antenna.pos.x - xDist*i, antenna.pos.y - yDist*i}
		if antinode1.x < 0 || antinode1.x >= antennasMap.width || antinode1.y < 0 || antinode1.y >= antennasMap.height {
			break
		}
		antennasMap.addAntinode(antinode1)
	}
	for i := 0; ; i++ {
		antinode2 := Coordinate{antenna2.pos.x + xDist*i, antenna2.pos.y + yDist*i}
		if antinode2.x < 0 || antinode2.x >= antennasMap.width || antinode2.y < 0 || antinode2.y >= antennasMap.height {
			break
		}
		antennasMap.addAntinode(antinode2)
	}
}

// parseAntennas parses the input and populates the antennas map
func parseAntennas(input []string, antennasMap *Map) {
	for y, row := range input {
		for x, cell := range row {
			if cell == '.' {
				continue
			}
			antennasMap.addAntenna(cell, x, y)
		}
	}
}

func printMap(antennasMap *Map, withoutAntennas bool) {
	mp := make([][]rune, 0)
	// fill with dots
	for i := 0; i < antennasMap.height; i++ {
		row := make([]rune, antennasMap.width)
		for j := 0; j < antennasMap.width; j++ {
			row[j] = '.'
		}
		mp = append(mp, row)
	}
	if !withoutAntennas {
		// fill with Antennas
		for _, antennas := range antennasMap.antennas {
			for _, antenna := range antennas {
				mp[antenna.pos.y][antenna.pos.x] = antenna.frequency
			}
		}
	}
	// fill with Antinodes as "#" if they are not Antennas
	for _, antinode := range antennasMap.antinodes {
		if mp[antinode.y][antinode.x] == '.' {
			mp[antinode.y][antinode.x] = '#'
		}
	}
	util.PrintRuneMap(mp)
}
