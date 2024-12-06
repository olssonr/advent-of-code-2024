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
			numXmas += grid.searchXmas(letter, x, y)
		}
	}
	fmt.Println("Part1:", numXmas)

	numMasCrosses := 0
	for y, xs := range grid {
		for x, letter := range xs {
			if grid.searchMasCross(letter, x, y) {
				numMasCrosses++
			}
		}
	}
	fmt.Println("Part2:", numMasCrosses)
}

type Coordinate struct {
	x int
	y int
}

func newCoordinate(x, y int) Coordinate {
	return Coordinate{x: x, y: y}
}

func (coordinate Coordinate) isValid(grid Grid) bool {
	// TODO: Perhaps should calulcate this and store it in the grid? Make it a struct instead?
	height := len(grid)
	width := len(grid[0])

	return coordinate.x >= 0 && coordinate.y >= 0 && coordinate.x < width && coordinate.y < height
}

type Line []Coordinate

func newLine(coordinates ...Coordinate) (line Line) {
	line = append(line, coordinates...)

	return
}

func (line Line) isValid(grid Grid) bool {
	if utils.Any(line, func(coordinate Coordinate) bool {
		return !coordinate.isValid(grid)
	}) {
		return false
	} else {
		return true
	}
}

func (line Line) letters(grid Grid) (letters string) {
	for _, coordinate := range line {
		letters += grid[coordinate.y][coordinate.x]
	}

	return
}

type Grid [][]string

func newGrid(lines []string) (grid Grid) {
	grid = [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return
}

//*** Part1 ***

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
// SearchXmas returns the number of these occurences in the grid
func (grid Grid) searchXmas(letter string, x, y int) (numXmas int) {
	if letter != "X" {
		return 0
	}

	directions := grid.directions(x, y)

	for _, line := range directions {
		if grid.isXmas(line) {
			numXmas++
		}
	}

	return
}

// Directions returns a line for each direction
func (grid Grid) directions(x, y int) []Line {
	north := newLine(newCoordinate(x, y-1), newCoordinate(x, y-2), newCoordinate(x, y-3))
	northEast := newLine(newCoordinate(x+1, y-1), newCoordinate(x+2, y-2), newCoordinate(x+3, y-3))
	east := newLine(newCoordinate(x+1, y), newCoordinate(x+2, y), newCoordinate(x+3, y))
	southEast := newLine(newCoordinate(x+1, y+1), newCoordinate(x+2, y+2), newCoordinate(x+3, y+3))
	south := newLine(newCoordinate(x, y+1), newCoordinate(x, y+2), newCoordinate(x, y+3))
	southWest := newLine(newCoordinate(x-1, y+1), newCoordinate(x-2, y+2), newCoordinate(x-3, y+3))
	west := newLine(newCoordinate(x-1, y), newCoordinate(x-2, y), newCoordinate(x-3, y))
	northWest := newLine(newCoordinate(x-1, y-1), newCoordinate(x-2, y-2), newCoordinate(x-3, y-3))

	return []Line{north, northEast, east, southEast, south, southWest, west, northWest}
}

func (grid Grid) isXmas(line Line) bool {
	if !line.isValid(grid) {
		return false
	}

	return line.letters(grid) == "MAS"
}

// *** Part 2 ***

// Ok, so for part 2 we need to look for Two MAS that forms an X
//
// The following are all possilbe occurances of MAS Xs from a coordinate
// M.S S.M S.S M.M
// .A. .A. .A. .A.
// M.S S.M M.M S.S
// If we represent the first example as vector sets it would look like this [(0,0) (1,1) (2,2)] and [(0,2),(1,1), (2,0)]
//
// However we can simplify this problem by saying that only two combinations are available for each vector set, and we
// only need to check if the letters for both vector sets match any of the two combinations. If we simplify the vectors
// as a string it would look like this: (string1 == MAS || string1 == SAM) && (string2 == MAS || string2 == SAM)
// SearchMasCross returns true|false if the grid has a 3x3 sub grid starting in x,y that has a MAS cross (X-MAS)
func (grid Grid) searchMasCross(letter string, x, y int) bool {
	// This early return is just an optimization
	if letter == "A" {
		return false
	}

	line1, line2 := grid.cross(x, y)

	return grid.isMasCross(line1, line2)
}

// Cross returns two lines crossing each other in a 3x3 area
func (grid Grid) cross(x, y int) (line1, line2 []Coordinate) {
	line1 = newLine(newCoordinate(x, y), newCoordinate(x+1, y+1), newCoordinate(x+2, y+2))
	line2 = newLine(newCoordinate(x, y+2), newCoordinate(x+1, y+1), newCoordinate(x+2, y))

	return
}

func (grid Grid) isMasCross(line1, line2 Line) bool {
	if !line1.isValid(grid) || !line2.isValid(grid) {
		return false
	}

	if utils.Any([]Line{line1, line2}, func(line Line) bool {
		letters := line.letters(grid)
		return !(letters == "MAS" || letters == "SAM")
	}) {
		return false
	} else {
		return true
	}
}
