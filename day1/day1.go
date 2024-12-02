package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	lines := utils.PuzzleLines("day1/day1_puzzle_input.txt")

	list1, list2 := lists(lines)
	slices.Sort(list1)
	slices.Sort(list2)

	totalDistance := sumPairDifference(list1, list2)
	fmt.Println("Part 1:", totalDistance)

	similarityScore := 0
	for _, number := range list1 {
		similarityScore += number * utils.Count(list2, number)
	}

	fmt.Println("Part 2:", similarityScore)
}

func sumPairDifference(list1 []int, list2 []int) (distanceSum int) {
	for i := 0; i < len(list1); i++ {
		distanceSum += utils.AbsInt(list1[i] - list2[i])
	}

	return
}

func lists(lines []string) (list1, list2 []int) {
	for _, line := range lines {
		fields := strings.Fields(line)

		list1 = append(list1, utils.Toi(fields[0]))
		list2 = append(list2, utils.Toi(fields[1]))
	}

	return
}
