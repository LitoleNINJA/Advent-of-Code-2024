package day18

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"time"
)

func DAY18() {
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

var coords [][2]int
var n int = 71

func partOne(lines []string) int64 {
	ans := int64(0)

	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, line := range lines {
		x, y := utils.ReadXY(line, "", "", ",")
		coords = append(coords, [2]int{int(y), int(x)})
	}

	grid = mapMemory(grid, 1024)

	grid, ans = bfs(grid)

	// for _, row := range grid {
	// 	fmt.Printf("%s\n", row)
	// }
	// fmt.Println()

	return ans
}

func mapMemory(grid [][]byte, k int) [][]byte {
	for i := range k {
		x, y := coords[i][0], coords[i][1]
		grid[x][y] = '#'
	}

	return grid
}

func bfs(grid [][]byte) ([][]byte, int64) {
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{0, 0, 0})

	for len(queue) > 0 {
		x, y, dist := queue[0][0], queue[0][1], queue[0][2]
		fmt.Println(x, y, dist)
		queue = queue[1:]
		grid[x][y] = 'O'

		if x == n-1 && y == n-1 {
			return grid, int64(dist)
		}

		for _, d := range utils.Dir {
			nx, ny := x+d[0], y+d[1]
			if utils.InRangeGrid(grid, nx, ny) && grid[nx][ny] == '.' {
				queue = append(queue, [3]int{nx, ny, dist+1})
			}
		}
	}

	fmt.Println("end not reached !")
	return grid, -1
}