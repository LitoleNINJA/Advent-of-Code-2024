package day10

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"time"
)

func DAY10() {
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

func partOne(lines []string) int64 {
	ans := int64(0)

	for i, line := range lines {
		for j, ch := range line {
			if ch == '0' {
				ans += score(lines, i, j)
			}
		}
	}

	return ans
}

func partTwo(lines []string) int64 {
	ans := int64(0)

	for i, line := range lines {
		for j, ch := range line {
			if ch == '0' {
				ans += rating(lines, i, j)
			}
		}
	}

	return ans
}

func score(lines []string, i int, j int) int64 {
	ans := int64(0)

	n, m := len(lines), len(lines[0])
	vis := make([][]bool, n)
	for k := range vis {
		vis[k] = make([]bool, m)
	}

	dfs(lines, i, j, &vis)

	for _, row := range vis {
		for _, val := range row {
			if val {
				ans++
			}
		}
	}
	return ans
}

var dir = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func dfs(lines []string, i int, j int, vis *[][]bool) {
	if lines[i][j] == '9' {
		(*vis)[i][j] = true
		return
	}

	nextVal := byte(int(lines[i][j]-'0') + 1)
	for _, d := range dir {
		if utils.InRange(lines, i+d[0], j+d[1]) && byte(lines[i+d[0]][j+d[1]]-'0') == nextVal {
			dfs(lines, i+d[0], j+d[1], vis)
		}
	}
}

func rating(lines []string, i int, j int) int64 {
	if lines[i][j] == '9' {
		return 1
	}

	nextVal := byte(int(lines[i][j]-'0') + 1)
	cur := int64(0)
	for _, d := range dir {
		if utils.InRange(lines, i+d[0], j+d[1]) && byte(lines[i+d[0]][j+d[1]]-'0') == nextVal {
			cur += rating(lines, i+d[0], j+d[1])
		}
	}

	return cur
}
