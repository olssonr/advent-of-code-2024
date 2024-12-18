package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func main() {
	lines := utils.PuzzleLines("day7/day7_puzzle_input.txt")
	equations := []Equation{}
	for _, line := range lines {
		equations = append(equations, newEquation(line))
	}

	// Part1:
	// 1. build all possible permutations for each equation
	// 2. Go through all permutations per equation and check if an equation as a possible solution
	// 3. Sum the test value of all equations with possible solutions
	//
	// We can think of this as a binary tree
	// Example 3267: 81 40 27
	//
	//              81
	//           /     \
	//        +40      *40
	//        / \      /  \
	//     +27  *27  +27 *27
	//
	// And we can solve this with a recursive function

	totalCalibrationResult := 0
	for _, equation := range equations {
		sums := configurations(equation.numbers[0], equation.numbers[1:])
		if slices.Contains(sums, equation.testValue) {
			totalCalibrationResult += equation.testValue
		}
	}
	fmt.Println("Part1:", totalCalibrationResult)
}

func configurations(acc int, numbers []int) []int {
	if len(numbers) < 1 {
		// Base case: return the accumulated result in a slice
		return []int{acc}
	}

	head := numbers[0]
	tail := numbers[1:]

	// Recursively calculate the configurations and merge the results in a flat slice
	return append(configurations(acc+head, tail), configurations(acc*head, tail)...)
}

type Equation struct {
	testValue int
	numbers   []int
}

func newEquation(line string) (equation Equation) {
	re := regexp.MustCompile(`(\d*):\s([\d\s]*)`)
	matches := re.FindStringSubmatch(line)

	equation.testValue = utils.Toi(matches[1])
	for _, number := range strings.Split(matches[2], " ") {
		equation.numbers = append(equation.numbers, utils.Toi(number))
	}

	return
}
