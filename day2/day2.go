package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.PuzzleLines("day2/day2_puzzle_input.txt")

	var reports []Report
	for _, line := range lines {
		reports = append(reports, newReport(line))
	}

	numSafeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafeReports++
		}
	}

	fmt.Println(numSafeReports)
}

type Report struct {
	levels []int
}

func newReport(line string) (report Report) {
	fields := strings.Fields(line)
	for _, field := range fields {
		report.levels = append(report.levels, utils.Toi(field))
	}

	return
}

func isSafe(report Report) bool {
	var increasing bool
	var decreasing bool
	for i, level := range report.levels {
		if i == 0 {
			continue
		}
		previousLevel := report.levels[i-1]

		if level == previousLevel {
			return false
		}

		if level > previousLevel {
			increasing = true
		} else {
			decreasing = true
		}
		if increasing && decreasing {
			return false
		}

		difference := utils.AbsInt(previousLevel - level)
		if difference > 3 {
			return false
		}
	}

	return true
}
