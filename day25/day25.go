package day25

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"time"
)

func DAY25() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	result := partOne(lines)
	duration := time.Since(start)
	fmt.Printf("Part 1 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())

	/*start = time.Now()
	result = partTwo(lines)
	duration = time.Since(start)
	fmt.Printf("Part 2 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())*/
}

var locks [][5]int
var keys [][5]int

func partOne(lines []string) int64 {
	ans := int64(0)

	locks = make([][5]int, 0)
	keys = make([][5]int, 0)
	buf := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			val, isLock := getPinHeights(buf)
			if isLock {
				locks = append(locks, val)
			} else {
				keys = append(keys, val)
			}
			buf = make([]string, 0)
			continue
		}
		buf = append(buf, line)
	}

	for _, lock := range locks {
		for _, key := range keys {
			if isGoodFit(lock, key) {
				ans++
			}
		}
	}

	return ans
}

func getPinHeights(buf []string) ([5]int, bool) {
	var ans [5]int
	for i := range 5 {
		ans[i] = -1
	}

	for _, row := range buf {
		for i, ch := range row {
			if ch == '#' {
				ans[i]++
			}
		}
	}

	return ans, buf[0] == "#####"
}

func isGoodFit(lock [5]int, key [5]int) bool {
	for i := range 5 {
		if lock[i] + key[i] > 5 {
			return false
		}
	}
	return true
}