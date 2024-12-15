package day15

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"time"
)

func DAY15() {
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

func partOne(lines []string) int64 {
	ans := int64(0)

	var grid [][]byte
	for i, line := range lines {
		if len(line) == 0 {
			lines = lines[i+1:]
			break
		}
		grid = append(grid, []byte(line))
	}

	x, y := -1, -1
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == '@' {
				x, y = i, j
				break
			}
		}
		if x != -1 && y != -1 {
			break
		}
	}

	for _, line := range lines {
		for _, ch := range line {
			grid, x, y = makeMove(grid, x, y, ch)
		}
	}

	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 'O' {
				ans += int64(100*i + j)
			}
		}
	}

	// Print the final grid
	// for _, row := range grid {
	// 	fmt.Println(string(row))
	// }

	return ans
}

func makeMove(grid [][]byte, x, y int, ch rune) ([][]byte, int, int) {
	dx, dy := getMove(ch)
	nx, ny := x+dx, y+dy

	if utils.InRangeGrid(grid, nx, ny) && grid[nx][ny] == '.' {
		grid[x][y], grid[nx][ny] = grid[nx][ny], grid[x][y]
		x, y = nx, ny
	} else if utils.InRangeGrid(grid, nx, ny) && grid[nx][ny] == 'O' {
		for utils.InRangeGrid(grid, nx, ny) && grid[nx][ny] == 'O' {
			nx += dx
			ny += dy
		}

		if utils.InRangeGrid(grid, nx, ny) && grid[nx][ny] != '#' {
			for x != nx || y != ny {
				grid[nx][ny], grid[nx-dx][ny-dy] = grid[nx-dx][ny-dy], grid[nx][ny]
				nx -= dx
				ny -= dy
			}
		} else {
			return grid, x, y
		}
		x, y = x+dx, y+dy
	}

	return grid, x, y
}

func getMove(ch rune) (int, int) {
	if ch == '^' {
		return -1, 0
	} else if ch == 'v' {
		return 1, 0
	} else if ch == '<' {
		return 0, -1
	} else if ch == '>' {
		return 0, 1
	} else {
		return 0, 0
	}
}