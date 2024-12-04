package day4

import (
	"aoc2024/utils"
	"fmt"
	"os"
)

func DAY4() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 : %d\n", partOne(lines))
	fmt.Printf("Part 2 : %d\n", partTwo(lines))
}

func partOne(lines []string) int {
	ans := 0
	for i, line := range lines {
		for j, ch := range line {
			if ch == 'X' {
				ans += countXMAS(lines, i, j)
			}
		}
	}

	return ans
}

func partTwo(lines []string) int {
	ans := 0
	for i, line := range lines {
		for j, ch := range line {
			if ch == 'A' && isCrossMAS(lines, i, j) {
				ans++
			}
		}
	}

	return ans
}

func countXMAS(lines []string, i int, j int) int {
	count := 0
	dir := [8][2]int{
		{0, 1},   // right
		{1, 1},   // down-right
		{1, 0},   // down
		{1, -1},  // down-left
		{0, -1},  // left
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
	}

	for _, d := range dir {
		if isGood(lines, i, j, d[0], d[1]) {
			count++
		}
	}

	return count
}

func isGood(lines []string, i int, j int, dx int, dy int) bool {
	xmas := "XMAS"
	pos := 0
	for {
		if !inRange(lines, i, j) {
			return false
		}

		if lines[i][j] != xmas[pos] {
			return false
		}

		i += dx
		j += dy
		pos++
		if pos >= 4 {
			return true
		}
	}
}

func inRange(lines []string, i int, j int) bool {
	return (i >= 0 && i < len(lines) && j >= 0 && j < len(lines[0]))
}

func isCrossMAS(lines []string, i int, j int) bool {
	if inRange(lines, i-1, j-1) && inRange(lines, i+1, j+1) && inRange(lines, i+1, j-1) && inRange(lines, i-1, j+1) {
		if isMA(lines[i-1][j-1], lines[i+1][j+1]) && isMA(lines[i+1][j-1], lines[i-1][j+1]) {
			return true
		}
	}

	return false
}

func isMA(x string, y string) {
	
}
