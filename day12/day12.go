package day12

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"time"
)

func DAY12() {
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

	vis := make([][]bool, len(lines))
	for i := range vis {
		vis[i] = make([]bool, len(lines[0]))
	}

	for i, line := range lines {
		for j := range line {
			if !vis[i][j] && lines[i][j] != '.' {
				area, peri, _ := price(lines, i, j, &vis)
				ans += area * peri
			}
		}
	}

	return ans
}

func partTwo(lines []string) int64 {
	ans := int64(0)

	vis := make([][]bool, len(lines))
	for i := range vis {
		vis[i] = make([]bool, len(lines[0]))
	}

	for i, line := range lines {
		for j := range line {
			if !vis[i][j] && lines[i][j] != '.' {
				area, _, corners := price(lines, i, j, &vis)
				// fmt.Println(i, j, area, corners)
				ans += area * corners
			}
		}
	}

	return ans
}

func price(lines []string, i int, j int, vis *[][]bool) (int64, int64, int64) {
	area, perimeter, corners := int64(0), int64(0), int64(0)
	dfs(lines, i, j, vis, &area, &perimeter, &corners)

	return area, perimeter, corners
}

func dfs(lines []string, i int, j int, vis *[][]bool, area *int64, peri *int64, corners *int64) {
	(*vis)[i][j] = true
	*area++

	count := int64(0)
	for _, d := range utils.Dir {
		di, dj := i+d[0], j+d[1]
		if utils.InRange(lines, di, dj) && lines[di][dj] == lines[i][j] {
			count++
		}
	}
	*peri += 4 - count

	for x := 0; x < 4; x++ {
		y := (x+1) % 4
		d1, d2 := utils.Dir[x], utils.Dir[y]
		if isInnerCorner(lines, i, j, d1, d2) || isInnerCorner(lines, i, j, d2, d1) {
			*corners++
		} else if isOuterCorner(lines, i, j, d1, d2) || isOuterCorner(lines, i, j, d2, d1) {
			*corners++
		}

	}

	for _, d := range utils.Dir {
		di, dj := i+d[0], j+d[1]
		if utils.InRange(lines, di, dj) && !(*vis)[di][dj] && lines[di][dj] == lines[i][j] {
			dfs(lines, di, dj, vis, area, peri, corners)
		}
	}
}

func isInnerCorner(lines []string, i int, j int, d1 [2]int, d2 [2]int) bool {
	d1i, d1j, d2i, d2j := i+d1[0], j+d1[1], i+d2[0], j+d2[1]
	if lines[d1i][d1j] == lines[i][j] && lines[d2i][d2j] == lines[i][j] && lines[d1i][d2j] != lines[i][j] {
		// fmt.Println(i, j, d1i, d1j, d2i, d2j, true)
		return true
	}

	return false
}

func isOuterCorner(lines []string, i int, j int, d1 [2]int, d2 [2]int) bool {
	d1i, d1j, d2i, d2j := i+d1[0], j+d1[1], i+d2[0], j+d2[1]
	if lines[d1i][d1j] != lines[i][j] && lines[d2i][d2j] != lines[i][j] {
		// fmt.Println(i, j, d1i, d1j, d2i, d2j, true)
		return true
	}

	return false
}