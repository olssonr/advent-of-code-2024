package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PuzzleText(filename string) string {
	content, error := os.ReadFile(filename)
	if error != nil {
		fmt.Println("Error reading the file:", error)
	}

	return string(content)
}

func PuzzleLines(filename string) (lines []string) {
	lines = strings.Split(PuzzleText(filename), "\n")

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
