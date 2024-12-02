package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc2024/utils"
)

func DAY2() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	// fmt.Printf("Part 1 : %d\n", partOne(lines))
	fmt.Printf("Part 2 : %d\n", partTwo(lines))
}

func partOne(lines []string) int {
	var ans int = 0

	for _, line := range lines {
		levels := strings.Split(line, " ")

		if isGood(levels) {
			ans++
		}
	}

	return ans
}

func isGood(levels []string) bool {
	n := len(levels)

	diff := make([]int, 0)
	for i := 1; i < n; i++ {
		diff = append(diff, gap(levels[i], levels[i-1]))
	}

	isPositive := isPositive(diff)
	for _, d := range diff {
		if (isPositive && d <= 0) || (!isPositive && d >= 0) || utils.Abs(d) > 3 || utils.Abs(d) < 1 {
			return false
		}
	}

	// fmt.Println(levels)
	return true
}

func gap(x string, y string) int {
	x_int, _ := strconv.ParseInt(x, 10, 64)
	y_int, _ := strconv.ParseInt(y, 10, 64)

	return int(x_int - y_int)
}

func isPositive(diff []int) bool {
	pos, neg := 0, 0

	for _, d := range diff {
		if d >= 0 {
			pos++
		} else {
			neg++
		}
	}

	return pos > neg
}

func partTwo(lines []string) int {
	var ans int = 0

	for _, line := range lines {
		levels := strings.Split(line, " ")

		for i := 0; i < len(levels); i++ {
			newLevels := make([]string, len(levels)-1)
			copy(newLevels, levels[:i])
			copy(newLevels[i:], levels[i+1:])
			// fmt.Println(newLevels)
			if isGood(newLevels) {
				ans++
				break
			}
		}
	}

	return ans
}
