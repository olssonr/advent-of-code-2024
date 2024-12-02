package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Put some utility functions before the main function
// consider moving this to a separate file so can be used in next days

func PuzzleLines(filename string) (lines []string) {
	content, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println("Error reading the file:", error)
	}
	lines = strings.Split(string(content), "\n")

	return
}

func AbsInt(number int) int {
	if number > -1 {
		return number
	}

	return -number
}

func Toi(text string) (number int) {
	number, _ = strconv.Atoi(text)

	return
}

func Count(list []int, element int) (count int) {
	for _, listItem := range list {
		if listItem == element {
			count++
		}
	}

	return
}
