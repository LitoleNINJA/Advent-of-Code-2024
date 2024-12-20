package day20

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"time"
)

func DAY20() {
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

type Point struct {
	x, y int
}

var ex, ey int

func partOne(lines []string) int64 {
	ans := int64(0)

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	sx, sy := -1, -1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				sx, sy = i, j
			} else if grid[i][j] == 'E' {
				ex, ey = i, j
				grid[i][j] = '.'
			}
		}
	}

	dist := dijkstra(grid, sx, sy)
	
	for p1, d1 := range dist {
		for p2, d2 := range dist {
			d := utils.Abs(p1.x-p2.x) + utils.Abs(p1.y-p2.y) 
			if (d == 2) && utils.Abs(d1-d2)-d >= 100 {
				ans++
			}
		}
	}

	return ans
}

func dijkstra(grid [][]byte, i, j int) map[Point]int {
	dist := make(map[Point]int, 0)
	queue := make([]Point, 0)

	queue = append(queue, Point{i, j})
	dist[Point{i, j}] = 0

	for len(queue) > 0 {
		slices.SortFunc(queue, func(a, b Point) int {
			if dist[a] < dist[b] {
				return -1
			} else if dist[a] > dist[b] {
				return 1
			}
			return 0
		})

		x, y := queue[0].x, queue[0].y
		queue = queue[1:]

		for _, d := range utils.Dir {
			dx, dy := x+d[0], y+d[1]
			if utils.InRangeGrid(grid, dx, dy) && grid[dx][dy] == '.' {
				if _, ok := dist[Point{dx, dy}]; !ok {
					dist[Point{dx, dy}] = dist[Point{x, y}] + 1
					queue = append(queue, Point{dx, dy})
				}
			}
		}
	}

	return dist
}
