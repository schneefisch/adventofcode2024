package _20

import "adventofcode2024/challenges/util"

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

type Cheat struct {
	posX, posY int
	scoreSaved int
}

func appendUnique(cheats []Cheat, cheat Cheat) []Cheat {
	for _, existing := range cheats {
		if existing.posX == cheat.posX && existing.posY == cheat.posY {
			return cheats
		}
	}
	return append(cheats, cheat)
}

func RaceCondition(filename string, threshold int) (int, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	maze := &util.Maze{}
	maze.Parse(input)

	_, visited := findShortestPath(maze)

	cheats := make([]Cheat, 0)
	cheats = findJumpPositions(maze, cheats, visited)

	// count cheats that save steps equal or above the threshold
	cheatCount := 0
	for _, cheat := range cheats {
		if cheat.scoreSaved >= threshold {
			cheatCount++
		}
	}

	cheatsExtended := findJumpPositionsExtended(maze, visited)
	cheatCountExtended := 0
	for _, cheat := range cheatsExtended {
		if cheat >= threshold {
			cheatCountExtended++
		}
	}

	// Implement me
	return cheatCount, cheatCountExtended, nil
}

// findJumpPositionsExtended finds all the "wall" tiles that are adjacent to two visited tiles
// ToDo: still not working, is finding too many cheats
func findJumpPositionsExtended(maze *util.Maze, visited *Visited) []int {
	cheats := make([]int, 0)
	for _, node := range visited.nodes {
		reachablePositions := findReachablePositions(maze, node.pos)

		// for each reachable position find the cheat score
		for _, targetPos := range reachablePositions {
			if targetScore, found := visited.find(targetPos); found {
				// calculate the shortest distance from the start to the target no matter the walls
				distance := distance(node.pos, targetPos)

				cheatScore := node.score - targetScore - distance
				if cheatScore > 0 {
					cheats = append(cheats, cheatScore)
				}
			}
		}
	}
	return cheats
}

// distance calculates the shortest distance from the start to the target in a 2D grid
func distance(position, targetPos util.Position) int {
	distX := position.X - targetPos.X
	distY := position.Y - targetPos.Y
	if distX < 0 {
		distX = -distX
	}
	if distY < 0 {
		distY = -distY
	}
	return distX + distY
}

// findReachablePositions finds all the positions that are reachable from the given position
// using the Manhattan distance
func findReachablePositions(maze *util.Maze, position util.Position) []util.Position {
	reachable := make([]util.Position, 0)
	maxJumpDistance := 20

	// check all positions within the manhattan distance of maxJumpDistance
	for x := position.X - maxJumpDistance; x <= position.X+maxJumpDistance; x++ {
		for y := position.Y - maxJumpDistance; y <= position.Y+maxJumpDistance; y++ {
			newPos := util.Position{X: x, Y: y}
			if maze.IsInGrid(newPos) && maze.TileAt(newPos).Kind != util.Wall {
				reachable = append(reachable, newPos)
			}
		}
	}

	return reachable
}

// findJumpPositions finds all the "wall" tiles that are adjacent to two visited tiles
func findJumpPositions(maze *util.Maze, cheats []Cheat, visited *Visited) []Cheat {
	directions := []util.Direction{util.North, util.East, util.South, util.West}
	for _, node := range visited.nodes {
		for _, dir := range directions {
			newPos := nextPos(node.pos, dir)
			if maze.IsInGrid(newPos) && maze.TileAt(newPos).Kind == util.Wall {
				// check if the next position in the same direction after the wall is also a visited tile
				posAfterWall := nextPos(newPos, dir)
				if maze.IsInGrid(posAfterWall) && maze.TileAt(posAfterWall).Kind != util.Wall {
					jumpScore, found := visited.find(posAfterWall)
					if found {
						// in addition to the possible jump position, I can also directly calculate the score of the cheat
						// by subtracting the score of the visited tile from the score of the current tile
						// need to also subtract 2 from the cheatScore, because two steps are requried for the jump
						cheatScore := jumpScore - node.score - 2
						if cheatScore > 0 {
							cheats = appendUnique(cheats, Cheat{newPos.X, newPos.Y, cheatScore})
						}
					}
				}
			}
		}
	}
	return cheats
}

// findShortestPath uses a depthFirstSearch to find the shortest path in the maze
// returns the score of the shortest path, equivalent of the number of steps
func findShortestPath(maze *util.Maze) (int, *Visited) {
	stack := []Node{{maze.Start, 0}}
	visited := &Visited{nodes: []Node{}}
	visited.add(maze.Start, 0)
	minSteps := -1
	directions := []util.Direction{util.North, util.East, util.South, util.West}

	for len(stack) > 0 {
		// pop the "most recent" entry from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if maze.Grid[current.pos.Y][current.pos.X].Kind == util.End {
			if minSteps == -1 || current.score < minSteps {
				minSteps = current.score
			}
			continue
		}

		for _, dir := range directions {
			newPos := nextPos(current.pos, dir)

			if maze.IsInGrid(newPos) {
				tile := maze.TileAt(newPos)
				newScore := current.score + 1
				if tile.Kind != util.Wall {
					if oldScore, found := visited.find(newPos); !found || newScore < oldScore {
						visited.add(newPos, newScore)
						stack = append(stack, Node{newPos, newScore})
					}
				}
			}
		}
	}

	return minSteps, visited
}

func nextPos(current util.Position, dir util.Direction) util.Position {
	newPos := util.Position{X: current.X, Y: current.Y, Dir: dir}
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
