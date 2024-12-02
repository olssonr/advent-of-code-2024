package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
