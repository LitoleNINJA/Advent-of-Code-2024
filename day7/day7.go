package day7

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func DAY7() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	result := partOne(lines)
	duration := time.Since(start) 
	fmt.Printf("Part 1 : %d, Time Taken : %v\n", result, duration)

	start = time.Now()
	result = partTwo(lines)
	duration = time.Since(start) 
	fmt.Printf("Part 2 : %d, Time Taken : %v\n", result, duration)
}

func partOne(lines []string) int64 {
	ans := int64(0)

	for _, line := range lines {
		target, values, _ := strings.Cut(line, ":")
		if res, ok := isGood(target, values, false); ok {
			ans += res
		}
	}

	return ans
}

func partTwo(lines []string) int64 {
	ans := int64(0)

	for _, line := range lines {
		target, values, _ := strings.Cut(line, ":")
		if res, ok := isGood(target, values, true); ok {
			ans += res
		}
	}

	return ans
}

func isGood(t string, v string, isExtraOP bool) (int64, bool) {
	target, _ := strconv.ParseInt(t, 10, 64)
	values, _ := utils.StringSliceToInt(strings.Fields(v))

	slices.Reverse(values)
	allValues := possibleValues(values, isExtraOP)
	// fmt.Println(allValues)

	if slices.Contains(allValues, target) {
		return target, true
	} else {
		return 0, false
	}

}

func possibleValues(values []int64, isExtraOp bool) []int64 {
	if len(values) == 1 {
		return values
	}

	ans := make([]int64, 0)
	rem := possibleValues(values[1:], isExtraOp)

	for _, val := range rem {
		ans = append(ans, values[0]+val)
		ans = append(ans, values[0]*val)
		if isExtraOp {
			ans = append(ans, join(val, values[0]))
		}
	}

	return ans
}

func join(x int64, y int64) int64 {
	var sb strings.Builder
	sb.WriteString(strconv.FormatInt(x, 10))
	sb.WriteString(strconv.FormatInt(y, 10))

	result, _ := strconv.ParseInt(sb.String(), 10, 64)
	return result
}
