package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"regexp"
)

func main() {
	corruptedMemory := utils.PuzzleText("day3/day3_puzzle_input.txt")
	fmt.Println("Part1:", findMuls(corruptedMemory).sum())

	modifiedCorruptedMemory := applyDoDontsInstructions(corruptedMemory)
	fmt.Println("Part2:", findMuls(modifiedCorruptedMemory).sum())
}

// ApplyDoDontsInstructions removes disabled instructions from the string
func applyDoDontsInstructions(memory string) string {
	// Newlines are not part of . so we use [\s\S] instead
	// In order to make do optional we match with either do or end ($)
	re := regexp.MustCompile(`don't\(\)[\s\S]*?(do\(\)|$)`)
	return re.ReplaceAllString(memory, "")
}

type Muls []Mul

func findMuls(memory string) (muls Muls) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	instructionMatches := re.FindAllStringSubmatch(memory, -1)

	for _, instructionMatch := range instructionMatches {
		muls = append(muls, newMul(utils.Toi(instructionMatch[1]), utils.Toi(instructionMatch[2])))
	}

	return
}

func (muls Muls) sum() (total int) {
	for _, mul := range muls {
		total += mul.calculate()
	}

	return
}

type Mul struct {
	arg1 int
	arg2 int
}

func newMul(arg1, arg2 int) Mul {
	return Mul{arg1: arg1, arg2: arg2}
}

func (mul Mul) calculate() int {
	return mul.arg1 * mul.arg2
}
