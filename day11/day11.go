package day11

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func DAY11() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	result := partOne(lines)
	duration := time.Since(start)
	fmt.Printf("Part 1 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())

	start = time.Now()
	result = partTwo(lines)
	duration = time.Since(start)
	fmt.Printf("Part 2 : %-6d | Time Taken : %.6f seconds\n", result, duration.Seconds())
}

type Pair struct {
	x, y int64
}

var dp map[Pair]int64

func partOne(lines []string) int64 {
	ans := int64(0)
	dp = make(map[Pair]int64, 0)

	values, _ := utils.StringSliceToInt(strings.Fields(lines[0]))
	for _, val := range values {
		ans += recurse(val, 25)
	}

	return ans
}

func partTwo(lines []string) int64 {
	ans := int64(0)
	dp = make(map[Pair]int64, 0)

	values, _ := utils.StringSliceToInt(strings.Fields(lines[0]))
	for _, val := range values {
		ans += recurse(val, 75)
	}

	return ans
}

func splitVal(val int64) (int64, int64) {
	valString := strconv.FormatInt(val, 10)
	mid := len(valString) / 2
	left, _ := strconv.ParseInt(valString[:mid], 10, 64)
	right, _ := strconv.ParseInt(valString[mid:], 10, 64)
	return left, right
}

func recurse(val int64, k int64) int64 {
	if k == 0 {
		return 1
	}

	if val == 0 {
		dp[Pair{val, k}] = recurse(1, k-1)
		return dp[Pair{val, k}]
	}

	if v, ok := dp[Pair{val, k}]; ok {
		return v
	}

	ans := int64(0)
	if utils.DigitCount(val)%2 == 0 {
		left, right := splitVal(val)
		ans = recurse(left, k-1) + recurse(right, k-1)
	} else {
		ans = recurse(val*2024, k-1)
	}

	dp[Pair{val, k}] = ans
	return ans
}
