package _14

import (
	"adventofcode2024/challenges/util"
	"fmt"
	"log"
)

type Robot struct {
	// x, y are the coordinates from the top left corner
	// vX, vY are the velocities for each iteration
	x, y, vX, vY int
}

func restroomRedoubt(filename string, width, height, iterations int) (int, int) {
	input, err := util.ReadLines(filename)
	if err != nil {
		log.Fatal(err)
	}
	robots := make([]Robot, 0)
	robots = parseRobots(input, robots)

	robots = moveRobots(robots, width, height, iterations)

	//log.Printf("Robots: %v", robots)
	// now I need to count quadrants

	//log.Printf("width / 2: %d", width/2)
	//log.Printf("height / 2: %d", height/2)
	numRobots := make([]int, 4)
	numRobots = countQuadrants(robots, width, height, numRobots)

	safetyFactor := 1
	for _, num := range numRobots {
		safetyFactor *= num
	}
	return safetyFactor, 0
}

func countQuadrants(robots []Robot, width int, height int, numRobots []int) []int {
	// we have 4 quadrants, so we need to check for each robot in which quadrant it is
	// the exact middle of the map does not count as any quadrant and is ignored!
	// because the map-width and map-height are always odd, we can use the following:
	quadrantWidth := width / 2
	quadrantHeight := height / 2
	for _, robot := range robots {
		if robot.x < quadrantWidth && robot.y < quadrantHeight {
			numRobots[0]++
		} else if robot.x > quadrantWidth && robot.y < quadrantHeight {
			numRobots[1]++
		} else if robot.x < quadrantWidth && robot.y > quadrantHeight {
			numRobots[2]++
		} else if robot.x > quadrantWidth && robot.y > quadrantHeight {
			numRobots[3]++
		}
	}
	return numRobots
}

func moveRobots(robots []Robot, width, height, iterations int) []Robot {
	for i := 0; i < iterations; i++ {
		for j, robot := range robots {
			robot = moveRobot(robot, width, height)
			robots[j] = robot
		}
	}
	return robots
}

func moveRobot(robot Robot, width, height int) Robot {
	// we have a map with the given width and height.

	// first, move the robot
	robot.x += robot.vX
	robot.y += robot.vY

	// check if the robot is outside the map in any direction, if so, teleport it to the other side
	if robot.x < 0 {
		robot.x += width
	}
	if robot.x >= width {
		robot.x -= width
	}
	if robot.y < 0 {
		robot.y += height
	}
	if robot.y >= height {
		robot.y -= height
	}
	return robot
}

func parseRobots(input []string, robots []Robot) []Robot {
	for _, line := range input {
		robot := Robot{}
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.vX, &robot.vY)
		if err != nil {
			log.Fatal(err)
		}
		robots = append(robots, robot)
	}
	return robots
}
