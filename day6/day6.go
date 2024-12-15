package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.PuzzleLines("day6/day6_puzzle_input.txt")
	labMap := newLabMap(lines)
	guard := newGuard(up, labMap.grid.findGuardPosition())

	//labMap.grid.print()
	for guard.patrol(labMap) {
		//labMap.grid.print()
	}

	fmt.Println("Part 1:", labMap.grid.numDistinctPositions())
}

const (
	up    = 0
	down  = 180
	left  = 270
	right = 90
)

type Coordinate struct {
	x int
	y int
}

func newCoordinate(x, y int) Coordinate {
	return Coordinate{x: x, y: y}
}

func (coordinate Coordinate) isValid(labMap LabMap) bool {
	return coordinate.x >= 0 && coordinate.y >= 0 && coordinate.x < labMap.width && coordinate.y < labMap.height
}

func (coordinate Coordinate) up() Coordinate {
	coordinate.y -= 1

	return coordinate
}

func (coordinate Coordinate) right() Coordinate {
	coordinate.x += 1

	return coordinate
}

func (coordinate Coordinate) down() Coordinate {
	coordinate.y += 1

	return coordinate
}

func (coordinate Coordinate) left() Coordinate {
	coordinate.x -= 1

	return coordinate
}

type Guard struct {
	facing   int
	position Coordinate
}

func newGuard(facing int, position Coordinate) (guard Guard) {
	guard.facing = facing
	guard.position = position

	return
}

func (guard Guard) positionAhead() (coordinate Coordinate) {
	switch guard.facing {
	case up:
		coordinate = guard.position.up()
	case down:
		coordinate = guard.position.down()
	case left:
		coordinate = guard.position.left()
	case right:
		coordinate = guard.position.right()
	}

	return
}

func (guard *Guard) turn() {
	guard.facing = (guard.facing + 90) % 360
}

func (guard *Guard) patrol(labMap LabMap) bool {
	positionAhead := guard.positionAhead()
	if !positionAhead.isValid(labMap) {
		return false
	}

	if labMap.positionHasObstacle(positionAhead) {
		guard.turn()
	} else {
		guard.position = positionAhead
	}

	labMap.markVisited(guard.position)

	return true
}

type LabMap struct {
	grid   Grid
	width  int
	height int
}

func newLabMap(lines []string) (labMap LabMap) {
	labMap.grid = newGrid(lines)
	labMap.width = len(labMap.grid)
	labMap.height = len(labMap.grid[0])

	return
}

func (labMap *LabMap) markVisited(cooardinate Coordinate) {
	labMap.grid[cooardinate.y][cooardinate.x] = "X"
}

func (labMap LabMap) positionHasObstacle(coordinate Coordinate) bool {
	return labMap.grid[coordinate.y][coordinate.x] == "#"
}

type Grid [][]string

func newGrid(lines []string) (grid Grid) {
	grid = [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return
}

func (grid Grid) findGuardPosition() (coordinate Coordinate) {
	for y, xs := range grid {
		for x, character := range xs {
			if character == "^" {
				coordinate = newCoordinate(x, y)
			}
		}
	}

	return
}

func (grid Grid) numDistinctPositions() (sum int) {
	for _, line := range grid {
		for _, character := range line {
			if character == "X" || character == "^" {
				sum++
			}
		}
	}

	return
}

func (grid Grid) print() {
	fmt.Println()
	for _, line := range grid {
		fmt.Println(line)
	}
}
