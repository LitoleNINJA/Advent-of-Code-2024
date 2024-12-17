package day16

import (
	"aoc2024/utils"
	"container/heap"
	"fmt"
	"os"
	"time"
)

func DAY16() {
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

var ex, ey int

func partOne(lines []string) int64 {
	ans := int64(0)
	var grid [][]byte
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	sx, sy := -1, -1
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 'S' {
				sx, sy = i, j
			} else if grid[i][j] == 'E' {
				ex, ey = i, j
			}
		}
	}

	ans = dijkstra(grid, sx, sy)

	return ans
}

type Point struct {
	x, y int
}

func dijkstra(grid [][]byte, x, y int) int64 {
	n, m := len(grid), len(grid[0])
	dist := make([][]int64, n)
	for i := range dist {
		dist[i] = make([]int64, m)
		for j := range dist[i] {
			dist[i][j] = int64(1 << 32)
		}
	}

	dist[x][y] = 0
	pq := make(PriorityQueue, 1)
	i := 0
	pq[0] = &Item{
		value:    Point{x, y},
		priority: 0,
		index:    i,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		x, y := item.value.x, item.value.y
		if x == ex && y == ey {
			return dist[x][y]
		}

		if int64(item.priority) > dist[x][y] {
			continue
		}

		for _, d := range utils.Dir {
			nx, ny := x+d[0], y+d[1]
			if !utils.InRangeGrid(grid, nx, ny) {
				continue
			}

			if dist[nx][ny] > dist[x][y] + int64(grid[nx][ny] - '0') {
		}
	}
}
