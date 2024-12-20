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
	isInGrid := true
	isInLoop := false
	for isInGrid && !isInLoop {
		//labMap.grid.print()
		isInGrid, isInLoop = guard.patrol(labMap)
	}

	fmt.Println("Part 1:", labMap.grid.numDistinctPositions())

	// Part2
	// First, draw the map with all visited positions marked (already done in part1)
	// Then try adding an obstacle to each visited position
	// Restart the patrol loop now and check for two things:
	// 1. If patrol returns false, the guard is about to go out of the grid, and there is no loop
	// 2. If we visit the same positition twice and we are facing the same direction we are in a loop
	numDifferentObstructionPositions := 0
	for _, position := range labMap.grid.visitedPositions() {
		// We need to work a new copy of the map, would a deep copy be more efficient?
		mapWithObstacle := newLabMap(lines)
		mapWithObstacle.grid[position.y][position.x] = "O"
		guard := newGuard(up, mapWithObstacle.grid.findGuardPosition())

		isInGrid := true
		isInLoop := false
		for isInGrid && !isInLoop {
			//mapWithObstacle.grid.print()
			isInGrid, isInLoop = guard.patrol(mapWithObstacle)
		}

		if isInLoop {
			numDifferentObstructionPositions++
			//fmt.Println(numDifferentObstructionPositions)
		}
	}

	fmt.Println("Part 2:", numDifferentObstructionPositions)

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

func (guard *Guard) patrol(labMap LabMap) (inMap bool, inLoop bool) {
	positionAhead := guard.positionAhead()
	if !positionAhead.isValid(labMap) {
		return false, false
	}

	if labMap.positionHasObstacle(positionAhead) || labMap.positionHasNewObstacle(positionAhead) {
		guard.turn()
	} else {
		guard.position = positionAhead

		oldMark := labMap.grid.getMark(guard.position)
		labMap.markVisited(*guard)
		newMark := labMap.grid.getMark(guard.position)

		// Another way to discover the loop would be to store the guard (coordinate, facing) in a set and check that
		// But this was more fun :D
		if oldMark == newMark {
			return true, true
		}
	}

	return true, false
}

func (guard Guard) mark() (mark string) {
	switch guard.facing {
	case up:
		mark = "^"
	case down:
		mark = "v"
	case left:
		mark = "<"
	case right:
		mark = ">"
	}

	return
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

func (labMap *LabMap) markVisited(guard Guard) {
	labMap.grid[guard.position.y][guard.position.x] = guard.mark()
}

func (labMap LabMap) positionHasObstacle(coordinate Coordinate) bool {
	return labMap.grid[coordinate.y][coordinate.x] == "#"
}

func (labMap LabMap) positionHasNewObstacle(coordinate Coordinate) bool {
	return labMap.grid[coordinate.y][coordinate.x] == "O"
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
			if character == "^" || character == "v" || character == "<" || character == ">" {
				sum++
			}
		}
	}

	return
}

func (grid Grid) visitedPositions() (positions []Coordinate) {
	for y, line := range grid {
		for x, character := range line {
			if character == "^" || character == "v" || character == "<" || character == ">" {
				positions = append(positions, newCoordinate(x, y))
			}
		}
	}

	return
}

func (grid Grid) getMark(coordinate Coordinate) string {
	return grid[coordinate.y][coordinate.x]
}

func (grid Grid) print() {
	fmt.Println()
	for _, line := range grid {
		fmt.Println(line)
	}
}
