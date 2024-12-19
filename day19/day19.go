package day19

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func DAY19() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	result := partOne(lines)
	duration := time.Since(start)
	fmt.Printf("Part 1 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())

	// start = time.Now()
	// result = partTwo(lines)
	// duration = time.Since(start)
	// fmt.Printf("Part 2 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())
}

var patterns []string
var dp map[string]int

func partOne(lines []string) int64 {
	ans := int64(0)

	patterns = strings.Split(lines[0], ", ")
	dp = make(map[string]int, 0)
	dp[""] = 1
	for _, p := range patterns {
		recurse(p)
	}

	for i := 2; i < len(lines); i++ {
		ans += int64(recurse(lines[i]))
	}

	return ans
}

func recurse(line string) int {
	if val, ok := dp[line]; ok {
		return val
	}

	ans := 0
	for i := range len(line) {
		cur := line[:i+1]
		if slices.Contains(patterns, cur) {
			ans += recurse(line[i+1:])
		}
	}

	dp[line] = ans
	return ans
}
