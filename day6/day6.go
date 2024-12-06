package day6

import (
	"aoc2024/utils"
	"fmt"
	"os"
)

func DAY6() {
	lines, err := utils.ReadLines()
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	// fmt.Printf("Part 1 : %d\n", partOne(lines))
	fmt.Printf("Part 2 : %d\n", partTwo(lines))
}

func partOne(lines []string) int {
	ans := 0
	x, y := -1, -1
	for i, line := range lines {
		for j, ch := range line {
			if ch == '^' {
				x, y = i, j
				break
			}
		}
		if x != -1 && y != -1 {
			break
		}
	}

	dfs(x-1, y, -1, 0, &lines)

	return ans
}

func dfs(i int, j int, dx int, dy int, lines *[]string) {
	if !inRange(i, j, lines) {
		return
	}

	if inRange(i+dx, j+dy, lines) && (*lines)[i+dx][j+dy] == '#' {
		dx, dy = rotate90(dx, dy)
	}

	line := []byte((*lines)[i])
	line[j] = 'X'
	(*lines)[i] = string(line)

	dfs(i+dx, j+dy, dx, dy, lines)
}

func inRange(i int, j int, lines *[]string) bool {
	return (i >= 0 && i < len(*lines) && j >= 0 && j < len((*lines)[0]))
}

func rotate90(i int, j int) (int, int) {
	return j, -i
}

func partTwo(lines []string) int {
	ans := 0
	for i, line := range lines {
		bytes := []byte(line)
		for j, ch := range line {
			if ch == '.' {
				bytes[j] = '#'
				if isStuck()
			}
		}
		lines[i] = string(bytes)
	}
	return ans
}
