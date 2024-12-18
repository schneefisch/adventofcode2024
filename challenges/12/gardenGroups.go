package _12

import (
	"adventofcode2024/challenges/util"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Fence struct {
	x, y int
	dir  Direction
}

type Field struct {
	x         int
	y         int
	plant     rune
	perimeter []Direction
}

type Region struct {
	start     *Field
	fields    []*Field
	plant     rune
	area      int
	perimeter int
	sides     int
}

type Garden struct {
	gardenMap [][]rune
	regions   []*Region
}

func (g *Garden) isFieldInRegions(field *Field) bool {
	for _, region := range g.regions {
		for _, regionField := range region.fields {
			if regionField.x == field.x && regionField.y == field.y {
				return true
			}
		}
	}
	return false
}

func (g *Garden) isInGarden(f *Field) bool {
	return f.x >= 0 && f.x < len(g.gardenMap[0]) && f.y >= 0 && f.y < len(g.gardenMap)
}

func NewGarden() *Garden {
	return &Garden{
		gardenMap: make([][]rune, 0),
		regions:   make([]*Region, 0),
	}
}

func GardenGroups(filename string) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	garden := NewGarden()
	parseGarden(input, garden)
	identifyRegions(garden)

	//util.PrintRuneMap(garden.gardenMap)
	//log.Println("Regions:")
	result := 0
	for _, region := range garden.regions {
		//log.Printf("Region with plant %c has %d fields and %d perimeter\n", region.plant, region.area, region.perimeter)
		result += region.area * region.perimeter
	}

	result2 := 0

	return result, result2, nil
}

func identifyRegions(garden *Garden) {
	for h, row := range garden.gardenMap {
		for w, plant := range row {
			field := &Field{
				x:     w,
				y:     h,
				plant: plant,
			}
			field.perimeter = garden.getPerimeters(field)
			if garden.isFieldInRegions(field) {
				// jump to the next field, this is already contained in a region
				continue
			}
			// if it's not contained in a region already, create a new region and
			// find all fields that belong to it
			region := &Region{
				start:  field,
				plant:  plant,
				fields: make([]*Field, 0),
			}
			region.fields = append(region.fields, field)

			// find all fields that belong to this region
			region.fields = findFieldsInRegion(region.fields, field, garden)

			// calculate the area of the region
			region.area = len(region.fields)

			// calculate the perimeter of the region
			region.perimeter = 0
			for _, f := range region.fields {
				region.perimeter += len(f.perimeter)
			}

			// calculate the sides of the region
			region.sides = garden.getSides(region)

			// add new region to the garden
			garden.regions = append(garden.regions, region)
		}
	}
}

func findFieldsInRegion(fields []*Field, field *Field, garden *Garden) []*Field {
	// check all sides of the field
	// - is in range of the garden
	// - is same plant
	// - is not already in the fields slice
	// if all conditions are met, add the field to the fields slice
	// and call this function recursively with the new field
	surroundingFields := []*Field{
		{x: field.x, y: field.y - 1},
		{x: field.x + 1, y: field.y},
		{x: field.x, y: field.y + 1},
		{x: field.x - 1, y: field.y},
	}
	for _, f := range surroundingFields {
		if garden.isInGarden(f) && garden.belongsToRegion(field, f) && !contains(fields, f) {
			f.plant = garden.gardenMap[f.y][f.x]
			f.perimeter = garden.getPerimeters(f)
			fields = append(fields, f)
			newFields := findFieldsInRegion(fields, f, garden)
			fields = appendUnique(fields, newFields)
		}
	}

	// return the fields slice
	return fields
}

func appendUnique(original []*Field, new []*Field) []*Field {
	for _, field := range new {
		if !contains(original, field) {
			original = append(original, field)
		}
	}
	return original
}

// contains checks if a field is already in the fields slice
func contains(fields []*Field, search *Field) bool {
	for _, field := range fields {
		if field.x == search.x && field.y == search.y {
			return true
		}
	}
	return false
}

func (g *Garden) belongsToRegion(origin *Field, newField *Field) bool {
	return g.isInGarden(newField) && g.samePlant(origin, newField)
}

func (g *Garden) samePlant(field1 *Field, field2 *Field) bool {
	if !g.isInGarden(field1) || !g.isInGarden(field2) {
		return false
	}
	plant1 := g.gardenMap[field1.y][field1.x]
	plant2 := g.gardenMap[field2.y][field2.x]
	return plant1 == plant2
}

