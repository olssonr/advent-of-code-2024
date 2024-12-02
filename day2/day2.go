package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	lines := utils.PuzzleLines("day2/day2_puzzle_input.txt")

	var reports []Report
	for _, line := range lines {
		reports = append(reports, newReport(line))
	}

	var failedReports []Report
	for _, report := range reports {
		if !isSafe(report) {
			failedReports = append(failedReports, report)
		}
	}

	numSafeReports := len(reports) - len(failedReports)
	fmt.Println("Part 1: ", numSafeReports)

	numTolerableReports := 0
	for _, failedReport := range failedReports {
		if tolerable(failedReport) {
			numTolerableReports++
		}
	}

	fmt.Println("Part 2: ", numSafeReports+numTolerableReports)
}

func tolerable(report Report) bool {
	for i := 0; i < len(report.levels); i++ {
		reportVariation := copyReport(report)
		reportVariation.levels = slices.Delete(reportVariation.levels, i, i+1)

		if isSafe(reportVariation) {
			return true
		}
	}

	return false
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

func copyReport(report Report) (reportCopy Report) {
	levelsCopy := make([]int, len(report.levels))
	copy(levelsCopy, report.levels)
	reportCopy = Report{levels: levelsCopy}

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
