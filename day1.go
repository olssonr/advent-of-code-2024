package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Put some utility functions before the main function
// consider moving this to a separate file so can be used in next days

func puzzleLines(filename string) (lines []string) {
	content, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println("Error reading the file:", error)
	}
	lines = strings.Split(string(content), "\n")

	return
}

func absInt(number int) int {
	if number > -1 {
		return number
	}

	return -number
}

func toi(text string) (number int) {
	number, _ = strconv.Atoi(text)

	return
}

func main() {
	lines := puzzleLines("day1_puzzle_input.txt")

	list1, list2 := lists(lines)
	slices.Sort(list1)
	slices.Sort(list2)

	totalDistance := sumPairDifference(list1, list2)
	fmt.Println(totalDistance)
}

func sumPairDifference(list1 []int, list2 []int) (distanceSum int) {
	for i := 0; i < len(list1); i++ {
		distanceSum += absInt(list1[i] - list2[i])
	}

	return
}

func lists(lines []string) (list1, list2 []int) {
	for _, line := range lines {
		fields := strings.Fields(line)

		list1 = append(list1, toi(fields[0]))
		list2 = append(list2, toi(fields[1]))
	}

	return
}
