package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.PuzzleLines("day4/day4_puzzle_input.txt")
	grid := newGrid(lines)

	numXmas := 0
	for y, xs := range grid {
		for x, letter := range xs {
			numXmas += grid.search(letter, x, y)
		}
	}

	fmt.Println("Part1:", numXmas)
}

type Coordinate struct {
	x int
	y int
}

func newCoordinate(x, y int) Coordinate {
	return Coordinate{x: x, y: y}
}

type Grid [][]string

func newGrid(lines []string) (grid Grid) {
	grid = [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return
}

// Search returns the number of matches of XMAS that originates from x,y
//
// Fetches all directions e.g. 3 letters in each direction
// Then go through all such directions and check if they all have the correct letters
// If we can't fetch all letters (out of grid) then we don't need to check that direction
//
// The following are all possible occurences of XMAS from a coordinate
// S..S..S
// .A.A.A.
// ..MMM..
// SAMXMAS
// ..MMM..
// .A.A.A.
// S..S..S
//
// We have 8 directions e.g N, NE, E, SE, S, SW, W, NW
// NW N NE
// W. * E.
// SW S SE
func (grid Grid) search(letter string, x, y int) (numXmas int) {
	if letter != "X" {
		return 0
	}

	directions := grid.directions(x, y)

	for _, direction := range directions {
		if grid.isXmas(direction...) {
			numXmas++
		}
	}

	return
}

// Directions returns the coordinates for each direction
// If the coordinates for the direction are out of the grid the direction is skipped
// TODO: grid is not used yet, perhaps we should check the coordinates are valid here instead?
func (grid Grid) directions(x, y int) [][]Coordinate {
	north := []Coordinate{newCoordinate(x, y-1), newCoordinate(x, y-2), newCoordinate(x, y-3)}
	northEast := []Coordinate{newCoordinate(x+1, y-1), newCoordinate(x+2, y-2), newCoordinate(x+3, y-3)}
	east := []Coordinate{newCoordinate(x+1, y), newCoordinate(x+2, y), newCoordinate(x+3, y)}
	southEast := []Coordinate{newCoordinate(x+1, y+1), newCoordinate(x+2, y+2), newCoordinate(x+3, y+3)}
	south := []Coordinate{newCoordinate(x, y+1), newCoordinate(x, y+2), newCoordinate(x, y+3)}
	southWest := []Coordinate{newCoordinate(x-1, y+1), newCoordinate(x-2, y+2), newCoordinate(x-3, y+3)}
	west := []Coordinate{newCoordinate(x-1, y), newCoordinate(x-2, y), newCoordinate(x-3, y)}
	northWest := []Coordinate{newCoordinate(x-1, y-1), newCoordinate(x-2, y-2), newCoordinate(x-3, y-3)}

	return [][]Coordinate{north, northEast, east, southEast, south, southWest, west, northWest}
}

func (grid Grid) isXmas(coordinates ...Coordinate) bool {
	if utils.Any(coordinates, func(coordinate Coordinate) bool {
		return !grid.isValid(coordinate)
	}) {
		return false
	}

	first := coordinates[0]
	second := coordinates[1]
	third := coordinates[2]
	letters := []string{grid[first.y][first.x], grid[second.y][second.x], grid[third.y][third.x]}

	return strings.Join(letters, "") == "MAS"
}

func (grid Grid) isValid(coordinate Coordinate) bool {
	// TODO: Perhaps should calulcate this and store it in the grid? Make it a struct instead?
	height := len(grid)
	width := len(grid[0])

	return coordinate.x >= 0 && coordinate.y >= 0 && coordinate.x < width && coordinate.y < height
}