func (g *Garden) getPerimeters(f *Field) []Direction {
	perimeters := make([]Direction, 0)
	top := &Field{x: f.x, y: f.y - 1}
	right := &Field{x: f.x + 1, y: f.y}
	bottom := &Field{x: f.x, y: f.y + 1}
	left := &Field{x: f.x - 1, y: f.y}

	if !g.isInGarden(top) || !g.samePlant(f, top) {
		perimeters = append(perimeters, Up)
	}
	if !g.isInGarden(right) || !g.samePlant(f, right) {
		perimeters = append(perimeters, Right)
	}
	if !g.isInGarden(bottom) || !g.samePlant(f, bottom) {
		perimeters = append(perimeters, Down)
	}
	if !g.isInGarden(left) || !g.samePlant(f, left) {
		perimeters = append(perimeters, Left)
	}
	return perimeters
}

// getSides finds the sides by walking the perimeter of the region.
// Since we are iterating row by row (top-down) and field by field (left-right)
// we can assume, that the start-field is always in the top-left corner of the region
// we can walk the perimeter by walking clock-wise around the region and increase the
// side counter every time we change the direction
func (g *Garden) getSides(region *Region) int {
	sides := 0
	visited := make([]Fence, 0)
	// walk the fields in the region
	for _, field := range region.fields {
		// also if it has no perimeter, it's not part of the perimeter
		if len(field.perimeter) == 0 {
			continue
		}

		var allPerimeters []Fence
		for _, pe := range field.perimeter {
			allPerimeters = append(allPerimeters, Fence{x: field.x, y: field.y, dir: pe})
		}

		for _, perimeter := range allPerimeters {
			// check if it's already visited
			if containsFence(visited, perimeter) {
				continue
			}

			// if it's not visited, walk the perimeter
			sides += g.walkPerimeter(perimeter, visited)
		}

		// if the field has two sides, it might be a corner or might have perimeter on both sides
	}
	return sides
}

func (g *Garden) walkPerimeter(perimeter Fence, visited []Fence) int {
	if containsFence(visited, perimeter) {
		return 0
	}
	visited = append(visited, perimeter)
	sides := 1
	next, corner := g.nextPerimeter(perimeter)
	if corner {
		sides++
	}
	sides += g.walkPerimeter(next, visited)
	return sides
}

func (g *Garden) nextPerimeter(perimeter Fence) (Fence, bool) {
	var next Fence
	// if the current fence has also a clock-wise perimeter, we can go around the
	// corner and increase the sides counter

	switch perimeter.dir {
	case Up:
		next = Fence{x: perimeter.x + 1, y: perimeter.y, dir: Up}
	case Right:
		next = Fence{x: perimeter.x, y: perimeter.y + 1, dir: Right}
	case Down:
		next = Fence{x: perimeter.x - 1, y: perimeter.y, dir: Down}
	case Left:
		next = Fence{x: perimeter.x, y: perimeter.y - 1, dir: Left}
	}
	// check if the next field is in the garden
	if g.isInGarden(&Field{x: next.x, y: next.y}) &&
		g.samePlant(&Field{x: perimeter.x, y: perimeter.y}, &Field{x: next.x, y: next.y}) {
		return next, false
	}
	// get clock-wise corner
	switch perimeter.dir {
	case Up:
		next = Fence{x: perimeter.x + 1, y: perimeter.y, dir: Right}
	case Right:
		next = Fence{x: perimeter.x, y: perimeter.y + 1, dir: Down}
	case Down:
		next = Fence{x: perimeter.x - 1, y: perimeter.y, dir: Left}
	case Left:
		next = Fence{x: perimeter.x, y: perimeter.y - 1, dir: Up}
	}

	return next, true
}

func containsFence(visited []Fence, perimeter Fence) bool {
	for _, fence := range visited {
		if fence.x == perimeter.x && fence.y == perimeter.y && fence.dir == perimeter.dir {
			return true
		}
	}
	return false
}

func parseGarden(input []string, garden *Garden) {
	garden.gardenMap = make([][]rune, len(input))
	for i, row := range input {
		fields := make([]rune, len(row))
		for j, field := range row {
			fields[j] = field
		}
		garden.gardenMap[i] = fields
	}
}
